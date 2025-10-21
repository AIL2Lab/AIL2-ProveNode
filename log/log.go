package log

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log = logrus.New()

func InitLogrus(levelString string, logFile string) error {
	logLevel, err := logrus.ParseLevel(levelString)
	if err != nil {
		return err
	}
	Log.SetFormatter(&logrus.JSONFormatter{
		DisableHTMLEscape: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	Log.SetReportCaller(true)
	Log.SetLevel(logLevel)
	// w1 := os.Stdout
	// w2 := &lumberjack.Logger{
	// 	Filename:   logFile,
	// 	MaxBackups: 10,
	// 	MaxAge:     10,
	// 	MaxSize:    10,
	// 	Compress:   true,
	// 	// LocalTime:  false,
	// }
	// Log.SetOutput(io.MultiWriter(w1, w2))
	Log.SetOutput(&lumberjack.Logger{
		Filename:   logFile,
		MaxBackups: 30,
		MaxAge:     28,
		MaxSize:    50,
		Compress:   true,
		// LocalTime:  false,
	})

	return nil
}
