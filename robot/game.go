package robot

import (
	"edy-robot/msg"
	"github.com/name5566/leaf/log"
	"math/rand"
	"time"
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
	a.writeMsg(&msg.C2S_UsrnPwdLogin{
		Username:    unionids[count],
		Password:   "123456789",
	})
	count++
}

func (a *Agent) enterRoom() {
	a.writeMsg(&msg.C2S_EnterRoom{})
}

func (a *Agent) signIn() {
	log.Debug("是否已报名：%v    是否已在比赛：%v", a.isSign(), a.playerData.isPlay)
	if !a.isSign() && !a.playerData.isPlay {
		log.Debug("报名")
		a.writeMsg(&msg.C2S_Apply{
			MatchId:a.matchids[rand.Intn(len(a.matchids))],
			Action:1,
		})
		a.signOutTimer = time.AfterFunc(10 * time.Second, func() {
			if a != nil {
				a.signOut()
			}
		})
	}
}

func (a *Agent) signOut() {
	if a.isSign() && !a.playerData.isPlay {
		log.Debug("退签")
		a.writeMsg(&msg.C2S_Apply{
			MatchId: a.currMatchid,
			Action:2,
		})
		a.currMatchid = ""
		time.AfterFunc(10 * time.Second, func() {
			if a != nil {
				a.signIn()
			}
		})
	}
}

func (a *Agent) doBid(scores []int) {
	a.writeMsg(&msg.C2S_LandlordBid{
		Score: scores[rand.Intn(len(scores))],
	})
}

func (a *Agent) doDouble() {
	temp := rand.Intn(2)
	double := false
	if temp & 1 == 0 {
		double = true
	}
	a.writeMsg(&msg.C2S_LandlordDouble{
		Double:double,
	})
}

func (a *Agent) doDiscard(actionDiscardType int) {
	if len(a.playerData.Hint) >= 1 {
		Delay(func() {
			a.writeMsg(&msg.C2S_LandlordDiscard{
				Cards:a.playerData.Hint[rand.Intn(len(a.playerData.Hint))],
			})
		})
	} else {
		if actionDiscardType == 1 {
			Delay(func() {
				a.writeMsg(&msg.C2S_LandlordDiscard{
					Cards:[]int{a.playerData.hands[rand.Intn(len(a.playerData.hands))]},
				})
			})
			return
		}
		Delay(func() {
			a.writeMsg(&msg.C2S_LandlordDiscard{})
		})
	}
}

func (a *Agent) isSign() bool {
	return a.currMatchid != ""
}