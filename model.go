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
	// CodeOutOfStock 缺货
	CodeOutOfStock Code = 30023
	// CodeApikeyError 密钥认证错误或已失效
	CodeApikeyError Code = 30039
	// CodeRequiresSecondaryVerification 需要二次验证
	CodeRequiresSecondaryVerification Code = 30043
	// CodeInvalidVerification 验证无效(验证码错误)
	CodeInvalidVerification Code = 30047
	// CodePortIsAlreadyInUse 端口已被使用
	CodePortIsAlreadyInUse Code = 70020
	// CodeProductHasExpired 产品已过期
	CodeProductHasExpired Code = 70021
	// CodeTheCurrentStateOfTheProductCannotPerformThisOperation 产品当前状态无法执行此操作
	CodeTheCurrentStateOfTheProductCannotPerformThisOperation Code = 70026
	// CodeAgreementRequired 需提供协议
	CodeAgreementRequired Code = 70057
)

// 基础响应
type BaseResponse struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

type BasicOperationResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
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

// 兑换积分物品
type RedeemPointsForItemRequest struct {
	ItemID int `json:"item_id"`
}

// 积分商城物品响应
type PointsMallItemsResponse struct {
	Code int `json:"code"`
	Data []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Points         int    `json:"points"`
		Type           string `json:"type"`
		AvailableStock int    `json:"available_stock"`
		FriendlyName   string `json:"friendly_name"`
		ItemData       struct {
			Color           string      `json:"color,omitempty"`
			FriendlyName    string      `json:"friendly_name,omitempty"`
			UsableScenes    string      `json:"usable_scenes,omitempty"`
			AvailableDays   int         `json:"available_days,omitempty"`
			UsableProduct   string      `json:"usable_product,omitempty"`
			BaseLimit       int         `json:"base_limit,omitempty"`
			PublicPoint     int         `json:"public_point,omitempty"`
			Type            string      `json:"type,omitempty"`
			UsableDuration  string      `json:"usable_duration,omitempty"`
			UsablePlanID    int         `json:"usable_plan_id,omitempty"`
			Value           interface{} `json:"value,omitempty"`
			ProductType     string      `json:"product_type,omitempty"`
			ProductSubtype  string      `json:"product_subtype,omitempty"`
			DurationSeconds int         `json:"duration_seconds,omitempty"`
			ProductConfig   struct {
				OsID         int    `json:"os_id"`
				PlanID       int    `json:"plan_id"`
				Subtype      string `json:"subtype"`
				Duration     int    `json:"duration"`
				PayMode      string `json:"pay_mode"`
				PanelUser    string `json:"panel_user"`
				EggTypeID    int    `json:"egg_type_id"`
				WithCoupon   int    `json:"with_coupon"`
				CPULimitMode bool   `json:"cpu_limit_mode"`
				Config       struct {
					CPU        int `json:"cpu"`
					Backup     int `json:"backup"`
					Memory     int `json:"memory"`
					NetIn      int `json:"net_in"`
					NetOut     int `json:"net_out"`
					Database   int `json:"database"`
					BaseDisk   int `json:"base_disk"`
					DataDisk   int `json:"data_disk"`
					Allocation int `json:"allocation"`
				} `json:"config"`
			} `json:"product_config,omitempty"`
			Desc    string `json:"desc,omitempty"`
			ImgURL  string `json:"img_url,omitempty"`
			DescURL string `json:"desc_url,omitempty"`
		} `json:"item_data"`
		BuyLimit      int    `json:"buy_limit"`
		SenderID      int    `json:"sender_id"`
		FirstSend     bool   `json:"first_send"`
		ByInvite      bool   `json:"by_invite"`
		Color         string `json:"color"`
		Order         int    `json:"order"`
		PublicTime    int    `json:"public_time"`
		EndDate       int    `json:"end_date"`
		AutoRefresh   int    `json:"auto_refresh"`
		RefreshLimit  int    `json:"refresh_limit"`
		MoneyRequired int    `json:"money_required"`
	} `json:"data"`
}

// 请求二次验证
type Request2FARequest struct {
	Type string `json:"type"`
}

// 验证二次验证结果
type Verify2FAResultRequest struct {
	AuthCode int `json:"auth_code"`
}

// 设置IP描述
type SetRcsEipDescriptionRequest struct {
	IP          string `json:"ip"`
	Description string `json:"description"`
}

// RCS列表
type RcsListResponse struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords int `json:"TotalRecords"`
		Records      []struct {
			ExpDate         int    `json:"ExpDate"`
			ExpireNotice    int    `json:"ExpireNotice"`
			AutoRenew       bool   `json:"AutoRenew"`
			UnsubscribeAble bool   `json:"UnsubscribeAble"`
			Try             bool   `json:"Try"`
			ID              int    `json:"ID"`
			UID             int    `json:"UID"`
			PlanID          int    `json:"PlanID"`
			CreateDate      int    `json:"CreateDate"`
			NodeUUID        string `json:"NodeUUID"`
			Node            struct {
				UUID              string      `json:"UUID"`
				AuthKey           string      `json:"AuthKey"`
				Region            string      `json:"Region"`
				IPRegion          string      `json:"IpRegion"`
				Machine           string      `json:"Machine"`
				Product           string      `json:"Product"`
				Subtype           string      `json:"Subtype"`
				ChineseName       string      `json:"ChineseName"`
				PhysicalNode      string      `json:"PhysicalNode"`
				Config            string      `json:"Config"`
				Stock             interface{} `json:"Stock"`
				StatusData        string      `json:"StatusData"`
				ShowMonitorData   string      `json:"ShowMonitorData"`
				UpdateTime        string      `json:"UpdateTime"`
				GitRepositoryName string      `json:"GitRepositoryName"`
				CertifyRequired   bool        `json:"CertifyRequired"`
				IsDisableBackup   bool        `json:"IsDisableBackup"`
				IsHidden          bool        `json:"IsHidden"`
				NodeName          string      `json:"NodeName"`
			} `json:"Node"`
			Status                  string `json:"Status"`
			StopReason              string `json:"StopReason"`
			RewardPointsToBeCollect int    `json:"RewardPointsToBeCollect"`
			Tag                     string `json:"Tag"`
			OsID                    int    `json:"OsID"`
			OsName                  string `json:"OsName"`
			HostName                string `json:"HostName"`
			DefaultPass             string `json:"DefaultPass"`
			MainIPv4                string `json:"MainIPv4"`
			IntIPv4                 string `json:"IntIPv4"`
			UsageData               struct {
				CPU     float64 `json:"CPU"`
				MaxMem  int64   `json:"MaxMem"`
				FreeMem int64   `json:"FreeMem"`
				UsedMem int     `json:"UsedMem"`
				Disks   struct {
					NAMING_FAILED struct {
						Total int64 `json:"Total"`
						Used  int64 `json:"Used"`
					} `json:"/"`
				} `json:"Disks"`
				DiskRead    int         `json:"DiskRead"`
				DiskWrite   int         `json:"DiskWrite"`
				NetOut      float64     `json:"NetOut"`
				NetIn       int         `json:"NetIn"`
				SmartHealth interface{} `json:"SmartHealth"`
				SmartTemp   int         `json:"SmartTemp"`
				UpdateTime  int         `json:"UpdateTime"`
			} `json:"UsageData"`
			Zone                 string `json:"Zone"`
			NatPublicIP          string `json:"NatPublicIP"`
			NatPublicDomain      string `json:"NatPublicDomain"`
			NATSpareDomain       string `json:"NATSpareDomain"`
			NetIn                int    `json:"NetIn"`
			NetOut               int    `json:"NetOut"`
			NowNetIn             int    `json:"NowNetIn"`
			NowNetOut            int    `json:"NowNetOut"`
			NetMode              string `json:"NetMode"`
			BridgeSyncing        bool   `json:"BridgeSyncing"`
			VnetID               int    `json:"VnetID"`
			UpdateTime           int    `json:"UpdateTime"`
			FwSyncTime           int    `json:"FwSyncTime"`
			FwMode               string `json:"FwMode"`
			AbCPULimit           int    `json:"AbCpuLimit"`
			AbNetLimit           int    `json:"AbNetLimit"`
			AbWhiteReason        string `json:"AbWhiteReason"`
			TrafficBytes         int64  `json:"TrafficBytes"`
			TrafficResetDate     int    `json:"TrafficResetDate"`
			TrafficBytesToday    int    `json:"TrafficBytesToday"`
			TrafficBytesDayLimit int64  `json:"TrafficBytesDayLimit"`
			TrafficOnLimit       int    `json:"TrafficOnLimit"`
			Plan                 struct {
				ID              int         `json:"id"`
				Region          string      `json:"region"`
				Subtype         string      `json:"subtype"`
				PlanName        string      `json:"plan_name"`
				Machine         string      `json:"machine"`
				ChargeType      string      `json:"charge_type"`
				Chinese         string      `json:"chinese"`
				IsFree          bool        `json:"is_free"`
				PointRenewPrice interface{} `json:"point_renew_price"`
				IsSelling       bool        `json:"is_selling"`
				Price           int         `json:"price"`
				TrafficBaseGb   int         `json:"traffic_base_gb"`
				TrafficPrice    struct {
					Num300  int `json:"300"`
					Num1024 int `json:"1024"`
					Num2048 int `json:"2048"`
				} `json:"traffic_price"`
				CPU            int         `json:"cpu"`
				Memory         int         `json:"memory"`
				NetIn          int         `json:"net_in"`
				NetOut         int         `json:"net_out"`
				IPPrices       interface{} `json:"ip_prices"`
				IPSelling      interface{} `json:"ip_selling"`
				AutoRestock    int         `json:"auto_restock"`
				AvailableStock int         `json:"available_stock"`
				GpuMemorySize  int         `json:"gpu_memory_size"`
				DgpuDevType    string      `json:"dgpu_dev_type"`
				WebbarConfig   interface{} `json:"webbar_config"`
				NoBackup       bool        `json:"no_backup"`
				DiskPrice      struct {
					Ssd float64 `json:"ssd"`
					Hdd float64 `json:"hdd"`
				} `json:"disk_price"`
			} `json:"Plan"`
			OsInfo struct {
				ID             int    `json:"id"`
				Region         string `json:"region"`
				Subtype        string `json:"subtype"`
				Machine        string `json:"machine"`
				Name           string `json:"name"`
				Version        string `json:"version"`
				SyncStatus     string `json:"sync_status"`
				OsType         string `json:"os_type"`
				ChineseName    string `json:"chinese_name"`
				Icon           string `json:"icon"`
				IsWithBbr      bool   `json:"is_with_bbr"`
				IsEol          bool   `json:"is_eol"`
				IsAvailable    bool   `json:"is_available"`
				Order          int    `json:"order"`
				LatestFilename string `json:"latest_filename"`
				NoVMAgent      bool   `json:"no_vm_agent"`
			} `json:"OsInfo"`
			CPU                   int    `json:"CPU"`
			Memory                int    `json:"Memory"`
			Disk                  int    `json:"Disk"`
			RBSKeepLast           int    `json:"RBSKeepLast"`
			RBSAutoBackup         string `json:"RBSAutoBackup"`
			RBSLastAutoBackupDate int    `json:"RBSLastAutoBackupDate"`
			FastAppInstallTaskKey string `json:"FastAppInstallTaskKey"`
			GPUDevice             string `json:"GPUDevice"`
			GPUMemorySize         int    `json:"GPUMemorySize"`
			DGPUEnable            bool   `json:"DGPUEnable"`
			NoPrimaryGPU          bool   `json:"NoPrimaryGPU"`
			WebbarMinutes         int    `json:"WebbarMinutes"`
			WebbarResetDate       int    `json:"WebbarResetDate"`
		} `json:"Records"`
	} `json:"data"`
}

// 创建RCS
type CreateRcsRequest struct {
	AddDiskSize  int    `json:"add_disk_size"`  // 额外硬盘容量GB
	PlanID       int    `json:"plan_id"`        // 套餐ID
	Duration     int    `json:"duration"`       // 创建时长(天)
	OsID         int    `json:"os_id"`          // 系统ID
	WithEipNum   int    `json:"with_eip_num"`   // 创建IP数量
	WithEipFlags string `json:"with_eip_flags"` // 是否开启高防，us_ddosip -> 美国高防，nb_ddosip -> 宁波高防
	WithEipType  string `json:"with_eip_type"`  // ipv4(默认)/ipv6
	WithCouponID int    `json:"with_coupon_id"` // 优惠券ID
	Try          bool   `json:"try"`            // 是否为试用
	NodeUUID     string `json:"node_uuid"`      // 指定节点(管理员可用，用户不可用)
	AppVars      []struct {
		AppID int `json:"app_id"`
		Vars  any `json:"vars,omitempty"`
	} `json:"app_vars"` // 预装应用
	Zone string `json:"zone"` // 内网可用区
}

// 创建RCS响应
type CreateRcsResopnse struct {
	Code int `json:"code"`
	Data struct {
		ExpDate         int    `json:"ExpDate"`
		ExpireNotice    int    `json:"ExpireNotice"`
		AutoRenew       bool   `json:"AutoRenew"`
		UnsubscribeAble bool   `json:"UnsubscribeAble"`
		Try             bool   `json:"Try"`
		ID              int    `json:"ID"`
		UID             int    `json:"UID"`
		PlanID          int    `json:"PlanID"`
		CreateDate      int    `json:"CreateDate"`
		NodeUUID        string `json:"NodeUUID"`
		Node            struct {
			UUID              string      `json:"UUID"`
			AuthKey           string      `json:"AuthKey"`
			Region            string      `json:"Region"`
			IPRegion          string      `json:"IpRegion"`
			Machine           string      `json:"Machine"`
			Product           string      `json:"Product"`
			Subtype           string      `json:"Subtype"`
			ChineseName       string      `json:"ChineseName"`
			PhysicalNode      string      `json:"PhysicalNode"`
			Config            string      `json:"Config"`
			Stock             interface{} `json:"Stock"`
			StatusData        string      `json:"StatusData"`
			ShowMonitorData   string      `json:"ShowMonitorData"`
			UpdateTime        string      `json:"UpdateTime"`
			GitRepositoryName string      `json:"GitRepositoryName"`
			CertifyRequired   bool        `json:"CertifyRequired"`
			IsDisableBackup   bool        `json:"IsDisableBackup"`
			IsHidden          bool        `json:"IsHidden"`
			NodeName          string      `json:"NodeName"`
		} `json:"Node"`
		Status                  string `json:"Status"`
		StopReason              string `json:"StopReason"`
		RewardPointsToBeCollect int    `json:"RewardPointsToBeCollect"`
		Tag                     string `json:"Tag"`
		OsID                    int    `json:"OsID"`
		OsName                  string `json:"OsName"`
		HostName                string `json:"HostName"`
		DefaultPass             string `json:"DefaultPass"`
		MainIPv4                string `json:"MainIPv4"`
		IntIPv4                 string `json:"IntIPv4"`
		Zone                    string `json:"Zone"`
		NatPublicIP             string `json:"NatPublicIP"`
		NatPublicDomain         string `json:"NatPublicDomain"`
		NATSpareDomain          string `json:"NATSpareDomain"`
		NetIn                   int    `json:"NetIn"`
		NetOut                  int    `json:"NetOut"`
		NowNetIn                int    `json:"NowNetIn"`
		NowNetOut               int    `json:"NowNetOut"`
		NetMode                 string `json:"NetMode"`
		BridgeSyncing           bool   `json:"BridgeSyncing"`
		VnetID                  int    `json:"VnetID"`
		UpdateTime              int    `json:"UpdateTime"`
		FwSyncTime              int    `json:"FwSyncTime"`
		FwMode                  string `json:"FwMode"`
		AbCPULimit              int    `json:"AbCpuLimit"`
		AbNetLimit              int    `json:"AbNetLimit"`
		AbWhiteReason           string `json:"AbWhiteReason"`
		TrafficBytes            int64  `json:"TrafficBytes"`
		TrafficResetDate        int    `json:"TrafficResetDate"`
		TrafficBytesToday       int    `json:"TrafficBytesToday"`
		TrafficBytesDayLimit    int    `json:"TrafficBytesDayLimit"`
		TrafficOnLimit          int    `json:"TrafficOnLimit"`
		Plan                    struct {
			ID              int         `json:"id"`
			Region          string      `json:"region"`
			Subtype         string      `json:"subtype"`
			PlanName        string      `json:"plan_name"`
			Machine         string      `json:"machine"`
			ChargeType      string      `json:"charge_type"`
			Chinese         string      `json:"chinese"`
			IsFree          bool        `json:"is_free"`
			PointRenewPrice interface{} `json:"point_renew_price"`
			IsSelling       bool        `json:"is_selling"`
			Price           int         `json:"price"`
			TrafficBaseGb   int         `json:"traffic_base_gb"`
			TrafficPrice    struct {
				Num200  int `json:"200"`
				Num1024 int `json:"1024"`
			} `json:"traffic_price"`
			CPU            int         `json:"cpu"`
			Memory         int         `json:"memory"`
			NetIn          int         `json:"net_in"`
			NetOut         int         `json:"net_out"`
			IPPrices       interface{} `json:"ip_prices"`
			IPSelling      interface{} `json:"ip_selling"`
			AutoRestock    int         `json:"auto_restock"`
			AvailableStock int         `json:"available_stock"`
			GpuMemorySize  int         `json:"gpu_memory_size"`
			DgpuDevType    string      `json:"dgpu_dev_type"`
			WebbarConfig   interface{} `json:"webbar_config"`
			NoBackup       bool        `json:"no_backup"`
			DiskPrice      struct {
				Ssd float64 `json:"ssd"`
				Hdd float64 `json:"hdd"`
			} `json:"disk_price"`
		} `json:"Plan"`
		OsInfo struct {
			ID             int    `json:"id"`
			Region         string `json:"region"`
			Subtype        string `json:"subtype"`
			Machine        string `json:"machine"`
			Name           string `json:"name"`
			Version        string `json:"version"`
			SyncStatus     string `json:"sync_status"`
			OsType         string `json:"os_type"`
			ChineseName    string `json:"chinese_name"`
			Icon           string `json:"icon"`
			IsWithBbr      bool   `json:"is_with_bbr"`
			IsEol          bool   `json:"is_eol"`
			IsAvailable    bool   `json:"is_available"`
			Order          int    `json:"order"`
			LatestFilename string `json:"latest_filename"`
			NoVMAgent      bool   `json:"no_vm_agent"`
		} `json:"OsInfo"`
		CPU                   int    `json:"CPU"`
		Memory                int    `json:"Memory"`
		Disk                  int    `json:"Disk"`
		RBSKeepLast           int    `json:"RBSKeepLast"`
		RBSAutoBackup         string `json:"RBSAutoBackup"`
		RBSLastAutoBackupDate int    `json:"RBSLastAutoBackupDate"`
		FastAppInstallTaskKey string `json:"FastAppInstallTaskKey"`
		GPUDevice             string `json:"GPUDevice"`
		GPUMemorySize         int    `json:"GPUMemorySize"`
		DGPUEnable            bool   `json:"DGPUEnable"`
		NoPrimaryGPU          bool   `json:"NoPrimaryGPU"`
		WebbarMinutes         int    `json:"WebbarMinutes"`
		WebbarResetDate       int    `json:"WebbarResetDate"`
	} `json:"data"`
}

// RCS 详情
type RcsDetails struct {
	Code int `json:"code"`
	Data struct {
		Data struct {
			ExpDate         int    `json:"ExpDate"`
			ExpireNotice    int    `json:"ExpireNotice"`
			AutoRenew       bool   `json:"AutoRenew"`
			UnsubscribeAble bool   `json:"UnsubscribeAble"`
			Try             bool   `json:"Try"`
			ID              int    `json:"ID"`
			UID             int    `json:"UID"`
			PlanID          int    `json:"PlanID"`
			CreateDate      int    `json:"CreateDate"`
			NodeUUID        string `json:"NodeUUID"`
			Node            struct {
				UUID              string      `json:"UUID"`
				AuthKey           string      `json:"AuthKey"`
				Region            string      `json:"Region"`
				IPRegion          string      `json:"IpRegion"`
				Machine           string      `json:"Machine"`
				Product           string      `json:"Product"`
				Subtype           string      `json:"Subtype"`
				ChineseName       string      `json:"ChineseName"`
				PhysicalNode      string      `json:"PhysicalNode"`
				Config            string      `json:"Config"`
				Stock             interface{} `json:"Stock"`
				StatusData        string      `json:"StatusData"`
				ShowMonitorData   string      `json:"ShowMonitorData"`
				UpdateTime        string      `json:"UpdateTime"`
				GitRepositoryName string      `json:"GitRepositoryName"`
				CertifyRequired   bool        `json:"CertifyRequired"`
				IsDisableBackup   bool        `json:"IsDisableBackup"`
				IsHidden          bool        `json:"IsHidden"`
				NodeName          string      `json:"NodeName"`
			} `json:"Node"`
			Status                  string `json:"Status"`
			StopReason              string `json:"StopReason"`
			RewardPointsToBeCollect int    `json:"RewardPointsToBeCollect"`
			Tag                     string `json:"Tag"`
			OsID                    int    `json:"OsID"`
			OsName                  string `json:"OsName"`
			HostName                string `json:"HostName"`
			DefaultPass             string `json:"DefaultPass"`
			MainIPv4                string `json:"MainIPv4"`
			IntIPv4                 string `json:"IntIPv4"`
			UsageData               struct {
				CPU     float64 `json:"CPU"`
				MaxMem  int64   `json:"MaxMem"`
				FreeMem int64   `json:"FreeMem"`
				UsedMem int     `json:"UsedMem"`
				Disks   struct {
					NAMING_FAILED struct {
						Total int64 `json:"Total"`
						Used  int64 `json:"Used"`
					} `json:"/"`
				} `json:"Disks"`
				DiskRead    int         `json:"DiskRead"`
				DiskWrite   float64     `json:"DiskWrite"`
				NetOut      float64     `json:"NetOut"`
				NetIn       float64     `json:"NetIn"`
				SmartHealth interface{} `json:"SmartHealth"`
				SmartTemp   int         `json:"SmartTemp"`
				UpdateTime  int         `json:"UpdateTime"`
			} `json:"UsageData"`
			Zone                 string `json:"Zone"`
			NatPublicIP          string `json:"NatPublicIP"`
			NatPublicDomain      string `json:"NatPublicDomain"`
			NATSpareDomain       string `json:"NATSpareDomain"`
			NetIn                int    `json:"NetIn"`
			NetOut               int    `json:"NetOut"`
			NowNetIn             int    `json:"NowNetIn"`
			NowNetOut            int    `json:"NowNetOut"`
			NetMode              string `json:"NetMode"`
			BridgeSyncing        bool   `json:"BridgeSyncing"`
			VnetID               int    `json:"VnetID"`
			UpdateTime           int    `json:"UpdateTime"`
			FwSyncTime           int    `json:"FwSyncTime"`
			FwMode               string `json:"FwMode"`
			AbCPULimit           int    `json:"AbCpuLimit"`
			AbNetLimit           int    `json:"AbNetLimit"`
			AbWhiteReason        string `json:"AbWhiteReason"`
			TrafficBytes         int64  `json:"TrafficBytes"`
			TrafficResetDate     int    `json:"TrafficResetDate"`
			TrafficBytesToday    int    `json:"TrafficBytesToday"`
			TrafficBytesDayLimit int64  `json:"TrafficBytesDayLimit"`
			TrafficOnLimit       int    `json:"TrafficOnLimit"`
			Plan                 struct {
				ID              int         `json:"id"`
				Region          string      `json:"region"`
				Subtype         string      `json:"subtype"`
				PlanName        string      `json:"plan_name"`
				Machine         string      `json:"machine"`
				ChargeType      string      `json:"charge_type"`
				Chinese         string      `json:"chinese"`
				IsFree          bool        `json:"is_free"`
				PointRenewPrice interface{} `json:"point_renew_price"`
				IsSelling       bool        `json:"is_selling"`
				Price           int         `json:"price"`
				TrafficBaseGb   int         `json:"traffic_base_gb"`
				TrafficPrice    struct {
					Num300  int `json:"300"`
					Num1024 int `json:"1024"`
					Num2048 int `json:"2048"`
				} `json:"traffic_price"`
				CPU      int `json:"cpu"`
				Memory   int `json:"memory"`
				NetIn    int `json:"net_in"`
				NetOut   int `json:"net_out"`
				IPPrices struct {
					NAMING_FAILED int `json:""`
					Ipv6          int `json:"ipv6"`
				} `json:"ip_prices"`
				IPSelling      interface{} `json:"ip_selling"`
				AutoRestock    int         `json:"auto_restock"`
				AvailableStock int         `json:"available_stock"`
				GpuMemorySize  int         `json:"gpu_memory_size"`
				DgpuDevType    string      `json:"dgpu_dev_type"`
				WebbarConfig   interface{} `json:"webbar_config"`
				NoBackup       bool        `json:"no_backup"`
				DiskPrice      struct {
					Ssd float64 `json:"ssd"`
					Hdd float64 `json:"hdd"`
				} `json:"disk_price"`
			} `json:"Plan"`
			OsInfo struct {
				ID             int    `json:"id"`
				Region         string `json:"region"`
				Subtype        string `json:"subtype"`
				Machine        string `json:"machine"`
				Name           string `json:"name"`
				Version        string `json:"version"`
				SyncStatus     string `json:"sync_status"`
				OsType         string `json:"os_type"`
				ChineseName    string `json:"chinese_name"`
				Icon           string `json:"icon"`
				IsWithBbr      bool   `json:"is_with_bbr"`
				IsEol          bool   `json:"is_eol"`
				IsAvailable    bool   `json:"is_available"`
				Order          int    `json:"order"`
				LatestFilename string `json:"latest_filename"`
				NoVMAgent      bool   `json:"no_vm_agent"`
			} `json:"OsInfo"`
			CPU                   int    `json:"CPU"`
			Memory                int    `json:"Memory"`
			Disk                  int    `json:"Disk"`
			RBSKeepLast           int    `json:"RBSKeepLast"`
			RBSAutoBackup         string `json:"RBSAutoBackup"`
			RBSLastAutoBackupDate int    `json:"RBSLastAutoBackupDate"`
			FastAppInstallTaskKey string `json:"FastAppInstallTaskKey"`
			GPUDevice             string `json:"GPUDevice"`
			GPUMemorySize         int    `json:"GPUMemorySize"`
			DGPUEnable            bool   `json:"DGPUEnable"`
			NoPrimaryGPU          bool   `json:"NoPrimaryGPU"`
			WebbarMinutes         int    `json:"WebbarMinutes"`
			WebbarResetDate       int    `json:"WebbarResetDate"`
		} `json:"Data"`
		UpgradeablePlans []struct {
			ID              int         `json:"id"`
			Region          string      `json:"region"`
			Subtype         string      `json:"subtype"`
			PlanName        string      `json:"plan_name"`
			Machine         string      `json:"machine"`
			ChargeType      string      `json:"charge_type"`
			Chinese         string      `json:"chinese"`
			IsFree          bool        `json:"is_free"`
			PointRenewPrice interface{} `json:"point_renew_price"`
			IsSelling       bool        `json:"is_selling"`
			Price           int         `json:"price"`
			TrafficBaseGb   int         `json:"traffic_base_gb"`
			TrafficPrice    interface{} `json:"traffic_price"`
			CPU             int         `json:"cpu"`
			Memory          int         `json:"memory"`
			NetIn           int         `json:"net_in"`
			NetOut          int         `json:"net_out"`
			IPPrices        interface{} `json:"ip_prices"`
			IPSelling       interface{} `json:"ip_selling"`
			AutoRestock     int         `json:"auto_restock"`
			AvailableStock  int         `json:"available_stock"`
			GpuMemorySize   int         `json:"gpu_memory_size"`
			DgpuDevType     string      `json:"dgpu_dev_type"`
			WebbarConfig    interface{} `json:"webbar_config"`
			NoBackup        bool        `json:"no_backup"`
			DiskPrice       struct {
				Ssd float64 `json:"ssd"`
				Hdd float64 `json:"hdd"`
			} `json:"disk_price"`
		} `json:"UpgradeablePlans"`
		RBSList []struct {
			ID        int    `json:"ID"`
			UID       int    `json:"UID"`
			ProductID int    `json:"ProductID"`
			NodeUUID  string `json:"NodeUUID"`
			Node      struct {
				UUID              string      `json:"UUID"`
				AuthKey           string      `json:"AuthKey"`
				Region            string      `json:"Region"`
				IPRegion          string      `json:"IpRegion"`
				Machine           string      `json:"Machine"`
				Product           string      `json:"Product"`
				Subtype           string      `json:"Subtype"`
				ChineseName       string      `json:"ChineseName"`
				PhysicalNode      string      `json:"PhysicalNode"`
				Config            string      `json:"Config"`
				Stock             interface{} `json:"Stock"`
				StatusData        string      `json:"StatusData"`
				ShowMonitorData   string      `json:"ShowMonitorData"`
				UpdateTime        string      `json:"UpdateTime"`
				GitRepositoryName string      `json:"GitRepositoryName"`
				CertifyRequired   bool        `json:"CertifyRequired"`
				IsDisableBackup   bool        `json:"IsDisableBackup"`
				IsHidden          bool        `json:"IsHidden"`
				NodeName          string      `json:"NodeName"`
			} `json:"Node"`
			Label          string `json:"Label"`
			FileName       string `json:"FileName"`
			PackSize       int64  `json:"PackSize"`
			CreateTime     int    `json:"CreateTime"`
			FinishTime     int    `json:"FinishTime"`
			Retry          int    `json:"Retry"`
			AdditionalInfo struct {
				Osid           int    `json:"OSID"`
				OSName         string `json:"OSName"`
				OSBaseDiskSize int    `json:"OSBaseDiskSize"`
				OSDataDiskSize int    `json:"OSDataDiskSize"`
				Slots          struct {
					Num0 struct {
						DiskType string `json:"DiskType"`
						Backup   bool   `json:"Backup"`
						Size     int    `json:"Size"`
					} `json:"0"`
				} `json:"Slots"`
			} `json:"AdditionalInfo"`
			Status string `json:"Status"`
		} `json:"RBSList"`
		NatList   []interface{} `json:"NatList"`
		EDiskList []struct {
			ID       int    `json:"ID"`
			Slot     int    `json:"Slot"`
			UID      int    `json:"UID"`
			DiskType string `json:"DiskType"`
			Tag      string `json:"Tag"`
			OSName   string `json:"OSName"`
			Vid      int    `json:"VID"`
			Size     int    `json:"Size"`
			Backup   bool   `json:"Backup"`
		} `json:"EDiskList"`
		EIPList []struct {
			ID          int    `json:"ID"`
			IPRegion    string `json:"IpRegion"`
			Region      string `json:"Region"`
			Type        string `json:"Type"`
			IP          string `json:"IP"`
			Gateway     string `json:"Gateway"`
			Block       string `json:"Block"`
			UID         int    `json:"UID"`
			Vid         int    `json:"VID"`
			CreateDate  int    `json:"CreateDate"`
			Flags       string `json:"Flags"`
			VlanID      int    `json:"VlanID"`
			Description string `json:"Description"`
		} `json:"EIPList"`
		RenewPointPrice struct {
			Num7  int `json:"7"`
			Num31 int `json:"31"`
		} `json:"RenewPointPrice"`
		FastInstallAppTask []any `json:"FastInstallAppTask"`
		VNets              []any `json:"VNets"`
	} `json:"data"`
}

// RCS创建备份
type CreateRcsBackupRequest struct {
	Label string `json:"label"` // 备份名称
}

// RCS重装系统
type ReinstallRcsRequest struct {
	AppVars []struct {
		AppID int  `json:"app_id"`
		Retry bool `json:"retry"` // 重发之前的任务,此项存在时,无需传入参数
		Vars  any  `json:"vars"`
	} `json:"app_vars"` // 当空数组时,进行单次任务下发
	OsID     int  `json:"os_id"`     // 系统ID
	ResetOsd bool `json:"reset_osd"` // 重置系统盘容量
}

// RCS管理弹性云盘
type RcsManagesElasticCloudDisksRequest struct {
	Actions []struct {
		Type   string `json:"type"`   // 操作类型: expand: 扩容, create: 创建
		Action any    `json:"action"` // 操作参数,RcsManagesElasticCloudDisksExpand或RcsManagesElasticCloudDisksCreate
	} `json:"actions"`
}

// RCS管理弹性云盘-扩容
type RcsManagesElasticCloudDisksExpand struct {
	EdiskID  int  `json:"edisk_id"`   // 弹性云盘ID
	SizeInGb int  `json:"size_in_gb"` // 操作容量
	Backup   bool `json:"backup"`     // 支持备份
}

// RCS管理弹性云盘-创建
type RcsManagesElasticCloudDisksCreate struct {
	SizeInGb int    `json:"size_in_gb"` // 操作容量
	DiskType string `json:"disk_type"`  // 磁盘类型(ssd/hdd)
	Backup   bool   `json:"backup"`     // 支持备份
	Tag      string `json:"tag"`        // 标签
}

// 创建并绑定弹性IP到RCS
type CreateAndBindIpToRcsRequest struct {
	WithFlags  string `json:"with_flags"`   // IP特征(可选): 应该是高防: us_ddosip -> 美国高防，nb_ddosip -> 宁波高防
	WithIPNum  int    `json:"with_ip_num"`  // IP数量
	WithIPType string `json:"with_ip_type"` // ipv4/ipv6
}

// 更换IP
type ChangeRcsIPRequest struct {
	DisableOldIPReason string `json:"disable_old_ip_reason"` // 可选
	IP                 string `json:"ip"`                    // IP地址
	ToIP               string `json:"to_ip"`                 // 可选
}

// 放弃IP
type DisCardRcsIPRequest struct {
	IP string `json:"ip"`
}

// 防火墙规则列表
type RcsFirewallRuleList struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords int `json:"TotalRecords"`
		Records      []struct {
			ID            int    `json:"ID"`
			VID           int    `json:"v_id"`
			IsEnable      bool   `json:"is_enable"`
			Pos           int    `json:"pos"`
			SourceAddress string `json:"source_address"`
			DestPort      string `json:"dest_port"`
			SourcePort    string `json:"source_port"`
			Protocol      string `json:"protocol"`
			Action        string `json:"action"`
			Description   string `json:"description"`
		} `json:"Records"`
	} `json:"data"`
}

// 创建/设置防火墙规则
type SetRcsFirewallRuleRequest struct {
	Action        string `json:"action"`         // 动作，accept/drop，接受或者丢弃
	Description   string `json:"description"`    // 备注(可选)
	DestPort      string `json:"dest_port"`      // 代表本机的目的端口，可以用-来链接，空白代表所有端口(可选)
	ID            int    `json:"id"`             // 规则ID(可选)
	IsEnable      bool   `json:"is_enable"`      //是否启用该规则(可选)
	Protocol      string `json:"protocol"`       // 协议，udp/tcp/icmp，空白代表所有(可选)
	SourceAddress string `json:"source_address"` // 代表来源的地址，可以用-链接范围，或者用逗号来分割多个地址，可以使用网络，CIDR格式，空则代表所有地址(可选)
	SourcePort    string `json:"source_port"`    // 一般不填(防反射)(可选)
}

// 移动防火墙规则优先级
type MobileRcsFirewallRulePriorityRequest struct {
	NewPos int `json:"newPos"`
}

// RCS监控数据
/* 格式如下:
{
    "code": 200,
    "data": {
        "Columns": [
            "time",
            "cpu",
            "freemem",
            "diskwrite",
            "diskread",
            "netout",
            "netin"
        ],
        "Values": [
            [
                1762004960,
                0.0310156592414152,
                null,
                39321.6,
                0,
                4297.5,
                5647.1
            ]
		}
}
*/
type RcsMonitoringData struct {
	Code int `json:"code"`
	Data struct {
		Columns []string    `json:"Columns"`
		Values  [][]float64 `json:"Values"`
	} `json:"data"`
}

// 添加NAT端口映射
type AddRcsNatPortMappingRequest struct {
	PortIn   int    `json:"port_in"`   // >= 1 <= 65535
	PortOut  int    `json:"port_out"`  // >= 10000 <= 60000
	PortType string `json:"port_type"` // tcp/udp/tcp_udp
	Tag      string `json:"tag"`       // 可选
}

// RCS续费
type RenewRcsRequest struct {
	Duration     int `json:"duration"`       // 续费时长(天)
	WithCouponID int `json:"with_coupon_id"` // 优惠券ID
}

// RCS自动续费选项
type EnableRcsAutoRenewRequest struct {
	AutoRenewOption bool `json:"auto_renew_option"`
}

// RCS重置密码
type ResetRcsPasswordRequest struct {
	Password string `json:"password"` // 新密码,留空则自动生成
}

// 设置RCS标签
type SetRcsTagRequest struct {
	TagName string `json:"tag_name"`
}

// RCS充流量
type ChargeRcsTraficRequest struct {
	TrafficInGb int `json:"traffic_in_gb"` // 充多少G
}

// RCS限流
type LimitRcsTrafficRequest struct {
	DayTrafficInGb int `json:"day_traffic_in_gb"` // 日流量阈值(G)
	TrafficLimit   int `json:"traffic_limit"`     // 限制带宽(M)
}

// RCS升级
type UpgradeRcsRequest struct {
	DestPlan     int `json:"dest_plan"`      // 升级到的套餐ID
	WithCouponID int `json:"with_coupon_id"` // 优惠券ID,默认为0
}

// RCS操作系统列表
type RcsOSList struct {
	Code int `json:"code"`
	Data []struct {
		ID             int    `json:"id"`      // 系统ID
		Region         string `json:"region"`  // 地域
		Subtype        string `json:"subtype"` // 类型(kvm)
		Machine        string `json:"machine"` // unknown
		Name           string `json:"name"`    // 英文名
		Version        string `json:"version"` // 版本
		SyncStatus     string `json:"sync_status"`
		OsType         string `json:"os_type"`         // 系统类型(windows/linux)
		ChineseName    string `json:"chinese_name"`    // 中文名
		Icon           string `json:"icon"`            // 图标
		IsWithBbr      bool   `json:"is_with_bbr"`     // 是否支持BBR
		IsEol          bool   `json:"is_eol"`          // 是否已过时
		IsAvailable    bool   `json:"is_available"`    // 是否可用
		Order          int    `json:"order"`           // 排序
		LatestFilename string `json:"latest_filename"` // 最新文件名
		NoVMAgent      bool   `json:"no_vm_agent"`     // 是否无虚拟机Agent
	} `json:"data"`
}

// RCS使用情况列表
type RcsUsageList struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords int `json:"TotalRecords"`
		Records      []struct {
			ID     int    `json:"ID"`
			Tag    string `json:"Tag"`
			IP     string `json:"IP"`
			System string `json:"System"`
			Region string `json:"Region"`
			Status string `json:"Status"`
			Usage  struct {
				CPU     float64 `json:"CPU"`
				MaxMem  int64   `json:"MaxMem"`
				FreeMem int64   `json:"FreeMem"`
				UsedMem int     `json:"UsedMem"`
				Disks   struct {
					NAMING_FAILED struct {
						Total int64 `json:"Total"`
						Used  int64 `json:"Used"`
					} `json:"/"`
				} `json:"Disks"`
				DiskRead    int         `json:"DiskRead"`
				DiskWrite   float64     `json:"DiskWrite"`
				NetOut      float64     `json:"NetOut"`
				NetIn       float64     `json:"NetIn"`
				SmartHealth interface{} `json:"SmartHealth"`
				SmartTemp   int         `json:"SmartTemp"`
				UpdateTime  int         `json:"UpdateTime"`
			} `json:"Usage"`
		} `json:"Records"`
	} `json:"data"`
}

// 工单列表
type WorkorderList struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords int `json:"TotalRecords"`
		Records      []struct {
			ID                 int    `json:"ID"`  // 工单ID
			UID                int    `json:"UID"` // 用户ID
			ExtUserInfo        string `json:"ExtUserInfo"`
			UserName           string `json:"UserName"` // 用户名
			UserEmail          string `json:"UserEmail"`
			UserVip            string `json:"UserVip"`            // 用户vip等级
			UserIcon           string `json:"UserIcon"`           // 用户头像
			Title              string `json:"Title"`              // 工单标题
			Type               string `json:"Type"`               // 工单类型
			RelatedProductType string `json:"RelatedProductType"` // 关联产品类型
			RelatedProductID   int    `json:"RelatedProductID"`   // 关联产品id
			IsUrgent           int    `json:"IsUrgent"`           // 是否为紧急工单
			Status             string `json:"Status"`             // 状态(finished/answered/waiting)
			Time               int    `json:"Time"`               // 工单创建时间
			LastTime           int    `json:"LastTime"`
			WaitBeginTime      int    `json:"WaitBeginTime"`
			AssistID           int    `json:"AssistID"` // 客服id
			AuthStatus         string `json:"AuthStatus"`
			AuthTime           int    `json:"AuthTime"`
			AuthID             int    `json:"AuthID"`
		} `json:"Records"`
	} `json:"data"`
}

// 工单详情
type WorkorderDetail struct {
	Code int `json:"code"`
	Data struct {
		ID                 int    `json:"ID"`
		UID                int    `json:"UID"`
		ExtUserInfo        string `json:"ExtUserInfo"`
		UserName           string `json:"UserName"`
		UserEmail          string `json:"UserEmail"`
		UserVip            string `json:"UserVip"`
		UserIcon           string `json:"UserIcon"`
		Title              string `json:"Title"`
		Content            string `json:"Content"`
		Type               string `json:"Type"`
		RelatedProductType string `json:"RelatedProductType"`
		RelatedProductID   int    `json:"RelatedProductID"`
		IsUrgent           int    `json:"IsUrgent"`
		Status             string `json:"Status"`
		Time               int    `json:"Time"`
		LastTime           int    `json:"LastTime"`
		WaitBeginTime      int    `json:"WaitBeginTime"`
		AssistID           int    `json:"AssistID"`
		AuthStatus         string `json:"AuthStatus"`
		AuthTime           int    `json:"AuthTime"`
		AuthID             int    `json:"AuthID"`
		Discuss            []struct {
			ID             int    `json:"ID"`
			IsAssist       bool   `json:"IsAssist"`
			UID            int    `json:"UID"`
			UserName       string `json:"UserName"`
			UserEmail      string `json:"UserEmail"`
			UserVip        string `json:"UserVip"`
			UserIcon       string `json:"UserIcon"`
			Content        string `json:"Content"`
			Time           int    `json:"Time"`
			WaitTime       int    `json:"WaitTime"`
			LastEditedTime int    `json:"LastEditedTime"`
			IsScored       bool   `json:"IsScored"`
		} `json:"Discuss"`
	} `json:"data"`
}

// 创建工单
type CreateWorkerorderRequest struct {
	Content            string `json:"content"`
	IsAuthed           bool   `json:"is_authed"`            // 是否需要授权(可选)
	IsUrgent           int    `json:"is_urgent"`            // 是否为紧急工单
	RelatedProductID   int    `json:"related_product_id"`   // 关联产品类型(可选)
	RelatedProductType string `json:"related_product_type"` // 关联产品id(可选)
	Title              string `json:"title"`
	Type               string `json:"type"` // 工单类型
}

// 创建工单响应
type CreateWorkerorderResponse struct {
	Code int `json:"code"`
	Data struct {
		ID                 int    `json:"ID"`
		UID                int    `json:"UID"`
		ExtUserInfo        string `json:"ExtUserInfo"`
		UserName           string `json:"UserName"`
		UserEmail          string `json:"UserEmail"`
		UserVip            string `json:"UserVip"`
		UserIcon           string `json:"UserIcon"`
		Title              string `json:"Title"`
		Content            string `json:"Content"`
		Type               string `json:"Type"`
		RelatedProductType string `json:"RelatedProductType"`
		RelatedProductID   int    `json:"RelatedProductID"`
		IsUrgent           int    `json:"IsUrgent"`
		Status             string `json:"Status"`
		Time               int    `json:"Time"`
		LastTime           int    `json:"LastTime"`
		WaitBeginTime      int    `json:"WaitBeginTime"`
		AssistID           int    `json:"AssistID"`
		AuthStatus         string `json:"AuthStatus"`
		AuthTime           int    `json:"AuthTime"`
		AuthID             int    `json:"AuthID"`
	} `json:"data"`
}

// 产品授权请求
type ProductAuthRequest struct {
	ProductID   int    `json:"product_id"`   // 产品ID
	ProductType string `json:"product_type"` // 产品类型(rvh/rcs/rgs/rbm/ros)
}

// 回复工单请求
type ReplyWorkerorderRequest struct {
	Content string `json:"content"` // 回复内容
}

// 工单打分
type ScoreWorkerorderRequest struct {
	Score     int    `json:"score"`      // 分数(1-5)
	Reason    string `json:"reason"`     // (可选)
	IsSolved  bool   `json:"is_solved"`  // 是否解决(可选)
	DiscussID int    `json:"discuss_id"` // 回复ID(可选)
	Aid       int    `json:"aid"`        // 客服id(可选)
}

// 工单打分详情
type ScoreWorkerorderDetail struct {
	Code int `json:"code"`
	Data struct {
		UID         int    `json:"uid"`
		Aid         int    `json:"aid"`
		OrderID     int    `json:"order_id"`
		DiscussID   int    `json:"discuss_id"`
		Score       int    `json:"score"`
		Reason      string `json:"reason"`
		IsSolved    bool   `json:"is_solved"`
		DiscussTime int    `json:"DiscussTime"`
		ScoreTime   int    `json:"score_time"`
	} `json:"data"`
}

// 工单状态
type WorkorderStatus struct {
	Code int `json:"code"`
	Data struct {
		LastTime int    `json:"LastTime"`
		Status   string `json:"Status"`
	} `json:"data"`
}

// 设置工单状态请求
type SetWorkorderStatusRequest struct {
	Status string `json:"status"`
}

// 云应用区域信息
type RcaRegionInfo struct {
	Code int `json:"code"`
	Data []struct {
		Id                     int    `json:"id"`
		Name                   string `json:"name"`
		Chinese_name           string `json:"chinese_name"`
		Website_service_domain string `json:"website_service_domain"`
		SftpServiceDomain      string `json:"sftp_service_domain"`
		PublicServiceDomain    string `json:"public_service_domain"`
		PriceInfo              struct {
			Cpu     float64 `json:"cpu"`
			Memory  float64 `json:"memory"`
			Ipv4    float64 `json:"ipv4"`
			Traffic float64 `json:"traffic"`
			Disk    float64 `json:"disk"`
		} `json:"price_info"`
	} `json:"data"`
}

// 云应用雨点余额使用情况
type RcaRaindropUsage struct {
	Code int `json:"code"`
	Data struct {
		ExpectedRemainDays     int     `json:"expected_remain_days"`      // 预计剩余天数
		LastMonthUsage         float64 `json:"last_month_usage"`          // 上月使用量
		ExpectedNextMonthUsage float64 `json:"expected_next_month_usage"` // 预计下月使用量
		FreeTrialRemainDays    int     `json:"free_trial_remain_days"`    // 剩余免费试用天数
		IsBeforeFirstPayment   bool    `json:"is_before_first_payment"`
	} `json:"data"`
}

// 雨点套餐列表
type RaindropPlansList struct {
	Code int `json:"code"`
	Data []struct {
		Id        int    `json:"id"`
		Amount    int    `json:"amount"`
		Price     int    `json:"price"`
		IsSelling bool   `json:"is_selling"`
		PlanName  string `json:"plan_name"`
		Chinese   string `json:"chinese"`
	} `json:"data"`
}

// 雨点消费历史
type RaindropConsumeLog struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords int `json:"TotalRecords"`
		Records      []struct {
			Id        int     `json:"id"`
			Uid       int     `json:"uid"`
			Time      int     `json:"time"`
			Type      string  `json:"type"`
			ProductID int     `json:"product_id"`
			Amount    float64 `json:"amount"`
			Data      struct {
				BasicPrice   float64 `json:"basic_price"`
				TrafficBytes int     `json:"traffic_bytes"`
				TrafficPrice float64 `json:"traffic_price"`
			} `json:"data"`
		} `json:"Records"`
	} `json:"data"`
}

// 云应用购买雨点请求
type BuyRaindropRequest struct {
	PlanID       int `json:"plan_id"`
	WithCouponID int `json:"with_coupon_id"`
}

// 开通云应用产品请求
type ActivateRcaRequest struct {
	RegionID int `json:"region_id"`
}

// 雨点余额
type RaindropBalance struct {
	Code int     `json:"code"`
	Data float64 `json:"data"`
}

// 创建云应用项目请求
type CreateRcaProjectRequest struct {
	ChargeType  string `json:"charge_type"`  // 计费类型: 动态计费: elastic
	CPULimit    int    `json:"cpu_limit"`    //（仅限package模式）CPU限制（毫核），1核心=1000，0.1核=100
	DiskSize    int    `json:"disk_size"`    // 磁盘大小（GiB）
	Ipv4Count   int    `json:"ipv4_count"`   // 要添加的IPv4地址数量
	Ipv6Count   int    `json:"ipv6_count"`   // 要添加的IPv6地址数量
	MemoryLimit int    `json:"memory_limit"` // （仅限package模式）内存限制（MiB）
	Name        string `json:"name"`         // 名称

	RegionID int `json:"region_id"` // 部署区域
}

// 创建云应用项目响应
type CreateRcaProjectResponse struct {
	Code int `json:"code"`
	Data struct {
		ID         int    `json:"ID"`
		UID        int    `json:"UID"`
		PlanID     int    `json:"PlanID"`
		CreateDate int    `json:"CreateDate"`
		NodeUUID   string `json:"NodeUUID"`
		Node       struct {
			UUID              string `json:"UUID"`
			AuthKey           string `json:"AuthKey"`
			Region            string `json:"Region"`
			IPRegion          string `json:"IpRegion"`
			Machine           string `json:"Machine"`
			Product           string `json:"Product"`
			Subtype           string `json:"Subtype"`
			ChineseName       string `json:"ChineseName"`
			PhysicalNode      string `json:"PhysicalNode"`
			Config            string `json:"Config"`
			Stock             any    `json:"Stock"`
			StatusData        string `json:"StatusData"`
			ShowMonitorData   string `json:"ShowMonitorData"`
			UpdateTime        string `json:"UpdateTime"`
			GitRepositoryName string `json:"GitRepositoryName"`
			CertifyRequired   bool   `json:"CertifyRequired"`
			IsDisableBackup   bool   `json:"IsDisableBackup"`
			IsHidden          bool   `json:"IsHidden"`
			NodeName          string `json:"NodeName"`
		} `json:"Node"`
		Status                  string      `json:"Status"`
		StopReason              string      `json:"StopReason"`
		RewardPointsToBeCollect int         `json:"RewardPointsToBeCollect"`
		Tag                     string      `json:"Tag"`
		ExpDate                 int         `json:"ExpDate"`
		ExpireNotice            int         `json:"ExpireNotice"`
		AutoRenew               bool        `json:"AutoRenew"`
		UnsubscribeAble         bool        `json:"UnsubscribeAble"`
		Try                     bool        `json:"Try"`
		Name                    string      `json:"name"`
		RegionID                int         `json:"region_id"`
		Region                  interface{} `json:"region"`
		Namespace               string      `json:"namespace"`
		APIToken                string      `json:"APIToken"`
		ResourceLimits          struct {
			MaxCPU    int `json:"max_cpu"`
			MaxMemory int `json:"max_memory"`
		} `json:"resource_limits"`
		VolumeSize     int    `json:"volume_size"`
		ChargeType     string `json:"charge_type"`
		HourlyPrice    int    `json:"hourly_price"`
		NextChargeTime int    `json:"next_charge_time"`
		BackupTarget   struct {
			Type              string `json:"type"`
			S3Endpoint        string `json:"s3_endpoint"`
			S3Bucket          string `json:"s3_bucket"`
			S3AccessKey       string `json:"s3_access_key"`
			S3SecretKey       string `json:"s3_secret_key"`
			S3BackupDirectory string `json:"s3_backup_directory"`
		} `json:"backup_target"`
		SftpSetting   any  `json:"sftp_setting"`
		IdleAlarmFlag bool `json:"idle_alarm_flag"`
		PaymentDueEnd int  `json:"payment_due_end"`
	} `json:"data"`
}

// 云应用项目列表
type RcaProjectList struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords int `json:"TotalRecords"`
		Records      []struct {
			ID         int    `json:"ID"`
			UID        int    `json:"UID"`
			PlanID     int    `json:"PlanID"`
			CreateDate int    `json:"CreateDate"`
			NodeUUID   string `json:"NodeUUID"`
			Node       struct {
				UUID              string      `json:"UUID"`
				AuthKey           string      `json:"AuthKey"`
				Region            string      `json:"Region"`
				IPRegion          string      `json:"IpRegion"`
				Machine           string      `json:"Machine"`
				Product           string      `json:"Product"`
				Subtype           string      `json:"Subtype"`
				ChineseName       string      `json:"ChineseName"`
				PhysicalNode      string      `json:"PhysicalNode"`
				Config            string      `json:"Config"`
				Stock             interface{} `json:"Stock"`
				StatusData        string      `json:"StatusData"`
				ShowMonitorData   string      `json:"ShowMonitorData"`
				UpdateTime        string      `json:"UpdateTime"`
				GitRepositoryName string      `json:"GitRepositoryName"`
				CertifyRequired   bool        `json:"CertifyRequired"`
				IsDisableBackup   bool        `json:"IsDisableBackup"`
				IsHidden          bool        `json:"IsHidden"`
				NodeName          string      `json:"NodeName"`
			} `json:"Node"`
			Status                  string `json:"Status"`
			StopReason              string `json:"StopReason"`
			RewardPointsToBeCollect int    `json:"RewardPointsToBeCollect"`
			Tag                     string `json:"Tag"`
			ExpDate                 int    `json:"ExpDate"`
			ExpireNotice            int    `json:"ExpireNotice"`
			AutoRenew               bool   `json:"AutoRenew"`
			UnsubscribeAble         bool   `json:"UnsubscribeAble"`
			Try                     bool   `json:"Try"`
			Name                    string `json:"name"`
			RegionID                int    `json:"region_id"`
			Region                  struct {
				ID                   int    `json:"id"`
				Name                 string `json:"name"`
				ChineseName          string `json:"chinese_name"`
				WebsiteServiceDomain string `json:"website_service_domain"`
				SftpServiceDomain    string `json:"sftp_service_domain"`
				PublicServiceDomain  string `json:"public_service_domain"`
				PriceInfo            struct {
					CPU     float64 `json:"cpu"`
					Memory  float64 `json:"memory"`
					Ipv4    float64 `json:"ipv4"`
					Traffic float64 `json:"traffic"`
					Disk    float64 `json:"disk"`
				} `json:"price_info"`
			} `json:"region"`
			Namespace      string `json:"namespace"`
			APIToken       string `json:"APIToken"`
			ResourceLimits struct {
				MaxCPU    int `json:"max_cpu"`
				MaxMemory int `json:"max_memory"`
			} `json:"resource_limits"`
			UsageData struct {
				CPU             int    `json:"cpu"`
				Memory          int    `json:"memory"`
				NetOut          int    `json:"net_out"`
				NetIn           int    `json:"net_in"`
				DiskUsage       int    `json:"disk_usage"`
				TrafficToday    int    `json:"traffic_today"`
				Status          string `json:"status"`
				StatusReason    string `json:"status_reason"`
				AllocatedCPU    int    `json:"allocated_cpu"`
				AllocatedMemory int    `json:"allocated_memory"`
				AppCount        int    `json:"app_count"`
				WebsiteCount    int    `json:"website_count"`
				DatabaseCount   int    `json:"database_count"`
				Ipv4Count       int    `json:"ipv4_count"`
				HealthyPods     int    `json:"healthy_pods"`
				UnhealthyPods   int    `json:"unhealthy_pods"`
			} `json:"usage_data"`
			VolumeSize     int    `json:"volume_size"`
			ChargeType     string `json:"charge_type"`
			HourlyPrice    int    `json:"hourly_price"`
			NextChargeTime int    `json:"next_charge_time"`
			BackupTarget   struct {
				Type              string `json:"type"`
				S3Endpoint        string `json:"s3_endpoint"`
				S3Bucket          string `json:"s3_bucket"`
				S3AccessKey       string `json:"s3_access_key"`
				S3SecretKey       string `json:"s3_secret_key"`
				S3BackupDirectory string `json:"s3_backup_directory"`
			} `json:"backup_target"`
			SftpSetting   interface{} `json:"sftp_setting"`
			IdleAlarmFlag bool        `json:"idle_alarm_flag"`
			PaymentDueEnd int         `json:"payment_due_end"`
		} `json:"Records"`
	} `json:"data"`
}

const (
	TBPass   = "关注成功"
	BiliPass = "雨云爱你"
	QQPass   = "我爱雨云"
)
