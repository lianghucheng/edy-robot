package robot

import (
	"edy-robot/cluster"
	"edy-robot/conf"
	"edy-robot/msg"
	"github.com/name5566/leaf/log"
	"math/rand"
	"time"
)

func (a *Agent) isMe(pos int) bool {
	return a.playerData.Position == pos
}

func (a *Agent) sendHeartbeat() {
	a.writeMsg(&msg.C2S_Heartbeat{})
}

func (a *Agent) wechatLogin() {
	mu.Lock()
	defer mu.Unlock()
	a.playerData.Unionid = unionids[count]
	a.playerData.Nickname = nicknames[count]
	a.writeMsg(&msg.C2S_UsrnPwdLogin{
		//Username: unionids[count],
		Username: "test77",
		Password: "123456789",
	})
	count++
}

func (a *Agent) enterRoom() {
	a.writeMsg(&msg.C2S_EnterRoom{})
}

func (a *Agent) signIn() {
	log.Debug("是否已报名：%v    是否已在比赛：%v", a.isSign(), a.playerData.isPlay)
	if !a.isSign() && !a.playerData.isPlay {
		log.Debug("报名")
		if len(a.matchids) <= 0 {
			log.Debug("当前赛事个数：%v", len(a.matchids))
			time.AfterFunc(10*time.Second, a.signIn)
			return
		}
		matchid := a.matchids[rand.Intn(len(a.matchids))]
		data, ok := conf.GetConfMatchidRobot()[matchid]
		if !ok {
			log.Debug("该赛事尚未配置   %v", matchid)
			time.AfterFunc(10*time.Second, a.signIn)
			return
		}
		if data.Total <= 0 {
			log.Debug("该赛事使用的机器人为零，数量是：%v", data.Total)
			time.AfterFunc(10*time.Second, a.signIn)
			return
		}
		cluster.Mux.Lock()
		flag := 0
		if val, ok := cluster.RobotUseNum[matchid]; ok && val > data.Total {
			log.Debug("该赛事使用的机器人已满，数量是：%v", data.Total)
			time.AfterFunc(10*time.Second, a.signIn)
			flag = 1
		}
		cluster.Mux.Unlock()
		if flag == 1 {
			return
		}
		if data.Status == 1 {
			log.Debug("该赛事的机器人处于关闭状态，当前状态是：%v", data.Status)
			time.AfterFunc(10*time.Second, a.signIn)
			return
		}
		a.writeMsg(&msg.C2S_Apply{
			MatchId: matchid,
			Action:  1,
		})
		a.robotMem = matchid
		cluster.Mux.Lock()
		cluster.RobotUseNum[matchid]++
		log.Debug("RobotUseNum增加 %v", cluster.RobotUseNum[matchid])
		cluster.Mux.Unlock()
		a.signOutTimer = time.AfterFunc(10*time.Second, func() {
			if a != nil {
				a.signOut()
			}
		})
	}
}

func (a *Agent) signOut() {
	if a.isSign() && !a.playerData.isPlay {
		log.Debug("退签")
		matchid := a.robotMem
		_, ok := conf.GetConfMatchidRobot()[matchid]
		if !ok {
			log.Debug("异常情况")
			return
		}
		cluster.Mux.Lock()
		if cluster.RobotUseNum[matchid] > 0 {
			cluster.RobotUseNum[matchid]--
			log.Debug("RobotUseNum减少 %v", cluster.RobotUseNum[matchid])
		}
		cluster.Mux.Unlock()
		a.writeMsg(&msg.C2S_Apply{
			MatchId: a.currMatchid,
			Action:  2,
		})
		a.currMatchid = ""
		time.AfterFunc(10*time.Second, func() {
			if a != nil {
				a.signIn()
			}
		})
	}
}

func (a *Agent) doBid(scores []int) {
	a.writeMsg(&msg.C2S_LandlordBid{
		Score: scores[rand.Intn(len(scores))],
	})
}

func (a *Agent) doDouble() {
	temp := rand.Intn(2)
	double := false
	if temp&1 == 0 {
		double = true
	}
	a.writeMsg(&msg.C2S_LandlordDouble{
		Double: double,
	})
}

func (a *Agent) doDiscard(actionDiscardType int) {
	if len(a.playerData.Hint) >= 1 {
		Delay(func() {
			a.writeMsg(&msg.C2S_LandlordDiscard{
				Cards: a.playerData.Hint[rand.Intn(len(a.playerData.Hint))],
			})
		})
	} else {
		if actionDiscardType == 1 {
			Delay(func() {
				a.writeMsg(&msg.C2S_LandlordDiscard{
					Cards: []int{a.playerData.hands[rand.Intn(len(a.playerData.hands))]},
				})
			})
			return
		}
		Delay(func() {
			a.writeMsg(&msg.C2S_LandlordDiscard{})
		})
	}
}

func (a *Agent) isSign() bool {
	return a.currMatchid != ""
}
