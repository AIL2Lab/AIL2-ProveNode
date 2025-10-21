package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"AIL2-ProveNode/db"
	"AIL2-ProveNode/ail2"
	"AIL2-ProveNode/ail2/calculator"
	hmp "AIL2-ProveNode/http"
	"AIL2-ProveNode/log"
	"AIL2-ProveNode/types"
	"AIL2-ProveNode/ws"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var version string

var gpusupportfile string = "watch/gpus.json"

var defaultLogFormatter = func(params gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if params.IsOutputColor() {
		statusColor = params.StatusCodeColor()
		methodColor = params.MethodColor()
		resetColor = params.ResetColor()
	}

	if params.Latency > time.Minute {
		params.Latency = params.Latency.Truncate(time.Second)
	}
	return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v | %#v\n%s",
		params.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, params.StatusCode, resetColor,
		params.Latency,
		params.ClientIP,
		methodColor, params.Method, resetColor,
		params.Path,
		params.Request.UserAgent(),
		params.ErrorMessage,
	)
}

func main() {
	configPath := flag.String("config", "", "run using the configuration file")
	versionFlag := flag.Bool("version", false, "show version number and exit")
	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}
	if *configPath == "" {
		fmt.Println("run command like 'app -config ./config.json'")
		os.Exit(1)
	}
	cfg, err := types.LoadConfig(*configPath)
	if err != nil {
		fmt.Println("Failed to load JSON configuration file:", err)
		os.Exit(1)
	}
	if err := log.InitLogrus(cfg.LogLevel, cfg.LogFile); err != nil {
		fmt.Println("Initialize the log failed:", err)
		os.Exit(1)
	}

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := db.InitMongo(ctx, cfg.MongoDB.URI, cfg.MongoDB.Database, cfg.MongoDB.ExpireTime); err != nil {
		os.Exit(1)
	}
	if err := db.InitIP2LDB(cfg.IP2LDB.DatabasePath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.IP2LDB.Close()

	if err := dbc.InitDbcChain(ctx, cfg.Chain); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := calculator.LoadGpuList(gpusupportfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Create file watcher failed:", err)
	}
	// defer watcher.Close()

	wsCtx, wsCancel := context.WithCancel(context.Background())
	wsHub, err := ws.InitHub(wsCtx, cfg.NotifyThirdParty.OfflineNotify)
	if err != nil {
		fmt.Println("Failed to init websocket hub:", err)
		os.Exit(1)
	}

	pm := hmp.NewPrometheusMetrics(cfg.Prometheus.JobName)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Log.Warn("watcher event not ok")
					return
				}
				log.Log.Infof("event: %v", event)
				if event.Has(fsnotify.Write) {
					log.Log.Infof("modified file: %v", event.Name)
					if strings.HasSuffix(event.Name, gpusupportfile) {
						if err := calculator.LoadGpuList(event.Name); err != nil {
							log.Log.Errorf("gpu support file changed, but %v", err)
						}
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					log.Log.Warn("watcher error not ok")
					return
				}
				log.Log.Infof("watcher error: %v", err)
			}
		}
	}()

	if err := watcher.Add("watch"); err != nil {
		fmt.Println("Add watcher failed:", err)
		os.Exit(1)
	}

	gin.DefaultWriter = log.Log.Out
	// router := gin.Default()
	router := gin.New()
	// router.Use(func(ctx *gin.Context) {
	// 	start := time.Now()
	// 	ctx.Next()
	// 	latency := time.Since(start)
	// 	log.Log.WithFields(logrus.Fields{
	// 		"client_ip":  ctx.ClientIP(),
	// 		"timestamp":  start.Format(time.RFC1123),
	// 		"method":     ctx.Request.Method,
	// 		"path":       ctx.Request.URL.Path,
	// 		"protocol":   ctx.Request.Proto,
	// 		"status":     ctx.Writer.Status(),
	// 		"latency":    latency,
	// 		"user_agent": ctx.Request.UserAgent(),
	// 		"errors":     ctx.Errors.ByType(gin.ErrorTypePrivate).String(),
	// 	}).Info("request details")
	// })
	corsConfig := cors.DefaultConfig()
	// corsConfig.AllowAllOrigins = true
	corsConfig.AllowOrigins = []string{"*"}
	// corsConfig.AllowOrigins = []string{"https://example.com"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	// corsConfig.AllowHeaders = []string{"Origin", "Content-Type"}
	corsConfig.AllowCredentials = true
	router.Use(gin.LoggerWithFormatter(defaultLogFormatter), gin.Recovery(), cors.New(corsConfig))
	router.GET("/metrics/prometheus", pm.Metrics)
	// router.GET("/echo", ws.Echo)
	router.GET("/api/v0/location", hmp.Location)
	router.GET("/api/v0/calculator/point", calculator.CalculatePointFromHttp)
	// for dbc contract
	c0 := router.Group("/api/v0/contract")
	{
		c0.POST("/register", func(ctx *gin.Context) {
			wg.Add(1)
			defer wg.Done()
			hmp.RegisterMachine(ctx)
		})
		c0.POST("/unregister", func(ctx *gin.Context) {
			wg.Add(1)
			defer wg.Done()
			hmp.UnregisterMachine(ctx, wsHub.SendUnregisterNotify)
		})
		c0.POST("/online", func(ctx *gin.Context) {
			wg.Add(1)
			defer wg.Done()
			hmp.OnlineMachine(ctx)
		})
		c0.POST("/offline", func(ctx *gin.Context) {
			wg.Add(1)
			defer wg.Done()
			hmp.OfflineMachine(ctx)
		})
	}
	router.GET("/websocket", func(c *gin.Context) {
		wg.Add(1)
		defer wg.Done()
		// ws.Ws(c, pm)
		ws.Ws2(wsHub, c, wsCtx)
	})

	// log.Log.Fatal(router.Run(cfg.Addr))

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	// if cfg.Certificate.Cert != "" && cfg.Certificate.Key != "" {
	// 	cert, err := tls.LoadX509KeyPair(cfg.Certificate.Cert, cfg.Certificate.Key)
	// 	if err != nil {
	// 		log.Log.Fatalf("Failed to load x509 certificate: %v", err)
	// 	}
	// 	srv.TLSConfig = &tls.Config{
	// 		Certificates: []tls.Certificate{cert},
	// 		MinVersion:   tls.VersionTLS12,
	// 	}
	// }

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if cfg.Certificate.Cert != "" && cfg.Certificate.Key != "" {
			if err := srv.ListenAndServeTLS(cfg.Certificate.Cert, cfg.Certificate.Key); err != nil && err != http.ErrServerClosed {
				log.Log.Fatalf("Start server: %v", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Log.Fatalf("Start server: %v", err)
			}
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Log.Println("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx1, cancel1 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel1()
	if err := srv.Shutdown(ctx1); err != nil {
		log.Log.Fatal("Server forced to shutdown: ", err)
	}

	watcher.Close()
	// ws.ShutdownAllWsConns()
	wsHub.Close()
	wg.Wait()

	wsCancel()
	wsHub.Wait()

	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
	if err := db.MDB.Disconnect(ctx2); err != nil {
		log.Log.Fatal("Disconnect mongo client: ", err)
	}

	log.Log.Println("Server exiting")
}
