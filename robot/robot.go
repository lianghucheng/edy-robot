package robot

import (
	"czddz-robot/common"
	"czddz-robot/net"
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/timer"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	// addr = "ws://czddz.shenzhouxing.com:3658"
	// addr        = "ws://139.199.180.94:3658"
	addr        = "ws://192.168.1.168:3658"
	clients     []*net.Client
	unionids    []string
	nicknames   []string
	headimgurls []string
	loginIPs    []string
	count       = 0
	mu          sync.Mutex
	Play        *bool

	robotNumber = 1 // 机器人数量

	dispatcher *timer.Dispatcher
)

func init() {
	rand.Seed(time.Now().UnixNano())

	names, ips := make([]string, 0), make([]string, 0)
	var err error
	names, err = common.ReadFile("D:/robot_nickname.txt")
	names = common.Shuffle2(names)

	ips, _ = common.ReadFile("D:/czddz_ip2.txt")
	ips = common.Shuffle2(ips)
	if err == nil {
		nicknames = append(nicknames, names[:robotNumber]...)
		loginIPs = append(loginIPs, ips[:robotNumber]...)
	} else {
		log.Debug("read file error: %v", err)
	}
	temp := rand.Perm(robotNumber)
	log.Debug("loginIP: %v", loginIPs)
	log.Debug("nicknames: %v", nicknames)
	for i := 0; i < robotNumber; i++ {
		unionids = append(unionids, strconv.Itoa(i))
		headimgurls = append(headimgurls, "https://www.shenzhouxing.com/robot/"+strconv.Itoa(temp[i])+".jpg")
	}

	dispatcher = timer.NewDispatcher(0)
}

func Init() {
	Play = flag.Bool("Play", true, "control robot enter game")
	flag.Parse()
	log.Debug("Play: %v", *Play)
	client := new(net.Client)
	client.Addr = addr
	client.ConnNum = robotNumber
	client.ConnectInterval = 3 * time.Second
	client.HandshakeTimeout = 10 * time.Second
	client.PendingWriteNum = 100
	client.MaxMsgLen = 4096
	client.NewAgent = newAgent

	client.Start()
	clients = append(clients, client)
}

func Destroy() {
	for _, client := range clients {
		client.Close()
	}
}

type Agent struct {
	conn       *net.MyConn
	playerData *PlayerData
}

func newAgent(conn *net.MyConn) network.Agent {
	a := new(Agent)
	a.conn = conn
	a.playerData = newPlayerData()
	return a
}

func newPlayerData() *PlayerData {
	playerData := new(PlayerData)
	playerData.Position = -1
	return playerData
}

func (a *Agent) writeMsg(msg interface{}) {
	a.conn.WriteMsg2(msg)
	return
}

func (a *Agent) readMsg() {
	for {
		msg, err := a.conn.ReadMsg()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Debug("error: %v", err)
			}
			break
		}
		//log.Debug("%s", msg)
		jsonMap := map[string]interface{}{}
		err = json.Unmarshal(msg, &jsonMap)
		if err == nil {
			a.handleMsg(jsonMap)
		} else {
			log.Error("%v", err)
		}
	}
}

func (a *Agent) Run() {
	go func() {
		for {
			(<-dispatcher.ChanTimer).Cb()
		}
	}()

	go a.wechatLogin()
	a.readMsg()
}

func (a *Agent) OnClose() {

}
