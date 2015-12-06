package monitor

import (
	"../data"
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

	apiTimeoutCountCollector = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_timeout_counter",
			Help: "api timeout counter.",
		},
		[]string{"api_name"},
	)

	apiMaxUseTimeCollector = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "api_max_usetime",
			Help: "api max use time.",
		},
		[]string{"api_name"},
	)

	serviceAnalyseResults = make(map[string]*data.AnalyseResult)
	apiAnalyseResults     = make(map[string]*data.AnalyseResult)
)

func init() {
	prometheus.MustRegister(serviceTimeoutCountCollector)
	prometheus.MustRegister(serviceMaxUseTimeCollector)
	//prometheus.MustRegister(apiTimeoutCountCollector)
	//prometheus.MustRegister(apiMaxUseTimeCollector)
}

func StartServer(serverHost string, dbHost string, refreshInterval int) {
	flag.Parse()

	go func() {
		now := time.Now()
		gtTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), data.Location)

		for {

			dataList, err := data.LoadTrackingData(dbHost, &gtTime)
			if err != nil {
				log.Println(err)
				continue
			}
			data.TrackTimeoutAnalysePerService(0, dataList, serviceAnalyseResults)

			for k, _ := range serviceAnalyseResults {
				serviceTimeoutCountCollector.WithLabelValues(k).Set(float64(serviceAnalyseResults[k].Count))
				serviceMaxUseTimeCollector.WithLabelValues(k).Set(float64(serviceAnalyseResults[k].MaxUseTime) / 1000)
				//init data
				serviceAnalyseResults[k].Count = 0
				serviceAnalyseResults[k].MaxUseTime = 0
			}
			/*
				data.TrackTimeoutAnalysePerAPI(0, dataList, apiAnalyseResults)

				for k, _ := range apiAnalyseResults {
					apiTimeoutCountCollector.WithLabelValues(k).Set(float64(apiAnalyseResults[k].Count))
					apiMaxUseTimeCollector.WithLabelValues(k).Set(float64(apiAnalyseResults[k].MaxUseTime) / 1000)
					//init data
					apiAnalyseResults[k].Count = 0
					apiAnalyseResults[k].MaxUseTime = 0
				}
			*/
			time.Sleep(time.Second * time.Duration(refreshInterval))
		}
	}()

	addr := flag.String("listen-address", serverHost, "The address to listen on for HTTP requests.")
	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(*addr, nil)
}
