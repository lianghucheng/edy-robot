package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
)

var conf *RobotConf
var baseConf *BaseConf

type BaseConf struct {
	DBName       string
	DBAddr       string
	DBConnNum    int
	DBConfColl   string
	Model        string
	RobotNum     int
	GameWsAddr   string
	IpPath       string
	NicknamePath string
}

type RobotConf struct {
	ConfMatchidRobots map[string]*ConfMatchidRobot
}

type ConfMatchidRobot struct {
	Total  int
	Status int
}

func init() {
	conf = new(RobotConf)
	conf.ConfMatchidRobots = make(map[string]*ConfMatchidRobot) //todo:不初始化会不会报错
	baseConf = new(BaseConf)
	ReadBaseConf()
}

func GetConfMatchidRobot() map[string]*ConfMatchidRobot {
	return conf.ConfMatchidRobots
}

func ReadBaseConf() {
	b, err := ioutil.ReadFile("conf/base.json")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	err = json.Unmarshal(b, baseConf)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//util.PrintObject(baseConf)
}

func GetBaseConf() *BaseConf {
	return baseConf
}
