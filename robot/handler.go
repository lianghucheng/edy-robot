package robot

import (
	"czddz-robot/poker"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/timer"
	"math/rand"
	"strconv"
	"time"
)

func (a *Agent) handleMsg(jsonMap map[string]interface{}) {
	if _, ok := jsonMap["S2C_Heartbeat"]; ok {
		a.sendHeartbeat()
	} else if res, ok := jsonMap["S2C_Login"].(map[string]interface{}); ok {
		a.playerData.AccountID = int(res["AccountID"].(float64))
		a.playerData.Role = int(res["Role"].(float64))
		if a.playerData.Role != roleRobot {
			log.Debug("accID: %v 登录初始化", a.playerData.AccountID)
			a.setRobotData()
			return
		}
		if res["AnotherRoom"].(bool) {
			a.enterRoom()
		} else {
			if *Play {
				index, _ := strconv.Atoi(a.playerData.Unionid)
				if index > 99 {
					a.enterRedPacketMatchingRoom()

					CronFunc("10 0 12 * * *", func() {
						a.enterRedPacketMatchingRoom()
					})
					CronFunc("10 0 20 * * *", func() {
						a.enterRedPacketMatchingRoom()
					})
				} else {
					//delayTime, _ := strconv.Atoi(a.playerData.Unionid)
					//DelayDo(time.Duration(delayTime+10)*time.Second, a.enterBaseScoreMatchingRoom)
					DelayDo(time.Duration(10)*time.Second, a.enterBaseScoreMatchingRoom)
				}
			} else {
				log.Debug("accID: %v 登录", a.playerData.AccountID)
			}
		}
	} else if res, ok := jsonMap["S2C_EnterRoom"].(map[string]interface{}); ok {
		err := int(res["Error"].(float64))
		switch err {
		case 0:
			a.playerData.Position = int(res["Position"].(float64))
			a.playerData.RoomType = int(res["RoomType"].(float64))
			a.playerData.MaxPlayers = int(res["MaxPlayers"].(float64))
			roomNumber := res["RoomNumber"].(string)
			switch a.playerData.RoomType {
			case roomBaseScoreMatching:
				a.playerData.BaseScore = int(res["BaseScore"].(float64))
				log.Debug("accID: %v 进入房间: %v 底分: %v", a.playerData.AccountID, roomNumber, a.playerData.BaseScore)
			case roomRedPacketMatching:
				a.playerData.RedPacketType = int(res["RedPacketType"].(float64))
				log.Debug("accID: %v 进入房间: %v 红包: %v", a.playerData.AccountID, roomNumber, a.playerData.RedPacketType)
			}
			a.getAllPlayer()
		case 4:
			// S2C_EnterRoom_Unknown
			// 机器人进入房间不会创建，如果没有一人房或者两人房就返回这条错误
			// 延迟进入
			//log.Debug("无房间可进")
			//delayTime, _ := strconv.Atoi(a.playerData.Unionid)
			DelayDo(time.Duration(10)*time.Second, a.enterBaseScoreMatchingRoom)
		case 6: // S2C_EnterRoom_LackOfChips
			log.Debug("accID: %v 需要%v筹码才能进入", a.playerData.AccountID, res["MinChips"].(float64))
			a.addChips()
		case 7: // S2C_EnterRoom_NotRightNow
			log.Debug("accID: %v 红包比赛暂未开始", a.playerData.AccountID)
		}
	} else if res, ok := jsonMap["S2C_SitDown"].(map[string]interface{}); ok {
		pos := int(res["Position"].(float64))
		//if pos == a.playerData.MaxPlayers-1 {
		if pos == a.playerData.Position {
			Delay(func() {
				a.prepare()
			})
		}
	} else if res, ok := jsonMap["S2C_ExitRoom"].(map[string]interface{}); ok {
		err := int(res["Error"].(float64))
		switch err {
		case 0:
			pos := int(res["Position"].(float64))
			if a.isMe(pos) {
				Delay(func() {
					a.enterTheRoom()
				})
			} else if pos == a.playerData.MaxPlayers-1 {
				StopTimer(a.playerData.exitRoomTimer)
				a.exitRoom()
			}
		}
	} else if res, ok := jsonMap["S2C_Prepare"].(map[string]interface{}); ok {
		pos := int(res["Position"].(float64))
		//log.Debug("当前发送位置：%v 自己位置：%v", pos, a.playerData.Position)
		ready := res["Ready"].(bool)
		if a.isMe(pos) && ready {
			duration := 10 * time.Minute
			// duration := 5 * time.Second
			a.playerData.exitRoomTimer = time.AfterFunc(duration, func() {
				a.playerData.exitRoomTimer = nil
				a.exitRoom()
			})
		} else if pos == a.playerData.MaxPlayers-1 && ready {
			StopTimer(a.playerData.exitRoomTimer)
		}
	} else if res, ok := jsonMap["S2C_UpdatePokerHands"].(map[string]interface{}); ok {
		if a.isMe(int(res["Position"].(float64))) {
			if res["Hands"] != nil {
				a.playerData.hands = To1DimensionalArray(res["Hands"].([]interface{}))
				log.Debug("hands: %v", poker.ToCardsString(a.playerData.hands))
			}
		}
	} else if res, ok := jsonMap["S2C_ActionLandlordBid"].(map[string]interface{}); ok {
		if a.isMe(int(res["Position"].(float64))) {
			Delay(func() {
				analyzer := new(poker.LandlordAnalyzer)
				analyzer.Analyze(a.playerData.hands)
				if analyzer.HasKingBomb || analyzer.HasBomb {
					a.bid(true)
				} else {
					a.bid(false)
				}
			})
		}
	} else if res, ok := jsonMap["S2C_ActionLandlordGrab"].(map[string]interface{}); ok {
		if a.isMe(int(res["Position"].(float64))) {
			Delay(func() {
				a.grab(false)
			})
		}
	} else if _, ok := jsonMap["S2C_ActionLandlordDouble"].(map[string]interface{}); ok {
		Delay(func() {
			a.double(false)
		})
	} else if _, ok := jsonMap["S2C_ActionLandlordShowCards"].(map[string]interface{}); ok {
		Delay(func() {
			a.showCards(false)
		})
	} else if res, ok := jsonMap["S2C_ActionLandlordDiscard"].(map[string]interface{}); ok {
		if a.isMe(int(res["Position"].(float64))) {
			Delay(func() {
				a.systemHost()
			})
		}
	} else if _, ok := jsonMap["S2C_LandlordRoundResult"].(map[string]interface{}); ok {
		Delay(func() {
			a.enterTheRoom()
		})
	} else if res, ok := jsonMap["S2C_PayOK"].(map[string]interface{}); ok {
		log.Debug("accID: %v 加钱: %v", a.playerData.AccountID, res["Chips"].(float64))
		Delay(func() {
			a.enterTheRoom()
		})
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
	time.AfterFunc(time.Duration(rand.Intn(2)+3)*time.Second, func() {
		if cb != nil {
			cb()
		}
	})
}

func DelayDo(d time.Duration, cb func()) {
	if cb == nil {
		return
	}
	time.AfterFunc(d, cb)
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
