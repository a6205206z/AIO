package monitor

import (
	"../alarm"
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"time"
)

var (
	serviceTimeoutCountCollector = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_timeout_counter",
			Help: "service timeout counter.",
		},
		[]string{"service_name"},
	)

	serviceMaxUseTimeCollector = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "service_max_usetime",
			Help: "service max use time.",
		},
		[]string{"service_name"},
	)

	analyseResults = make(map[string]*alarm.AnalyseResult)
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
			alarm.TrackTimeoutAnalysePerService(0, dataList, analyseResults)

			for k, _ := range analyseResults {
				serviceTimeoutCountCollector.WithLabelValues(k).Set(float64(analyseResults[k].Count))
				serviceMaxUseTimeCollector.WithLabelValues(k).Set(float64(analyseResults[k].MaxUseTime) / 1000)
				//init data
				analyseResults[k].Count = 0
				analyseResults[k].MaxUseTime = 0
			}

			time.Sleep(time.Second * time.Duration(refreshInterval))
		}
	}()

	addr := flag.String("listen-address", serverHost, "The address to listen on for HTTP requests.")
	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(*addr, nil)
}
