package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// https://pkg.go.dev/github.com/prometheus/client_golang@v1.20.2/prometheus

type PrometheusMetrics struct {
	jobName             string
	reg                 *prometheus.Registry
	utilizationGPUGauge *prometheus.GaugeVec
	memoryTotalGauge    *prometheus.GaugeVec
	memoryUsedGauge     *prometheus.GaugeVec
}

func NewPrometheusMetrics(jobName string) *PrometheusMetrics {
	pm := &PrometheusMetrics{
		jobName: jobName,
		reg:     prometheus.NewRegistry(),
		utilizationGPUGauge: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "utilization_gpu",
				Help: "utilization of GPU",
			},
			[]string{"job", "instance"},
		),
		memoryTotalGauge: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "memory_total",
				Help: "total memory",
			},
			[]string{"job", "instance"},
		),
		memoryUsedGauge: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "memory_used",
				Help: "used memory",
			},
			[]string{"job", "instance"},
		),
	}
	pm.reg.MustRegister(pm.utilizationGPUGauge)
	pm.reg.MustRegister(pm.memoryTotalGauge)
	pm.reg.MustRegister(pm.memoryUsedGauge)
	return pm
}

// func (pm PrometheusMetrics) SetMetrics(id string, info types.WsMachineInfoRequest) {
// 	if pm.jobName == "" {
// 		return
// 	}
// 	pm.utilizationGPUGauge.WithLabelValues(pm.jobName, id).Set(float64(info.UtilizationGPU))
// 	pm.memoryTotalGauge.WithLabelValues(pm.jobName, id).Set(float64(info.MemoryTotal))
// 	pm.memoryUsedGauge.WithLabelValues(pm.jobName, id).Set(float64(info.MemoryUsed))
// }

func (pm PrometheusMetrics) DeleteMetrics(id string) {
	if pm.jobName == "" {
		return
	}
	pm.utilizationGPUGauge.DeleteLabelValues(pm.jobName, id)
	pm.memoryTotalGauge.DeleteLabelValues(pm.jobName, id)
	pm.memoryUsedGauge.DeleteLabelValues(pm.jobName, id)
}

func (pm PrometheusMetrics) Metrics(ctx *gin.Context) {
	w, r := ctx.Writer, ctx.Request
	if pm.jobName == "" {
		ctx.String(http.StatusInternalServerError, "job name is empty")
		return
	}

	// pm.utilizationGPUGauge.WithLabelValues("test", "machine1").Set(30)
	// pm.memoryTotalGauge.WithLabelValues("test", "machine1").Set(24564)
	// pm.memoryUsedGauge.WithLabelValues("test", "machine1").Set(22128)
	// pm.memoryUsedGauge.With(prometheus.Labels{"job": "test", "instance": "machine1"}).Set(22128)

	h := promhttp.HandlerFor(pm.reg, promhttp.HandlerOpts{Registry: pm.reg})
	h.ServeHTTP(w, r)
}
