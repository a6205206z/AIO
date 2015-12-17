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
		[]string{"service_name", "status_code", "client_ip"},
	)

	serviceReqAvgUseTime = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_avg_use_time",
			Help: "service avg use time",
		},
		[]string{"service_name", "status_code", "client_ip"},
	)

	serviceMaxUseTime = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_max_use_time",
			Help: "service max use time",
		},
		[]string{"service_name", "status_code", "client_ip"},
	)

	serviceUseTimeDistributions = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "service_use_time_distributions",
			Help:    "service use time distributions",
			Buckets: prometheus.LinearBuckets(0, 100000, 11),
		},
		[]string{"service_name", "status_code"},
	)

	serviceAnalyseResults = make(map[string]*data.AnalyseResult)
)

func init() {
	prometheus.MustRegister(serviceReqCountCollector)
	prometheus.MustRegister(serviceReqAvgUseTime)
	prometheus.MustRegister(serviceMaxUseTime)
	prometheus.MustRegister(serviceUseTimeDistributions)
}

func resetCollector() {
	//init collector
	serviceReqCountCollector.Reset()
	serviceReqAvgUseTime.Reset()
	serviceMaxUseTime.Reset()
	serviceUseTimeDistributions.Reset()
}

func StartServer(serverHost string, dbHost string, refreshInterval int) {
	flag.Parse()

	go func() {
		now := time.Now()
		gtTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), data.Location)

		for {
			resetCollector()
			dataList, err := data.LoadTrackingData(dbHost, &gtTime)
			if err != nil {
				log.Println(err)
				continue
			}
			for i := 0; i < len(dataList); i++ {
				serviceUseTimeDistributions.WithLabelValues(
					dataList[i].Appname,
					strconv.Itoa(dataList[i].Statuscode),
				).Observe(float64(dataList[i].Usetime))
			}

			data.TrackAnalysePerService(0, dataList, serviceAnalyseResults)

			for k, _ := range serviceAnalyseResults {

				serviceReqCountCollector.WithLabelValues(
					serviceAnalyseResults[k].ServiceName,
					strconv.Itoa(serviceAnalyseResults[k].StatusCode),
					serviceAnalyseResults[k].ClientIP,
				).Set(float64(serviceAnalyseResults[k].Count))

				serviceReqAvgUseTime.WithLabelValues(
					serviceAnalyseResults[k].ServiceName,
					strconv.Itoa(serviceAnalyseResults[k].StatusCode),
					serviceAnalyseResults[k].ClientIP,
				).Set(float64(serviceAnalyseResults[k].AvgUseTime))

				serviceMaxUseTime.WithLabelValues(
					serviceAnalyseResults[k].ServiceName,
					strconv.Itoa(serviceAnalyseResults[k].StatusCode),
					serviceAnalyseResults[k].ClientIP,
				).Set(float64(serviceAnalyseResults[k].MaxUseTime))

				//init data
				serviceAnalyseResults[k].Count = 0
				serviceAnalyseResults[k].AvgUseTime = 0
				serviceAnalyseResults[k].MaxUseTime = 0
			}
			time.Sleep(time.Second * time.Duration(refreshInterval))
		}
	}()

	addr := flag.String("listen-address", serverHost, "The address to listen on for HTTP requests.")
	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(*addr, nil)
}
