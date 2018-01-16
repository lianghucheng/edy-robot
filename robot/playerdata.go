package robot

import (
	"czddz-robot/poker"
	"math/rand"
)

// 房间类型
const (
	roomBaseScoreMatching = 1 // 底分匹配
	roomRedPacketMatching = 4 // 红包匹配
)

var (
	roomType      = []int{roomBaseScoreMatching, roomRedPacketMatching}
	baseScore     = []int{100, 5000, 10000}
	redPacketType = []int{1, 10}
)

type PlayerData struct {
	Unionid       string
	Nickname      string
	AccountID     int
	RoomType      int // 房间类型 0 练习 1 房卡匹配 2 私人
	BaseScore     int
	RedPacketType int
	Position      int
	Role          int

	hands    []int
	analyzer *poker.LandlordAnalyzer
}

func (playerData *PlayerData) getRandRoom() {
	switch roomType[rand.Intn(len(roomType))] {
	case roomBaseScoreMatching:
		playerData.RoomType = roomBaseScoreMatching
		playerData.BaseScore = baseScore[rand.Intn(len(baseScore))]
	case roomRedPacketMatching:
		playerData.RoomType = roomRedPacketMatching
		playerData.RedPacketType = redPacketType[rand.Intn(len(redPacketType))]
	}
}
