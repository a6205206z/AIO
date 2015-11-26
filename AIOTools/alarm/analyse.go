package alarm

type AnalyseResult struct {
	Count      int
	MaxUseTime int
}

func TrackTimeoutAnalyse(timeout int, trackingList []Tracking) map[string]*AnalyseResult {
	analyseResults := make(map[string]*AnalyseResult)

	for i := 0; i < len(trackingList); i++ {
		appname := trackingList[i].Appname
		if timeout < trackingList[i].Usetime {
			if analyseResults[appname] == nil {
				analyseResults[appname] = new(AnalyseResult)
			}
			analyseResults[appname].Count++
			if analyseResults[appname].MaxUseTime < trackingList[i].Usetime {
				analyseResults[appname].MaxUseTime = trackingList[i].Usetime
			}
		}
	}

	return analyseResults
}
