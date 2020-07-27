package db

import (
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
)

var MongoDB *mongodb.DialContext
var DB = "backstage"

func init() {
	var err error
	MongoDB, err = mongodb.Dial("mongodb://localhost", 100)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
