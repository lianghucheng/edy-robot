package msg

type C2S_Heartbeat struct {
}

type C2S_WeChatLogin struct {
	NickName   string
	Headimgurl string
	Unionid    string
}

type C2S_SetRobotData struct {
	LoginIP string
}

type C2S_EnterRoom struct{}

type C2S_ExitRoom struct{}

type C2S_LandlordMatching struct {
	RoomType      int // 房间类型: 1 底分匹配、4 红包匹配
	BaseScore     int // 底分: 100、5000、1万
	RedPacketType int // 红包种类(元): 1、10
}

type C2S_GetAllPlayers struct{}

type C2S_LandlordPrepare struct {
	ShowCards bool
}

type C2S_LandlordBid struct {
	Bid bool
}

type C2S_LandlordGrab struct {
	Grab bool
}

type C2S_LandlordDouble struct {
	Double bool
}

type C2S_LandlordShowCards struct {
	ShowCards bool
}

type C2S_SystemHost struct {
	Host bool
}
