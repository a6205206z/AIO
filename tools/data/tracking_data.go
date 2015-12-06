package data

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	LCAOTION_STR = "Asia/Chongqing"
)

var (
	Location *time.Location
)

func init() {
	Location, _ = time.LoadLocation(LCAOTION_STR)
}

type Tracking struct {
	Appname      string
	Location     string
	Url          string
	Clientip     string
	Realclientip string
	Statuscode   int
	Usetime      int
	Invoketime   time.Time
}

func LoadTrackingData(url string, gtTime *time.Time) ([]Tracking, error) {
	session, err := mgo.Dial(url)
	defer func() {
		if session != nil {
			session.Close()
		}
	}()
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	tracking_collection := session.DB("vesb").C("tracking")
	dataList := loadTrackingDataByTime(tracking_collection, gtTime)

	return dataList, err
}

func loadTrackingDataByTime(collection *mgo.Collection, lastCheckTime *time.Time) []Tracking {

	var results []Tracking
	iter := collection.Find(bson.M{"invoketime": bson.M{"$gt": lastCheckTime}}).Iter()
	defer func() {
		if iter != nil {
			iter.Close()
		}
	}()

	tracking := new(Tracking)
	for iter.Next(tracking) {
		results = append(results, *tracking)
		*lastCheckTime = tracking.Invoketime
	}
	return results
}
