package monitor

import (
	"../alarm"
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"time"
)

//var addr = flag.String("listen-address", ":8888", "The address to listen on for HTTP requests.")

var (
	serviceTimeoutCountCollector = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_timeout_counter",
			Help: "service timeout counter.",
		},
		[]string{"service_name"},
	)

	/*
		serviceMaxUseTimeCollector = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "service_max_usetime_minute",
				Help:    "service max use time per minute.",
				Buckets: prometheus.LinearBuckets(10, 10, 20),
			},
			[]string{"service_max_use_time"},
		)
	*/

	serviceMaxUseTimeCollector = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "service_max_usetime",
			Help: "service max use time.",
		},
		[]string{"service_name"},
	)
)

func init() {
	prometheus.MustRegister(serviceTimeoutCountCollector)
	prometheus.MustRegister(serviceMaxUseTimeCollector)
}

func StartServer(serverHost string, dbHost string, refreshInterval int) {
	flag.Parse()

	go func() {
		now := time.Now()
		gtTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), alarm.Location)
		for {
			dataList, err := alarm.LoadTrackingData(dbHost, &gtTime)
			if err != nil {
				log.Println(err)
				continue
			}
			analyseResults := alarm.TrackTimeoutAnalysePerService(0, dataList)

			for k, _ := range analyseResults {
				serviceTimeoutCountCollector.WithLabelValues(k).Set(float64(analyseResults[k].Count))
				serviceMaxUseTimeCollector.WithLabelValues(k).Set(float64(analyseResults[k].MaxUseTime) / 1000)
			}

			time.Sleep(time.Second * time.Duration(refreshInterval))
		}
	}()

	addr := flag.String("listen-address", serverHost, "The address to listen on for HTTP requests.")
	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(*addr, nil)
}
