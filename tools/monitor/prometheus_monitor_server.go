package monitor

import (
	"../data"
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	serviceReqCountCollector = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_req_counter",
			Help: "service request details counter.",
		},
		[]string{"service_name", "status_code", "client_ip", "url"},
	)

	serviceReqAvgUseTime = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_avg_use_time",
			Help: "service avg use time",
		},
		[]string{"service_name", "status_code", "client_ip", "url"},
	)

	serviceMaxUseTime = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_max_use_time",
			Help: "service max use time",
		},
		[]string{"service_name", "status_code", "client_ip", "url"},
	)

	serviceAnalyseResults = make(map[string]*data.AnalyseResult)
)

func init() {
	prometheus.MustRegister(serviceReqCountCollector)
	prometheus.MustRegister(serviceReqAvgUseTime)
	prometheus.MustRegister(serviceMaxUseTime)
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
			data.TrackAnalysePerService(0, dataList, serviceAnalyseResults)

			for k, _ := range serviceAnalyseResults {

				serviceReqCountCollector.WithLabelValues(
					serviceAnalyseResults[k].ServiceName,
					strconv.Itoa(serviceAnalyseResults[k].StatusCode),
					serviceAnalyseResults[k].ClientIP,
					serviceAnalyseResults[k].URL,
				).Set(float64(serviceAnalyseResults[k].Count))

				serviceReqAvgUseTime.WithLabelValues(
					serviceAnalyseResults[k].ServiceName,
					strconv.Itoa(serviceAnalyseResults[k].StatusCode),
					serviceAnalyseResults[k].ClientIP,
					serviceAnalyseResults[k].URL,
				).Set(float64(serviceAnalyseResults[k].AvgUsseTime))

				serviceMaxUseTime.WithLabelValues(
					serviceAnalyseResults[k].ServiceName,
					strconv.Itoa(serviceAnalyseResults[k].StatusCode),
					serviceAnalyseResults[k].ClientIP,
					serviceAnalyseResults[k].URL,
				).Set(float64(serviceAnalyseResults[k].MaxUseTime))

				//init data
				serviceAnalyseResults[k].Count = 0
				serviceAnalyseResults[k].AvgUsseTime = 0
				serviceAnalyseResults[k].MaxUseTime = 0
			}
			time.Sleep(time.Second * time.Duration(refreshInterval))

		}
	}()

	addr := flag.String("listen-address", serverHost, "The address to listen on for HTTP requests.")
	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(*addr, nil)
}
