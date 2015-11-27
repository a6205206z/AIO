package main

import (
	"./alarm"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	timeOut, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	gtTime := time.Date(year, month, day, hour, min, sec, nsec, loc)

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

}
