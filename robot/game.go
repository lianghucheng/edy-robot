package robot

import (
	"czddz-robot/msg"
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
	a.writeMsg(&msg.C2S_WeChatLogin{
		Unionid:    unionids[count],
		NickName:   nicknames[count],
		Headimgurl: headimgurl[count],
	})
	count++
}

func (a *Agent) setUserRobot() {
	a.writeMsg(&msg.C2S_SetUserRobot{})
}

func (a *Agent) enterRoom() {
	a.writeMsg(&msg.C2S_EnterRoom{})
}

func (a *Agent) startMatching(roomType int, baseScore int, redPacketType int) {
	a.writeMsg(&msg.C2S_LandlordMatching{
		RoomType:      roomType,
		BaseScore:     baseScore,
		RedPacketType: redPacketType,
	})
}

func (a *Agent) getAllPlayer() {
	a.writeMsg(&msg.C2S_GetAllPlayers{})
}

func (a *Agent) prepare() {
	a.writeMsg(&msg.C2S_LandlordPrepare{
		ShowCards: false,
	})
}

func (a *Agent) bid() {
	a.writeMsg(&msg.C2S_LandlordBid{
		Bid: true,
	})
}

func (a *Agent) grab(grab bool) {
	a.writeMsg(&msg.C2S_LandlordGrab{
		Grab: grab,
	})
}

func (a *Agent) systemHost() {
	a.writeMsg(&msg.C2S_SystemHost{
		Host: true,
	})
}
