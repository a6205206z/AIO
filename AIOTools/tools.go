package main

import (
	_ "./alarm"
	"./monitor"
	_ "fmt"
	"os"
	_ "strconv"
	_ "time"
)

func main() {
	/*
		timeOut, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		now := time.Now()
		gtTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), alarm.Location)
	*/
	/*
		email, err := alarm.LoadServiceRegisterUserEmail("uoko")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(email)
	*/
	/*
		for {
			dataList, err := alarm.LoadTrackingData("121.41.0.138", &gtTime)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			analyseResults := alarm.TrackTimeoutAnalyse(timeOut, dataList)

			for k, _ := range analyseResults {
				fmt.Printf("[%s request time out] count= %d max use time = %d\n", k, analyseResults[k].Count, analyseResults[k].MaxUseTime)
			}

			fmt.Printf("------------------%s----------------------\n", gtTime.String())
			time.Sleep(time.Minute)
		}
	*/
	dbHost := os.Args[1]
	monitor.StartServer(dbHost)
}
