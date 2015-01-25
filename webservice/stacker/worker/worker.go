/*
* @Author: souravray
* @Date:   2015-01-25 05:43:22
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-25 05:50:51
 */

package worker

import (
	"fmt"
	//"github.com/gophergala/tinyembassy/webservice/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
	"time"
)

type Payload struct {
	T          time.Time
	Counter    []string
	Badge      string
	Publisher  string
	DBURI      string
	DBName     string
	Campaining bson.ObjectId
}

func Z(payload Payload) {
	s, err := mgo.Dial(payload.DBURI)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	t := payload.T
	s.SetSafe(&mgo.Safe{})
	hour := fmt.Sprintf("t.h.%d%02d%02d%02d%02d%02d.i",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	c := s.DB(payload.DBName).C("counter")
	err = c.Update(bson.M{"id": payload.Campaining}, bson.M{"$inc": bson.M{"ct.i": 1, hour: 1}})
	if err != nil {
		fmt.Printf("Can't update document %v\n", err)
		os.Exit(1)
	}
}
