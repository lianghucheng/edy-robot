package msg

import "edy-robot/poker"

type C2S_Apply struct {
	MatchId string //赛事ID
	Action  int    //1:报名 2:取消报名
}

const (
	S2C_Error_MatchId  = 1 //赛事不存在
	S2C_Error_Coupon   = 2 //点券不足
	S2C_Error_Action   = 3 //已报名(等待开赛)
	S2C_Error_Match    = 4 //玩家已报名了其它赛事
	S2C_Error_Realname = 5 //玩家未实名
)

type S2C_Apply struct {
	Error  int
	RaceID string
	Action int
	Count  int //当前赛事人数
}

type C2S_RaceInfo struct {
}

type RaceInfo struct {
	ID        string  //赛事Id
	Desc      string  //赛事名称
	Award     float64 //赛事
	EnterFee  float64 //报名费
	ConDes    string  //赛事开赛条件
	JoinNum   int     //赛事报名人数
	StartTime int64   // 比赛开始时间
	StartType int     // 比赛开赛方式
	IsSign    bool    // 是否报名
	MatchType string
}

type S2C_RaceInfo struct {
	Races []RaceInfo
}

type C2S_UsrnPwdLogin struct {
	Username string
	Password string
}

type Customer struct {
	WeChat   string //微信
	Email    string //邮箱
	PhoneNum string //电话号码
}

type S2C_Login struct {
	AccountID         int
	Nickname          string
	Headimgurl        string
	Sex               int // 1 男、2 女
	Role              int // 1 玩家、2 代理、3 管理员、4 超管
	Token             string
	AnotherLogin      bool     // 其他设备登录
	FirstLogin        bool     // 首次登录
	AfterTaxAward     float64  // 税后奖金
	Coupon            int      // 点劵数量
	SignIcon          bool     //签到标签是否显示
	NewWelfareIcon    bool     //新人福利标签是否显示
	FirstRechargeIcon bool     //首充标签是否显示
	ShareIcon         bool     //分享推广标签是否显示
	Customer          Customer //客服
	RealName          string
	PhoneNum          string
	BankName          string //银行名称
	BankCardNoTail    string //银行卡号后四位
	SetNickName       bool
}

const (
	S2C_Close_LoginRepeated   = 1  // 您的账号在其他设备上线，非本人操作请注意修改密码
	S2C_Close_InnerError      = 2  // 登录出错，请重新登录
	S2C_Close_TokenInvalid    = 3  // 登录状态失效，请重新登录
	S2C_Close_UsernameInvalid = 5  // 登录出错，用户名无效
	S2C_Close_SystemOff       = 6  // 系统升级维护中，请稍后重试
	S2C_Close_RoleBlack       = 7  // 账号已冻结，请联系客服微信 S2C_Close.WeChatNumber
	S2C_Close_IPChanged       = 8  // 登录IP发生变化，非本人操作请注意修改密码
	S2C_Close_Code_Valid      = 9  // 验证码错误
	S2C_Close_Code_Error      = 10 // 验证码过期了
	S2C_Close_Pwd_Error       = 11 // 密码错误
	S2C_Close_Usrn_Nil        = 12 // 用户名不存在
	S2C_Close_Usrn_Exist      = 13 // 用户名不存在
)

type S2C_Close struct {
	Error        int
	WeChatNumber string
}

type S2C_EnterRoom struct {
	Error      int
	Position   int
	BaseScore  int
	MaxPlayers int // 最大玩家数
}

// 叫分动作
type S2C_ActionLandlordBid struct {
	Position  int
	Countdown int // 倒计时
	Score     []int
}

type C2S_LandlordBid struct {
	Score int //叫的分数
}

type S2C_ActionLandlordDouble struct {
	Countdown int // 倒计时
}

type C2S_LandlordDouble struct {
	Double bool
}

type S2C_LandlordRoundFinalResult struct {
	RoundResults []poker.LandlordPlayerRoundResult
	Countdown    int // 下一局开始时间
}

type S2C_ActionLandlordDiscard struct {
	IsErr             int
	ActionDiscardType int // 出牌动作类型
	Position          int
	Countdown         int     // 倒计时
	PrevDiscards      []int   // 上一次出的牌
	Hint              [][]int // 出牌提示
}

type C2S_LandlordDiscard struct {
	Cards []int
}

type S2C_UpdatePokerHands struct {
	Position      int
	Hands         []int // 手牌
	NumberOfHands int   // 手牌数量
}

type C2S_Heartbeat struct {
}

type C2S_EnterRoom struct{}

