package robot

import (
	"math/rand"
	"time"
)

// 房间类型
const (
	roomBaseScoreMatching = 1 // 底分匹配
	roomRedPacketMatching = 4 // 红包匹配
)

const (
	roleRobot = -2 // 机器人
)

var (
	roomType = []int{roomBaseScoreMatching, roomRedPacketMatching}
	// baseScore     = []int{100, 5000, 10000}
	baseScore     = []int{100}
	redPacketType = []int{1, 10}
)

type PlayerData struct {
	Unionid       string
	Nickname      string
	AccountID     int
	RoomType      int // 房间类型: 0 练习、1 底分匹配、4 红包匹配
	MaxPlayers    int
	BaseScore     int
	RedPacketType int
	Position      int
	Role          int

	hands         []int
	exitRoomTimer *time.Timer
}

func (playerData *PlayerData) getRandRedPacketMatchingRoom() {
	playerData.RoomType = roomRedPacketMatching
	playerData.RedPacketType = redPacketType[rand.Intn(len(redPacketType))]
}

func (playerData *PlayerData) getRandBaseScoreMatchingRoom() {
	playerData.RoomType = roomBaseScoreMatching
	playerData.BaseScore = baseScore[rand.Intn(len(baseScore))]
}

//func (playerData *PlayerData) getRandRoom() {
//	switch roomType[rand.Intn(len(roomType))] {
//	case roomBaseScoreMatching:
//		playerData.RoomType = roomBaseScoreMatching
//		playerData.BaseScore = baseScore[rand.Intn(len(baseScore))]
//	case roomRedPacketMatching:
//		playerData.RoomType = roomRedPacketMatching
//		playerData.RedPacketType = redPacketType[rand.Intn(len(redPacketType))]
//	}
//}
