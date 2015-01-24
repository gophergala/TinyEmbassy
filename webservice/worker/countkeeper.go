/*
* @Author: souravray
* @Date:   2015-01-25 01:02:42
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-25 02:14:27
 */

package worker

import (
	"fmt"
	//"github.com/gophergala/tinyembassy/webservice/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
	"os"
	"time"
)

func Z(t time, badge string, publisher string, campaing string, counter []string) {
	s, err := mgo.Dial(conf.DbURI)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	s.SetSafe(&mgo.Safe{})
	c := s.DB(conf.DbName).C("counter")
	fmt.Println(counter.Id)
	hour := fmt.Sprintf("t.h.%d%02d%02d%02d%02d%02d.i",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	err = c.Update(bson.M{"id": counter.Id}, bson.M{"$inc": bson.M{"ct.i": 1, hour: 1}})
	if err != nil {
		fmt.Printf("Can't update document %v\n", err)
		os.Exit(1)
	}
}
