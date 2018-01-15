package robot

import (
	"czddz-robot/poker"
	"github.com/name5566/leaf/log"
)

func (a *Agent) handleMsg(jsonMap map[string]interface{}) {
	if _, ok := jsonMap["S2C_Heartbeat"]; ok {
		a.sendHeartbeat()
	} else if res, ok := jsonMap["S2C_Login"].(map[string]interface{}); ok {
		a.playerData.AccountID = int(res["AccountID"].(float64))
		log.Debug("accID: %v 登录", a.playerData.AccountID)
		a.playerData.Role = int(res["Role"].(float64))
		if a.playerData.Role == 1 {
			a.setUserRobot()
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
		errCode := res["Error"].(float64)
		switch errCode {
		case 4:
			log.Debug("accID: %v 需要%v张房卡才能游戏", a.playerData.AccountID, int(res["RoomCards"].(float64)))
		}
	} else if res, ok := jsonMap["S2C_EnterRoom"].(map[string]interface{}); ok {
		if int(res["Error"].(float64)) == 0 {
			a.playerData.Position = int(res["Position"].(float64))
			a.playerData.RoomType = int(res["RoomType"].(float64))
			switch a.playerData.RoomType {
			case roomBaseScoreMatching:
				log.Debug("accID: %v 进入底分匹配房", a.playerData.AccountID)
			case roomRedPacketMatching:
				log.Debug("accID: %v 进入红包匹配房", a.playerData.AccountID)

			}
			a.getAllPlayer()
		}
	} else if res, ok := jsonMap["S2C_SitDown"].(map[string]interface{}); ok {
		if a.isMe(int(res["Position"].(float64))) {
			a.playerData.analyzer = new(poker.LandlordAnalyzer)
			a.prepare()
		}
	} else if res, ok := jsonMap["S2C_UpdatePokerHands"].(map[string]interface{}); ok {
		if a.isMe(int(res["Position"].(float64))) {
			a.playerData.hands = To1DimensionalArray(res["Hands"].([]interface{}))
			log.Debug("hands: %v", poker.ToCardsString(a.playerData.hands))
		}
	} else if res, ok := jsonMap["S2C_ActionLandlordBid"].(map[string]interface{}); ok {
		if a.isMe(int(res["Position"].(float64))) {
			Delay(func() {
				a.bid()
			})
		}
	} else if res, ok := jsonMap["S2C_ActionLandlordGrab"].(map[string]interface{}); ok {
		if a.isMe(int(res["Position"].(float64))) {
			Delay(func() {
				a.playerData.analyzer.Analyze(a.playerData.hands)
				grab := false
				if a.playerData.analyzer.Bomb() {
					grab = true
				}
				a.grab(grab)
			})
		}
	} else if res, ok := jsonMap["S2C_ActionLandlordDiscard"].(map[string]interface{}); ok {
		if a.isMe(int(res["Position"].(float64))) {
			Delay(func() {
				a.systemHost()
			})
		}
	} else if _, ok := jsonMap["S2C_LandlordRoundResult"].(map[string]interface{}); ok {
		a.enterRandRoom()
	}
}

func To1DimensionalArray(array []interface{}) []int {
	var newArray []int
	for _, v := range array {
		newArray = append(newArray, int(v.(float64)))
	}
	return newArray
}
