package db

import (
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
)

var MongoDB *mongodb.DialContext
var DB = "backstage"

func init() {
	var err error
	MongoDB, err = mongodb.Dial("mongodb://mongouser:eiTV^45i@10.66.129.0:27017", 100)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
