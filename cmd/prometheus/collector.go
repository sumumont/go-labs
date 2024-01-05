package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metric struct {
	MetricType prometheus.ValueType `json:"metricType"`
	Name       string               `json:"name"`
	Value      float64              `json:"value"`
	Desc       string               `json:"desc"`
	Tag        map[string]string    `json:"tag"`
}

type Collector interface {
	Update(ch chan<- prometheus.Metric) error
}

type SlurmCollector struct {
	Metrics []Metric `json:"metrics"`
}

const namespace = "slurm"

func (c *SlurmCollector) Update(ch chan<- prometheus.Metric) error {
	const subsystem = "node_status"

	metrics := c.Metrics
	for _, m := range metrics {
		labelValues := []string{"nodename"}

		desc := prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, m.Name),
			m.Desc,
			labelValues,
			nil,
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			m.MetricType,
			m.Value,
			labelValues...,
		)
	}
	return nil
}

// Describe 实现Collector接口的Describe方法
func (c *SlurmCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.customMetric.Desc()
}

// Collect 实现Collector接口的Collect方法
func (c *SlurmCollector) Collect(ch chan<- prometheus.Metric) {
	// 模拟收集指标的数据
	value := 42.0 // 这里可以根据需要更改指标的值
	c.customMetric.Set(value)

	// 将指标发送给Prometheus
	c.customMetric.Collect(ch)
}
