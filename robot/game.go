package robot

import (
	"czddz-robot/msg"
	"math/rand"
	"strconv"
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
		Headimgurl: headimgurls[count],
	})
	count++
}

func (a *Agent) setRobotData() {
	index, _ := strconv.Atoi(a.playerData.Unionid)
	a.writeMsg(&msg.C2S_SetRobotData{
		LoginIP: loginIPs[index],
	})
}

func (a *Agent) enterRoom() {
	a.writeMsg(&msg.C2S_EnterRoom{})
}

func (a *Agent) exitRoom() {
	a.writeMsg(&msg.C2S_ExitRoom{})
}

func (a *Agent) enterBaseScoreMatchingRoom() {
	a.playerData.getRandBaseScoreMatchingRoom()
	a.startMatching(roomBaseScoreMatching, a.playerData.BaseScore, 0)
}

func (a *Agent) enterRedPacketMatchingRoom() {
	a.playerData.getRandRedPacketMatchingRoom()
	a.startMatching(roomRedPacketMatching, 0, a.playerData.RedPacketType)
}

func (a *Agent) enterTheRoom() {
	switch a.playerData.RoomType {
	case roomBaseScoreMatching:
		a.startMatching(roomBaseScoreMatching, a.playerData.BaseScore, 0)
	case roomRedPacketMatching:
		a.startMatching(roomRedPacketMatching, 0, a.playerData.RedPacketType)
	}
}

func (a *Agent) addChips() {
	switch a.playerData.RoomType {
	case roomBaseScoreMatching:
		a.fakeWXPay(rand.Intn(30) + 20)
	case roomRedPacketMatching:
		a.fakeWXPay(rand.Intn(200) + 800)
	}
}

//func (a *Agent) enterRandRoom() {
//	a.playerData.getRandRoom()
//	switch a.playerData.RoomType {
//	case roomBaseScoreMatching:
//		a.startMatching(roomBaseScoreMatching, a.playerData.BaseScore, 0)
//	case roomRedPacketMatching:
//		a.startMatching(roomRedPacketMatching, 0, a.playerData.RedPacketType)
//	}
//}

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

func (a *Agent) bid(bid bool) {
	a.writeMsg(&msg.C2S_LandlordBid{
		Bid: bid,
	})
}

func (a *Agent) grab(grab bool) {
	a.writeMsg(&msg.C2S_LandlordGrab{
		Grab: grab,
	})
}

func (a *Agent) double(double bool) {
	a.writeMsg(&msg.C2S_LandlordDouble{
		Double: double,
	})
}

func (a *Agent) showCards(showCards bool) {
	a.writeMsg(&msg.C2S_LandlordShowCards{
		ShowCards: showCards,
	})
}

func (a *Agent) systemHost() {
	a.writeMsg(&msg.C2S_SystemHost{
		Host: true,
	})
}

func (a *Agent) fakeWXPay(totalFee int) {
	a.writeMsg(&msg.C2S_FakeWXPay{
		TotalFee: totalFee,
	})
}
