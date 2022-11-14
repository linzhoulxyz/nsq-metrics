package collector

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type producerStats []struct {
	val func(*producer) float64
	vec *prometheus.GaugeVec
}

// ProducerStats creates a new stats collector which is able to
// expose the procucer metrics of a nsqd node to Prometheus.
func ProducerStats(namespace string) StatsCollector {
	labels := []string{"client_id", "hostname", "version", "remote_address", "user_agent", "deflate", "snappy", "topic", "tls"}
	namespace += "_producer"

	stats := producerStats{
		{
			// TODO: Give state a descriptive name instead of a number.
			val: func(p *producer) float64 { return float64(p.State) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "state",
				Help:      "State of producer",
			}, labels),
		},
		{
			val: func(p *producer) float64 { return float64(p.ReadyCount) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "ready_count",
				Help:      "Ready count",
			}, labels),
		},
		{
			val: func(p *producer) float64 { return float64(p.InFlightCount) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "in_flight_count",
				Help:      "In flight count",
			}, labels),
		},
		{
			val: func(p *producer) float64 { return float64(p.MessageCount) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "message_count",
				Help:      "Message count",
			}, labels),
		},
		{
			val: func(p *producer) float64 { return float64(p.FinishCount) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "finish_count",
				Help:      "Finish count",
			}, labels),
		},
		{
			val: func(p *producer) float64 { return float64(p.RequeueCount) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "requeue_count",
				Help:      "Requeue count",
			}, labels),
		},
		{
			val: func(p *producer) float64 { return float64(p.SampleRate) },
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "sample_rate",
				Help:      "Sample Rate",
			}, labels),
		},
		{
			val: func(p *producer) float64 { return float64(p.PubCounts[0].Count) }, // TODO: Think about how to handle multiple topics.
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "pub_rate",
				Help:      "Producer publish Rate",
			}, labels),
		},
	}

	return stats
}

func (cs producerStats) set(s *stats) {
	for _, producer := range s.Producers {
		for _, pub_counts := range producer.PubCounts {
			labels := prometheus.Labels{
				"client_id":      producer.ClientId,
				"hostname":       producer.Hostname,
				"version":        producer.Version,
				"remote_address": producer.RemoteAddress,
				"user_agent":     producer.UserAgent,
				"deflate":        strconv.FormatBool(producer.Deflate),
				"snappy":         strconv.FormatBool(producer.Snappy),
				"topic":          pub_counts.Topic,
				"tls":            strconv.FormatBool(producer.TLS),
			}

			for _, c := range cs {
				c.vec.With(labels).Set(c.val(producer))
			}
		}
	}
}

func (cs producerStats) collect(out chan<- prometheus.Metric) {
	for _, c := range cs {
		c.vec.Collect(out)
	}
}

func (cs producerStats) describe(ch chan<- *prometheus.Desc) {
	for _, c := range cs {
		c.vec.Describe(ch)
	}
}

func (cs producerStats) reset() {
	for _, c := range cs {
		c.vec.Reset()
	}
}
