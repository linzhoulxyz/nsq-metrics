# NSQ Metrics

[![GoDoc](https://pkg.go.dev/badge/github.com/spiegeltechlab/nsq-metrics.svg)](https://pkg.go.dev/github.com/spiegeltechlab/nsq-metrics)

NSQ metrics exporter for prometheus.io, written in go.

## Building

    make

    OR

    go install github.com/spiegeltechlab/nsq-metrics@latest

## Config

env | flag | description | default
--- | --- | --- | ---
NSQ_METRICS_WEB_LISTEN_ADDRESS | web.listen | Address on which to expose metrics and web interface. | :9117
NSQ_METRICS_WEB_PATH | web.path | Path under which to expose metrics. | /metrics
NSQ_METRICS_NSQD_ADDRESS | nsqd.address | Address of the nsqd node. | http://localhost:4151/stats
NSQ_METRICS_ENABLED_COLLECTORS | collectors | Comma-separated list of collectors to use. See collectors below for possible values | stats.topics,stats.channels
NSQ_METRICS_NAMESPACE | namespace | Namespace for the NSQ metrics. | nsq
NSQ_METRICS_TLS_CA_CERT | tls.ca-cert | CA certificate file to be used for nsqd connections. | ""
NSQ_METRICS_TLS_CERT | tls.cert | TLS certificate file to be used for client connections to nsqd. | ""
NSQ_METRICS_TLS_KEY | tls.key | TLS key file to be used for TLS client connections to nsqd. | ""

## Collectors for ENABLED_COLLECTORS or collectors config

### Topics (stats.topics)
Exported metrics:
`nsq_topic_backend_depth`
`nsq_topic_channel_count`
`nsq_topic_depth`
`nsq_topic_e2e_latency_95_percentile`
`nsq_topic_e2e_latency_99_percentile`
`nsq_topic_message_count`

### Channels (stats.channels)
Exported metrics:
`nsq_channel_backend_depth`
`nsq_channel_client_count`
`nsq_channel_deferred_count`
`nsq_channel_depth`
`nsq_channel_e2e_latency_95p`
`nsq_channel_e2e_latency_99p`
`nsq_channel_in_flight_count`
`nsq_channel_message_count`
`nsq_channel_requeue_count`
`nsq_channel_requeue_count`
`nsq_channel_timeout_count`

### Clients (stats.clients)
`nsq_client_connect_ts`
`nsq_client_finish_count`
`nsq_client_in_flight_count`
`nsq_client_message_count`
`nsq_client_ready_count`
`nsq_client_requeue_count`
`nsq_client_sample_rate`
`nsq_client_state`

### Producers (stats.producers)
`nsq_producer_finish_count`
`nsq_producer_in_flight_count`
`nsq_producer_message_count`
`nsq_producer_pub_rate`
`nsq_producer_ready_count`
`nsq_producer_requeue_count`
`nsq_producer_sample_rate`
`nsq_producer_state`