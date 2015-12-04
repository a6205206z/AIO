package alarm

type AnalyseResult struct {
	Count      int
	MaxUseTime int
}

func TrackTimeoutAnalysePerAPI(timeout int, trackingList []Tracking) map[string]*AnalyseResult {
	analyseResults := make(map[string]*AnalyseResult)

	for i := 0; i < len(trackingList); i++ {
		analyseKey := trackingList[i].Appname + "-" + trackingList[i].Url
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

	return analyseResults
}

func TrackTimeoutAnalysePerService(timeout int, trackingList []Tracking) map[string]*AnalyseResult {
	analyseResults := make(map[string]*AnalyseResult)

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

	return analyseResults
}
