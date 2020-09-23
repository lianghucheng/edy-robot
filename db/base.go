package db

import (
	"edy-robot/conf"
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
)

var MongoDB *mongodb.DialContext
var DB = "backstage"

func init() {
	var err error
	baseConf := conf.GetBaseConf()
	log.Error("数据可地址:%v", baseConf.DBAddr)
	MongoDB, err = mongodb.Dial(baseConf.DBAddr, baseConf.DBConnNum)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
