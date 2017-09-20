package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

// In the response of localhost:8080/metrics
//
//HELP my_collected My collected metric
//# TYPE my_collected gauge
//my_collected 8
type myCollector struct {
	metricDesc *prometheus.Desc

	counter int64
}

func NewMyCollector() prometheus.Collector {
	return &myCollector{
		metricDesc: prometheus.NewDesc(
			"my_collected",
			"My collected metric",
			nil, nil),
		counter: 0,
	}
}

func (c *myCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.metricDesc
}

func (c *myCollector) Collect(ch chan<- prometheus.Metric) {
	c.counter++
	ch <- prometheus.MustNewConstMetric(c.metricDesc, prometheus.GaugeValue, float64(c.counter))
	//MustNewConstMetric() is used to avoid concurrent scrapes
}
