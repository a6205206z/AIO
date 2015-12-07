package data

type AnalyseResult struct {
	Count      int
	MaxUseTime int
}

func TrackTimeoutAnalysePerAPI(timeout int, trackingList []Tracking, analyseResults map[string]*AnalyseResult) {

	for i := 0; i < len(trackingList); i++ {
		analyseKey := trackingList[i].Appname + "-" + trackingList[i].Location
		if timeout < trackingList[i].Usetime {
			if analyseResults[analyseKey] == nil {
				analyseResults[analyseKey] = new(AnalyseResult)
			}
			analyseResults[analyseKey].Count++
			if analyseResults[analyseKey].MaxUseTime < trackingList[i].Usetime {
				analyseResults[analyseKey].MaxUseTime = trackingList[i].Usetime
			}
		}
	}
}

func TrackTimeoutAnalysePerService(timeout int, trackingList []Tracking, analyseResults map[string]*AnalyseResult) {

	for i := 0; i < len(trackingList); i++ {
		analyseKey := trackingList[i].Appname
		if timeout < trackingList[i].Usetime {
			if analyseResults[analyseKey] == nil {
				analyseResults[analyseKey] = new(AnalyseResult)
			}
			analyseResults[analyseKey].Count++
			if analyseResults[analyseKey].MaxUseTime < trackingList[i].Usetime {
				analyseResults[analyseKey].MaxUseTime = trackingList[i].Usetime
			}
		}
	}
}
