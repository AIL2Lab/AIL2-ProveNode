package http

import (
	"bytes"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
)

// go test -v -timeout 30s -count=1 -run TestPrometheusGauge AIL2-ProveNode/http
func TestPrometheusGauge(t *testing.T) {
	temp := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "temperature",
			Help: "The temperature of celsuis",
		},
		[]string{"job", "instance"},
	)
	temp.WithLabelValues("test", "c1").Set(36)
	temp.WithLabelValues("test", "c2").Set(39)
	hum := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "humidity",
			Help: "The humidity of celsuis",
		},
		[]string{"job", "instance"},
	)
	hum.WithLabelValues("test", "c1").Set(36)
	hum.WithLabelValues("test", "c2").Set(39)

	reg := prometheus.NewRegistry()
	reg.MustRegister(temp)
	reg.MustRegister(hum)

	metricsFamilies, err := reg.Gather()
	if err != nil || len(metricsFamilies) < 1 {
		t.Fatal("unexpected behavior of custom test registry ", err)
	}
	for i, mf := range metricsFamilies {
		t.Logf("metrics item %v: %v", i, mf.String())
		// prometheus.WriteToTextfile()
		buf := &bytes.Buffer{}
		written, err := expfmt.MetricFamilyToText(buf, mf)
		if err != nil {
			t.Fatalf("MetricFamilyToText failed: %v", err)
		}
		t.Logf("MetricFamilyToText %v bytes\n%v", written, buf.String())
	}

	t.Log("Remove instance of c2")

	temp.DeleteLabelValues("test", "c2")
	hum.DeleteLabelValues("test", "c2")

	// it will not panic when delete one lable twice
	temp.DeleteLabelValues("test", "c2")
	hum.DeleteLabelValues("test", "c2")

	metricsFamilies, err = reg.Gather()
	if err != nil || len(metricsFamilies) < 1 {
		t.Fatal("unexpected behavior of custom test registry ", err)
	}
	for i, mf := range metricsFamilies {
		t.Logf("metrics item %v: %v", i, mf.String())
		// prometheus.WriteToTextfile()
		buf := &bytes.Buffer{}
		written, err := expfmt.MetricFamilyToText(buf, mf)
		if err != nil {
			t.Fatalf("MetricFamilyToText failed: %v", err)
		}
		t.Logf("MetricFamilyToText %v bytes\n%v", written, buf.String())
	}

	t.Log("Remove metrics of humidity")

	reg.Unregister(hum)

	metricsFamilies, err = reg.Gather()
	if err != nil || len(metricsFamilies) < 1 {
		t.Fatal("unexpected behavior of custom test registry ", err)
	}
	for i, mf := range metricsFamilies {
		t.Logf("metrics item %v: %v", i, mf.String())
		// prometheus.WriteToTextfile()
		buf := &bytes.Buffer{}
		written, err := expfmt.MetricFamilyToText(buf, mf)
		if err != nil {
			t.Fatalf("MetricFamilyToText failed: %v", err)
		}
		t.Logf("MetricFamilyToText %v bytes\n%v", written, buf.String())
	}
}

// go test -v -timeout 30s -count=1 -run TestPrometheusTimestamp AIL2-ProveNode/http
func TestPrometheusTimestamp(t *testing.T) {
	desc := prometheus.NewDesc(
		"temperature",
		"The temperature of celsuis",
		nil, nil,
	)
	temperatureReportedByExternalSystem := 298.15
	timeReportedByExternalSystem := time.Date(2009, time.November, 10, 23, 0, 0, 12345678, time.UTC)
	s := prometheus.NewMetricWithTimestamp(
		timeReportedByExternalSystem,
		prometheus.MustNewConstMetric(
			desc, prometheus.GaugeValue, temperatureReportedByExternalSystem,
		),
	)

	_ = s

	// reg := prometheus.NewRegistry()
	// reg.MustRegister(s)

	// buf := &bytes.Buffer{}
	// enc := expfmt.NewEncoder(buf, expfmt.NewFormat(expfmt.TypeOpenMetrics))
	// enc.Encode(&s)
	// s.Write()
	// written, err := expfmt.MetricFamilyToText(buf, &s)
	// if err != nil {
	// 	t.Fatalf("MetricFamilyToText failed: %v", err)
	// }
	// t.Logf("MetricFamilyToText %v bytes\n%v", written, buf.String())
}
