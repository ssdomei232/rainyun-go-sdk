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
	WithEipType  string `json:"with_eip_type"`  // IPv4(默认)/IPv6
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
		FastInstallAppTask []interface{} `json:"FastInstallAppTask"`
		VNets              []interface{} `json:"VNets"`
	} `json:"data"`
}

const (
	TBPass   = "关注成功"
	BiliPass = "雨云爱你"
	QQPass   = "我爱雨云"
)
