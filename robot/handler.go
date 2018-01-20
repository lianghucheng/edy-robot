package robot

import (
	"czddz-robot/poker"
	"github.com/name5566/leaf/log"
	"math/rand"
	"time"
)

func (a *Agent) handleMsg(jsonMap map[string]interface{}) {
	if _, ok := jsonMap["S2C_Heartbeat"]; ok {
		a.sendHeartbeat()
	} else if res, ok := jsonMap["S2C_Login"].(map[string]interface{}); ok {
		a.playerData.AccountID = int(res["AccountID"].(float64))
		log.Debug("accID: %v 登录", a.playerData.AccountID)
		a.playerData.Role = int(res["Role"].(float64))
		if a.playerData.Role != -2 {
			a.setRobotData()
			return
		}
		if res["AnotherRoom"].(bool) {
			a.enterRoom()
		} else {
			if *Play {
				a.enterRandRoom()
			}
		}
	} else if res, ok := jsonMap["S2C_CreateRoom"].(map[string]interface{}); ok {
		err := res["Error"].(float64)
		switch err {
		case 4:
			log.Debug("accID: %v 需要%v张房卡才能游戏", a.playerData.AccountID, res["RoomCards"].(float64))
		}
	} else if res, ok := jsonMap["S2C_EnterRoom"].(map[string]interface{}); ok {
		err := res["Error"].(float64)
		switch err {
		case 0:
			a.playerData.Position = int(res["Position"].(float64))
			a.playerData.RoomType = int(res["RoomType"].(float64))
			roomNumber := res["RoomNumber"].(string)
			switch a.playerData.RoomType {
			case roomBaseScoreMatching:
				a.playerData.BaseScore = int(res["BaseScore"].(float64))
				log.Debug("accID: %v 进入房间:%v 底分: %v", a.playerData.AccountID, roomNumber, a.playerData.BaseScore)
			case roomRedPacketMatching:
				a.playerData.RedPacketType = int(res["RedPacketType"].(float64))
				log.Debug("accID: %v 进入房间:%v 红包: %v", a.playerData.AccountID, roomNumber, a.playerData.RedPacketType)
			}
			a.getAllPlayer()
		case 6:
			log.Debug("accID: %v 需要%v筹码才能进入", a.playerData.AccountID, res["MinChips"].(float64))
			Delay(func() {
				a.enterRandRoom()
			})
		case 7:
			log.Debug("accID: %v 比赛暂未开始", a.playerData.AccountID)
			Delay(func() {
				a.enterRandRoom()
			})
		}
	} else if res, ok := jsonMap["S2C_SitDown"].(map[string]interface{}); ok {
		if a.isMe(int(res["Position"].(float64))) {
			a.prepare()
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
				a.playerData.analyzer.Analyze(a.playerData.hands)
				if a.playerData.analyzer.HasKingBomb || a.playerData.analyzer.HasBomb {
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
	} else if res, ok := jsonMap["S2C_ActionLandlordDouble"].(map[string]interface{}); ok {
		if a.isMe(int(res["Position"].(float64))) {
			Delay(func() {
				a.double(false)
			})
		}
	} else if res, ok := jsonMap["S2C_ActionLandlordShowCards"].(map[string]interface{}); ok {
		if a.isMe(int(res["Position"].(float64))) {
			Delay(func() {
				a.showCards(false)
			})
		}
	} else if res, ok := jsonMap["S2C_ActionLandlordDiscard"].(map[string]interface{}); ok {
		if a.isMe(int(res["Position"].(float64))) {
			Delay(func() {
				a.systemHost()
			})
		}
	} else if _, ok := jsonMap["S2C_LandlordRoundResult"].(map[string]interface{}); ok {
		Delay(func() {
			a.enterRandRoom()
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
