# NSQ Metrics

[![GoDoc](https://pkg.go.dev/badge/github.com/leonkunert/nsq-metrics.svg)](https://pkg.go.dev/github.com/leonkunert/nsq-metrics)

NSQ metrics exporter for prometheus.io, written in go.

## Building

    make

    OR

    go install github.com/leonkunert/nsq-metrics@latest

## Config

env | flag | description | default
--- | --- | --- | ---
NSQ_METRICS_WEB_LISTEN_ADDRESS | web.listen | Address on which to expose metrics and web interface. | :9117
NSQ_METRICS_WEB_PATH | web.path | Path under which to expose metrics. | /metrics
NSQ_METRICS_NSQD_ADDRESS | nsqd.address | Address of the nsqd node. | http://localhost:4151/stats
NSQ_METRICS_ENABLED_COLLECTORS | collectors | Comma-separated list of collectors to use. | stats.topics,stats.channels
NSQ_METRICS_NAMESPACE | namespace | Namespace for the NSQ metrics. | nsq
NSQ_METRICS_TLS_CA_CERT | tls.ca-cert | CA certificate file to be used for nsqd connections. | ""
NSQ_METRICS_TLS_CERT | tls.cert | TLS certificate file to be used for client connections to nsqd. | ""
NSQ_METRICS_TLS_KEY | tls.key | TLS key file to be used for TLS client connections to nsqd. | ""
