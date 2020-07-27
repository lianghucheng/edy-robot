package db

import (
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
)

type RobotMatchNum struct {
	ID          int `bson:"_id"`
	MatchID     string
	MatchType   string
	MatchName   string
	PerMaxNum   int
	Total       int
	JoinNum     int
	Desc        string
	Status      int
	RobotStatus int
}

func SaveRobotJoinNum(matchid string, num int) {
	se := MongoDB.Ref()
	defer MongoDB.UnRef(se)
	data := new(RobotMatchNum)
	if err := se.DB(DB).C("robotmatchnum").Find(bson.M{"matchid": matchid}).One(data); err != nil {
		log.Error(err.Error())
		return
	}
	data.JoinNum = num
	se.DB(DB).C("robotmatchnum").Upsert(bson.M{"matchid": matchid}, data)
}
