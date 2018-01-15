package robot

import (
	"czddz-robot/common"
	"czddz-robot/conf"
	"czddz-robot/net"
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	//addr     = "ws://czddz.shenzhouxing.com:3658"
	addr        = "ws://192.168.1.141:3658"
	clients     []*net.Client
	unionids    []string
	nicknames   []string
	headimgurls []string
	count       = 0
	mu          sync.Mutex
	Play        *bool

	robotNumber = 100 // 机器人数量
)

func init() {
	rand.Seed(time.Now().UnixNano())
	names, err := conf.ReadName("D:/robot_nickname.txt")
	names = common.Shuffle2(names)
	if err == nil {
		nicknames = append(nicknames, names[:robotNumber]...)
	} else {
		log.Debug("read file error: %v", err)
	}
	temp := rand.Perm(robotNumber)
	log.Debug("nicknames: %v", nicknames)
	for i := 0; i < robotNumber; i++ {
		unionids = append(unionids, strconv.Itoa(i))
		headimgurls = append(headimgurls, "https://www.shenzhouxing.com/czddz/robot/"+strconv.Itoa(temp[i])+".jpg")
	}
}

func Init() {
	Play = flag.Bool("Play", false, "control robot enter game")
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
	go a.wechatLogin()
	a.readMsg()
}

func (a *Agent) OnClose() {

}
