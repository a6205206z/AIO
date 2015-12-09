package data

import (
	"crypto/sha1"
	"fmt"
	"io"
)

type AnalyseResult struct {
	ServiceName string
	StatusCode  int
	ClientIP    string
	URL         string
	Count       int
	MaxUseTime  int
	AvgUseTime  int
}

func TrackAnalysePerService(timeout int, trackingList []Tracking, analyseResults map[string]*AnalyseResult) {
	for i := 0; i < len(trackingList); i++ {
		analyseKey := hash(fmt.Sprintf("%s%d%s%s", trackingList[i].Appname, trackingList[i].Statuscode, trackingList[i].Realclientip, trackingList[i].Url))
		if timeout < trackingList[i].Usetime {
			if analyseResults[analyseKey] == nil {
				analyseResults[analyseKey] = new(AnalyseResult)
				analyseResults[analyseKey].ServiceName = trackingList[i].Appname
				analyseResults[analyseKey].StatusCode = trackingList[i].Statuscode
				analyseResults[analyseKey].ClientIP = trackingList[i].Realclientip
				analyseResults[analyseKey].URL = trackingList[i].Url
			}

			analyseResults[analyseKey].Count++
			if analyseResults[analyseKey].MaxUseTime < trackingList[i].Usetime {
				analyseResults[analyseKey].MaxUseTime = trackingList[i].Usetime
			}
			if trackingList[i].Usetime > 0 {
				analyseResults[analyseKey].AvgUseTime = (trackingList[i].Usetime + analyseResults[analyseKey].AvgUseTime) / 2
			}
		}
	}
}

func hash(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
