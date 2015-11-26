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

	gtTime := time.Now()

	for {
		dataList, _ := alarm.LoadTrackingData("192.168.200.22", &gtTime)
		analyseResults := alarm.TrackTimeoutAnalyse(timeOut, dataList)

		for k, _ := range analyseResults {
			fmt.Printf("[%s request time out] count= %d max use time = %d\n", k, analyseResults[k].Count, analyseResults[k].MaxUseTime)
		}

		fmt.Printf("------------------%s----------------------\n", gtTime.String())
		time.Sleep(time.Second * 10)
	}

}
