package rainyun_go_sdk

type Code int

const (
	// CodeInvalidInputParameter 输入参数无效
	CodeInvalidInputParameter Code = 10002
	// CodeVerificationCodeVerificationFailed 验证码验证失败
	CodeVerificationCodeVerificationFailed Code = 10004
	// CodeNeedLogin 需要登录
	CodeNeedLogin Code = 30002
	// CodeConditionsNotMet 未达条件
	CodeConditionsNotMet Code = 30011
	// CodeCorrespondingUserCannotBeFound 无法找到对应的用户
	CodeCorrespondingUserCannotBeFound Code = 30013
	// CodeApikeyError 密钥认证错误或已失效
	CodeApikeyError Code = 30039
)

// 基础响应
type BaseResponse struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

type BasicOperationResponse struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}

// 用户信息
type UserInfo struct {
	Code int `json:"code"`
	Data struct {
		ID                  int     `json:"ID"`
		Name                string  `json:"Name"`
		Email               string  `json:"Email"`
		Phone               string  `json:"Phone"`
		Money               float64 `json:"Money"`
		RegisterTime        int     `json:"RegisterTime"`
		QQOpenID            string  `json:"QQOpenID"`
		QQ                  any     `json:"QQ"`
		WechatOpenID        string  `json:"WechatOpenID"`
		CardCode            any     `json:"CardCode"`
		IconURL             string  `json:"IconUrl"`
		Points              int     `json:"Points"`
		MoneyFromPoint      int     `json:"MoneyFromPoint"`
		Inviter             int     `json:"Inviter"`
		APIKey              string  `json:"APIKey"`
		LastIP              string  `json:"LastIP"`
		BanReason           string  `json:"BanReason"`
		UnsubscribeCount    int     `json:"UnsubscribeCount"`
		AlipayAccount       string  `json:"AlipayAccount"`
		AlipayName          string  `json:"AlipayName"`
		LastLogin           string  `json:"LastLogin"`
		LastLoginArea       string  `json:"LastLoginArea"`
		LoginCount          int     `json:"LoginCount"`
		DLWallet            int     `json:"DLWallet"`
		DLLevel             int     `json:"DLLevel"`
		AdminGroup          string  `json:"AdminGroup"`
		TOTPSecret          string  `json:"TOTPSecret"`
		IsLoginEnableTFA    bool    `json:"IsLoginEnableTFA"`
		IsAllowPointUse     int     `json:"IsAllowPointUse"`
		ShareCode           string  `json:"ShareCode"`
		VipLevel            int     `json:"VipLevel"`
		IsAgent             bool    `json:"IsAgent"`
		Valid               bool    `json:"Valid"`
		Source              string  `json:"Source"`
		ConsumeMonthly      int     `json:"ConsumeMonthly"`
		ConsumeAll          int     `json:"ConsumeAll"`
		ConsumeQuarter      int     `json:"ConsumeQuarter"`
		ResellDaily         int     `json:"ResellDaily"`
		ResellMonthly       int     `json:"ResellMonthly"`
		ResellBeforeMonth   float64 `json:"ResellBeforeMonth"`
		ResellQuarter       float64 `json:"ResellQuarter"`
		ResellAll           float64 `json:"ResellAll"`
		StockDaily          int     `json:"StockDaily"`
		StockMonthly        int     `json:"StockMonthly"`
		StockQuarter        float64 `json:"StockQuarter"`
		StockAll            float64 `json:"StockAll"`
		SecondStockQuarter  int     `json:"SecondStockQuarter"`
		SecondStockAll      int     `json:"SecondStockAll"`
		SubUserMonthly      int     `json:"SubUserMonthly"`
		SubUserAll          int     `json:"SubUserAll"`
		ResellPointsMonthly int     `json:"ResellPointsMonthly"`
		ResellPointsAll     int     `json:"ResellPointsAll"`
		Vip                 struct {
			Title              string  `json:"Title"`
			SaleRequire        int     `json:"SaleRequire"`
			ResellRequire      int     `json:"ResellRequire"`
			CertifyRequired    bool    `json:"CertifyRequired"`
			SaleProfit         float64 `json:"SaleProfit"`
			ResellProfit       float64 `json:"ResellProfit"`
			SecondResellProfit float64 `json:"SecondResellProfit"`
			CanSendCoupons     bool    `json:"CanSendCoupons"`
			CanCustomCode      bool    `json:"CanCustomCode"`
			CanSendMsg         bool    `json:"CanSendMsg"`
			CanTryUsual        bool    `json:"CanTryUsual"`
			FreeDomainCount    int     `json:"FreeDomainCount"`
			FreeSSLCount       int     `json:"FreeSSLCount"`
			CanBeAgent         bool    `json:"CanBeAgent"`
			AgentTitle         string  `json:"AgentTitle"`
			StockRequire       int     `json:"StockRequire"`
			SecondStockRequire int     `json:"SecondStockRequire"`
			StockDiscount      float64 `json:"StockDiscount"`
			SecondStockProfit  float64 `json:"SecondStockProfit"`
		} `json:"VIP"`
		Certify          int    `json:"Certify"`
		LockPoints       int    `json:"LockPoints"`
		CertifyStatus    string `json:"CertifyStatus"`
		CertifyType      string `json:"CertifyType"`
		CertifyAuditNote string `json:"CertifyAuditNote"`
	} `json:"data"`
}

// UserRewardProducts 用户可用积分兑换的产品
type UserRewardProducts struct {
	Code int `json:"code"`
	Data struct {
		Rcs []interface{} `json:"rcs"`
		Rvh []interface{} `json:"rvh"`
		Rgs []interface{} `json:"rgs"`
		Ros []interface{} `json:"ros"`
		Rbm []interface{} `json:"rbm"`
	} `json:"data"`
}

// UserLogsResponse 用户日志响应结构体
type UserLogsResponse struct {
	Code int          `json:"code"`
	Data UserLogsData `json:"data"`
}

// UserLogsData 用户日志数据
type UserLogsData struct {
	TotalRecords int             `json:"TotalRecords"`
	Records      []UserLogRecord `json:"Records"`
}

// UserLogRecord 单条用户日志记录
type UserLogRecord struct {
	ID            int                    `json:"ID"`
	UID           int                    `json:"UID"`
	StartTime     int                    `json:"StartTime"`
	EndTime       int                    `json:"EndTime"`
	Type          string                 `json:"Type"`
	Duration      string                 `json:"Duration"`
	ProductID     int                    `json:"ProductID"`
	PlanID        int                    `json:"PlanID"`
	Price         float64                `json:"Price"`
	FromPoint     int                    `json:"FromPoint"`
	CutPrice      int                    `json:"CutPrice"`
	StockPrice    float64                `json:"StockPrice"`
	AgentID       int                    `json:"AgentID"`
	Valid         bool                   `json:"Valid"`
	Discard       bool                   `json:"Discard"`
	InvoiceIssued int                    `json:"InvoiceIssued"`
	Data          map[string]interface{} `json:"Data"`
	Region        string                 `json:"Region"`
}

// 发布优惠券给下级用户
type PublishCouponsToLowerLevelUsersRequest struct {
	BaseLimit      int    `json:"base_limit"`      // 满减条件(满多少才能用)
	Color          string `json:"color"`           // 颜色: waring: 黄，danger: 红，success: 绿...
	Count          int    `json:"count"`           // 发放数量
	ExpDate        int    `json:"exp_date"`        // 过期时间(timestamp)
	FriendlyName   string `json:"friendly_name"`   // 优惠券标题
	Type           string `json:"type"`            // 类型: discount:折扣, normal: 直减
	UID            int    `json:"uid"`             // 要发放到用户ID(如:114514),为空时则返回兑换码
	UsableDuration string `json:"usable_duration"` // unknown
	UsableProduct  string `json:"usable_product"`  // 可用产品(默认全部),","分隔: renew,create,upgrade
	UsableScenes   string `json:"usable_scenes"`   // 适用操作(默认全部),","分隔: rvh,rcs,rgs,ros,rbm
	Value          int    `json:"value"`           // 直减(元)/折扣(折), 折扣时:1~9:一~九折；11~99:一一~九九折
}

// 发送优惠券到积分商城
type PostCouponsToPointsMallRequest struct {
	AvailableDays  int    `json:"available_days"`  // 可用天数
	BaseLimit      int    `json:"base_limit"`      // 满减条件(满多少才能用)
	BuyLimit       int    `json:"buy_limit"`       // 领取次数限制
	Color          string `json:"color"`           // 颜色: waring: 黄, danger: 红, success: 绿, info: 蓝
	Count          int    `json:"count"`           // 发放数量
	EndDate        int    `json:"end_date"`        // 截止日期
	FirstSend      bool   `json:"first_send"`      // 绑定微信后立即自动领取(设置后不可手动领取并且不会显示)
	FriendlyName   string `json:"friendly_name"`   // 优惠券名称
	Name           string `json:"name"`            // 标识名称
	Order          int    `json:"order"`           // 排序,越大越靠前
	Points         int    `json:"points"`          // 领取积分
	Type           string `json:"type"`            // 类型: discount:折扣, normal: 直减
	UsableDuration string `json:"usable_duration"` // unknown
	UsableProduct  string `json:"usable_product"`  // 可用产品(默认全部),","分隔: renew,create,upgrade
	UsableScenes   string `json:"usable_scenes"`   // 适用操作(默认全部),","分隔: rvh,rcs,rgs,ros,rbm
	Value          int    `json:"value"`           // 直减(元)/折扣(折), 折扣时:1~9:一~九折；11~99:一一~九九折
}

const (
	TBPass   = "关注成功"
	BiliPass = "雨云爱你"
	QQPass   = "我爱雨云"
)
