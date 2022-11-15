package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spiegeltechlab/nsq-metrics/collector"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	// Version of nsq_exporter. Set at build time.
	Version = "0.0.0.dev"

	prefix = "NSQ_METRICS"

	// Env Vars
	listenAddressEnv     = prefix + "_WEB_LISTEN_ADDRESS"
	metricsPathEnv       = prefix + "_WEB_PATH"
	nsqdURLEnv           = prefix + "_NSQD_ADDRESS"
	enabledCollectorsEnv = prefix + "_ENABLED_COLLECTORS"
	namespaceEnv         = prefix + "_NAMESPACE"
	tlsCACertEnv         = prefix + "_TLS_CA_CERT"
	tlsCertEnv           = prefix + "_TLS_CERT"
	tlsKeyEnv            = prefix + "_TLS_KEY"

	// Default values for flags and env
	defaultListenAddress     = ":9117"
	defaultMetricsPath       = "/metrics"
	defaultNsqdURL           = "http://localhost:4151/stats"
	defaultEnabledCollectors = "stats.topics,stats.channels"
	defaultNamespace         = "nsq"
)

var (
	envListenAddress     = envString(listenAddressEnv, defaultListenAddress)
	envMetricsPath       = envString(metricsPathEnv, defaultMetricsPath)
	envNsqdURL           = envString(nsqdURLEnv, defaultNsqdURL)
	envEnabledCollectors = envString(enabledCollectorsEnv, defaultEnabledCollectors)
	envNamespace         = envString(namespaceEnv, defaultNamespace)

	listenAddress     = flag.String("web.listen", envListenAddress, "Address on which to expose metrics and web interface.")
	metricsPath       = flag.String("web.path", envMetricsPath, "Path under which to expose metrics.")
	nsqdURL           = flag.String("nsqd.address", envNsqdURL, "Address of the nsqd node.")
	enabledCollectors = flag.String("collect", envEnabledCollectors, "Comma-separated list of collectors to use.")
	namespace         = flag.String("namespace", envNamespace, "Namespace for the NSQ metrics.")
	tlsCACert         = flag.String("tls.ca_cert", "", "CA certificate file to be used for nsqd connections.")
	tlsCert           = flag.String("tls.cert", "", "TLS certificate file to be used for client connections to nsqd.")
	tlsKey            = flag.String("tls.key", "", "TLS key file to be used for TLS client connections to nsqd.")

	statsRegistry = map[string]func(namespace string) collector.StatsCollector{
		"topics":    collector.TopicStats,
		"channels":  collector.ChannelStats,
		"clients":   collector.ClientStats,
		"producers": collector.ProducerStats,
	}
)

func main() {
	flag.Parse()

	ex, err := createNsqExecutor()
	if err != nil {
		log.Fatalf("error creating nsq executor: %v", err)
	}
	prometheus.MustRegister(ex)

	http.Handle(*metricsPath, promhttp.Handler())
	if *metricsPath != "" && *metricsPath != "/" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`<html>
			<head><title>NSQ Exporter</title></head>
			<body>
			<h1>NSQ Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
		})
	}

	log.Print("listening to ", *listenAddress)
	err = http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func createNsqExecutor() (*collector.NsqExecutor, error) {
	nsqdURL, err := normalizeURL(*nsqdURL)
	if err != nil {
		return nil, err
	}

	ex, err := collector.NewNsqExecutor(*namespace, nsqdURL, *tlsCACert, *tlsCert, *tlsKey)
	if err != nil {
		log.Fatal(err)
	}
	for _, param := range strings.Split(*enabledCollectors, ",") {
		param = strings.TrimSpace(param)
		parts := strings.SplitN(param, ".", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid collector name: %s", param)
		}
		if parts[0] != "stats" {
			return nil, fmt.Errorf("invalid collector prefix: %s", parts[0])
		}

		name := parts[1]
		c, has := statsRegistry[name]
		if !has {
			return nil, fmt.Errorf("unknown stats collector: %s", name)
		}
		ex.Use(c(*namespace))
	}
	return ex, nil
}

func envString(name string, def string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return def
}

func normalizeURL(ustr string) (string, error) {
	ustr = strings.ToLower(ustr)
	if !strings.HasPrefix(ustr, "https://") && !strings.HasPrefix(ustr, "http://") {
		ustr = "http://" + ustr
	}

	u, err := url.Parse(ustr)
	if err != nil {
		return "", err
	}
	if u.Path == "" {
		u.Path = "/stats"
	}
	u.RawQuery = "format=json"
	return u.String(), nil
}
