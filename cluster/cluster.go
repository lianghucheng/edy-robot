package cluster

import (
	"edy-robot/conf"
	"fmt"
	"github.com/name5566/leaf/log"
	"net/http"
	"strconv"
	"sync"
)

var Mux *sync.Mutex
var RobotJoinNum = make(map[string]int)
var RobotUseNum = make(map[string]int)

func init() {
	Mux = &sync.Mutex{}
	go startHttp()
}

func startHttp() {
	http.HandleFunc("/conf/num-limit", handleConfNumLimit)
	http.HandleFunc("/conf/robot-status", handleRobotStatus)
	http.HandleFunc("/", handleIndex)
	log.Debug("启动http服务")
	if err := http.ListenAndServe(":9086", nil); err != nil {
		log.Error(err.Error())
		log.Debug("服务启动失败")
		return
	}
}

func handleConfNumLimit(w http.ResponseWriter, r *http.Request) {
	total := r.FormValue("robot_total")
	matchid := r.FormValue("matchid")
	robotTotal, _ := strconv.Atoi(total)
	log.Debug("接收handleConfNumLimit   %v  %v", total, matchid)
	Mux.Lock()
	defer Mux.Unlock()
	if _, ok := conf.GetConfMatchidRobot()[matchid]; ok {
		log.Debug("更新赛事机器人zong数量配置，配置前数量：%v", conf.GetConfMatchidRobot()[matchid].Total)
		conf.GetConfMatchidRobot()[matchid].Total = robotTotal
	} else {
		log.Debug("更新赛事机器人zong数量配置，配置前没有数量")
		conf.GetConfMatchidRobot()[matchid] = &conf.ConfMatchidRobot{
			Total: robotTotal,
		}
	}
	log.Debug("更新赛事机器人zong数量配置，配置后数量：%v    matchid:%v", conf.GetConfMatchidRobot()[matchid].Total, matchid)
	//todo: Save DB
	w.Write([]byte(fmt.Sprintf("更新赛事机器人zong数量配置，配置后数量：%v", conf.GetConfMatchidRobot()[matchid].Total)))
}

func handleRobotStatus(w http.ResponseWriter, r *http.Request) {
	status := r.FormValue("status")
	matchid := r.FormValue("matchid")
	robotStatus, _ := strconv.Atoi(status)
	Mux.Lock()
	defer Mux.Unlock()
	if _, ok := conf.GetConfMatchidRobot()[matchid]; ok {
		log.Debug("更新机器人状态配置，配置前状态：%v", conf.GetConfMatchidRobot()[matchid].Status)
		conf.GetConfMatchidRobot()[matchid].Status = robotStatus
	} else {
		log.Debug("更新机器人状态配置，配置前没有状态")
		conf.GetConfMatchidRobot()[matchid] = &conf.ConfMatchidRobot{
			Status: robotStatus,
		}
	}
	log.Debug("更新机器人状态配置，配置前状态：%v", conf.GetConfMatchidRobot()[matchid].Status)
	//todo: Save DB
	w.Write([]byte(fmt.Sprintf("更新机器人状态配置，配置前状态：%v", conf.GetConfMatchidRobot()[matchid].Status)))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf(`当前系统并发状态,\n机器人使用个数:%v\n机器人参赛个数：%v，赛事相关机器人情况:%+v`, RobotUseNum, RobotJoinNum, conf.GetConfMatchidRobot)))
}
