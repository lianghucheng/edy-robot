package robot

import (
	"edy-robot/msg"
	"edy-robot/poker"
	"encoding/json"
	"fmt"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/timer"
	"math/rand"
	"time"
)

func (a *Agent) handleMsg(jsonMap map[string]interface{}) {
	if _, ok := jsonMap["S2C_Heartbeat"]; ok {
		a.sendHeartbeat()
	} else if res, ok := jsonMap["S2C_Login"].(map[string]interface{}); ok {
		a.playerData.AccountID = int(res["AccountID"].(float64))
		log.Debug("登录成功！%v", a.playerData.AccountID)
		a.playerData.Role = int(res["Role"].(float64))
		if a.playerData.Role != roleRobot {
			log.Debug("accID: %v 登录初始化，因为不是机器人。", a.playerData.AccountID)
			//todo: Set character to robot.
		}
		if len(a.matchids) == 0 {
			a.writeMsg(&msg.C2S_RaceInfo{})
			Delay(func() {
				if a != nil {
					a.signIn()
				}
			})
		}
	} else if res, ok := jsonMap["S2C_UpdatePokerHands"].(map[string]interface{}); ok {
		m := new(msg.S2C_UpdatePokerHands)
		if parseObject(res, m) {
			if a.isMe(m.Position) {
				a.playerData.hands = m.Hands
				log.Debug("hands: %v", poker.ToCardsString(a.playerData.hands))
			}
		}
	} else if res, ok := jsonMap["S2C_RaceInfo"]; ok {
		log.Debug("收到赛事信息")
		msgRaceinfo := new(msg.S2C_RaceInfo)
		if parseObject(res, msgRaceinfo) {
			a.matchids = []string{}
			for _, race := range msgRaceinfo.Races {
				a.matchids = append(a.matchids, race.ID)
			}
		}
	} else if res, ok := jsonMap["S2C_Apply"]; ok {
		m := new(msg.S2C_Apply)
		if parseObject(res, m) {
			if m.Error == 0 && m.Action == 1 {
				log.Debug("报名成功")
				a.currMatchid = m.RaceID
			} else if m.Error != 0 {
				log.Debug("%v", m.Error)
				if a.signOutTimer != nil {
					a.signOutTimer.Stop()
					a.signOutTimer = nil
				}
				Delay(a.signIn)
			}
		}
	} else if res, ok := jsonMap["S2C_EnterRoom"]; ok {
		log.Debug("准备开始比赛")
		msgEnterRoom := new(msg.S2C_EnterRoom)
		if parseObject(res, msgEnterRoom) {
			if msgEnterRoom.Error == 0 {
				a.playerData.Position = msgEnterRoom.Position
			}
		}
		a.playerData.isPlay = true
		a.currMatchid = ""
	} else if res, ok := jsonMap["S2C_ActionLandlordBid"]; ok {
		m := new(msg.S2C_ActionLandlordBid)
		if parseObject(res, m) {
			if a.isMe(m.Position) {
				log.Debug("叫分")
				a.doBid(m.Score)
			}
		}
	} else if _, ok := jsonMap["S2C_ActionLandlordDouble"]; ok {
		log.Debug("加倍")
		a.doDouble()
	} else if _, ok := jsonMap["S2C_LandlordRoundFinalResult"]; ok {
		log.Debug("打完了一局")
		fmt.Println("打完了一局")
	} else if _, ok := jsonMap["S2C_MineRoundRank"]; ok {
		log.Debug("比赛结束")
		a.playerData.isPlay = false
		time.AfterFunc(10 * time.Second, func() {
			if a != nil {
				a.signIn()
			}
		})
	} else if res, ok := jsonMap["S2C_ActionLandlordDiscard"]; ok {
		m := new(msg.S2C_ActionLandlordDiscard)
		if parseObject(res, m) {
			if a.isMe(m.Position) {
				log.Debug("出牌动作")
				a.playerData.Hint = m.Hint
				a.doDiscard()
			}
		}
	}
}

func To1DimensionalArray(array []interface{}) []int {
	var newArray []int
	for _, v := range array {
		newArray = append(newArray, int(v.(float64)))
	}
	return newArray
}

func Delay(cb func()) {
	if cb != nil {
		time.AfterFunc(time.Duration(rand.Intn(2)+3)*time.Second, cb)
	}
}

func DelayDo(d time.Duration, cb func()) {
	if cb != nil {
		time.AfterFunc(d, cb)
	}
}

func CronFunc(expr string, cb func()) {
	cronExpr, _ := timer.NewCronExpr(expr)
	dispatcher.CronFunc(cronExpr, func() {
		if cb != nil {
			cb()
		}
	})
}

func StopTimer(t *time.Timer) {
	if t != nil {
		t.Stop()
		t = nil
	}
}

func parseObject(msg, obj interface{}) bool {
	b, err := json.Marshal(msg)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	err = json.Unmarshal(b, obj)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	return true
}