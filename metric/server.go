package metric

import (
	"github.com/bnb-chain/greenfield-data-marketplace-backend/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

const (
	MetricNameGnfdSavedBlock = "gnfd_saved_block_height"
	MetricNameBSCSavedBlock  = "bsc_saved_block_height"
)

type MetricService struct {
	MetricsMap map[string]prometheus.Metric
	cfg        *util.MonitorConfig
}

func NewMetricService(config *util.MonitorConfig) *MetricService {
	ms := make(map[string]prometheus.Metric, 0)
	labels := map[string]string{
		"gnfdChainId": config.GnfdChainId,
	}

	// Greenfield
	gnfdSavedBlockMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        MetricNameGnfdSavedBlock,
		Help:        "Saved block height for Greenfield in Database",
		ConstLabels: labels,
	})
	ms[MetricNameGnfdSavedBlock] = gnfdSavedBlockMetric
	prometheus.MustRegister(gnfdSavedBlockMetric)

	// BSC
	bscSavedBlockMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        MetricNameBSCSavedBlock,
		Help:        "Saved block height for BSC in Database",
		ConstLabels: labels,
	})
	ms[MetricNameBSCSavedBlock] = bscSavedBlockMetric
	prometheus.MustRegister(bscSavedBlockMetric)

	return &MetricService{
		MetricsMap: ms,
		cfg:        config,
	}
}

func (m *MetricService) Start() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9292", nil)
	if err != nil {
		panic(err)
	}
}

func (m *MetricService) SetGnfdSavedBlockHeight(height uint64) {
	m.MetricsMap[MetricNameGnfdSavedBlock].(prometheus.Gauge).Set(float64(height))
}

func (m *MetricService) SetBscSavedBlockHeight(height uint64) {
	m.MetricsMap[MetricNameBSCSavedBlock].(prometheus.Gauge).Set(float64(height))
}
