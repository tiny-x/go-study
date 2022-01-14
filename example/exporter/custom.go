package main

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

type Exporter struct {
	networkConnectionFd *prometheus.Desc
	processThreadCounts *prometheus.Desc
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.networkConnectionFd
	ch <- e.processThreadCounts
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	connections, err := net.Connections("tcp")
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}

	for _, connection := range connections {
		ch <- prometheus.MustNewConstMetric(
			e.networkConnectionFd,
			prometheus.GaugeValue,
			float64(connection.Fd),
			fmt.Sprintf("%s:%s", connection.Laddr.IP, strconv.Itoa(int(connection.Laddr.Port))),
			fmt.Sprintf("%s:%s", connection.Raddr.IP, strconv.Itoa(int(connection.Raddr.Port))),
			connection.Status,
			strconv.Itoa(int(connection.Type)),
			strconv.Itoa(int(connection.Family)),
			strconv.Itoa(int(connection.Pid)),
		)
	}

	processes, err := process.Processes()
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
	for _, p := range processes {

		ch <- prometheus.MustNewConstMetric(
			e.processThreadCounts,
			prometheus.GaugeValue,
			func() float64 {
				if cpuPercent, err := p.CPUPercent(); err == nil {
					return cpuPercent
				}
				return 0
			}(),
			strconv.Itoa(int(p.Pid)),
			func() string {
				if name, err := p.Name(); err == nil {
					return name
				}
				return ""
			}(),
			func() string {
				if cmd, err := p.Cmdline(); err == nil {
					return cmd
				}
				return ""
			}(),
		)
	}
}

func NewExporter() *Exporter {
	return &Exporter{
		networkConnectionFd: prometheus.NewDesc(
			"xhas_agent_network_connection_fd",
			"network connection info",
			[]string{"lAddr", "rAddr", "status", "type", "family", "pid"},
			nil,
		),
		processThreadCounts: prometheus.NewDesc(
			"xhas_agent_process_cpu_percent",
			"process cpu percent",
			[]string{"pid", "name", "cmd"},
			nil,
		),
	}
}

func main() {
	exporter := NewExporter()
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(exporter)

	gatherers := prometheus.Gatherers{
		//prometheus.DefaultGatherer,
		reg,
	}

	h := promhttp.HandlerFor(gatherers,
		promhttp.HandlerOpts{
			ErrorLog:      log.NewErrorLogger(),
			ErrorHandling: promhttp.ContinueOnError,
		})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	log.Fatal(http.ListenAndServe(":8999", nil))
}
