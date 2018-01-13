package robot

import "czddz-robot/poker"

type PlayerData struct {
	Unionid   string
	Nickname  string
	AccountID int
	RoomType  int // 房间类型 0 练习 1 房卡匹配 2 私人
	Position  int
	Role      int

	hands    []int
	analyzer *poker.LandlordAnalyzer
}
