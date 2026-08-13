package rgs

// 创建游戏云请求
type CreateRgsRequest struct {
	AppVars []struct {
		AppID int  `json:"app_id"`
		Retry bool `json:"retry"` // 重发之前的任务,此项存在时,无需传入参数
		Vars  struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"vars"`
	} `json:"app_vars"` // 当空数组时,进行单次任务下发(可选)
	Config struct {
		Allocation int `json:"allocation"`
		Backup     int `json:"backup"`
		BaseDisk   int `json:"base_disk"`
		CPU        int `json:"cpu"`
		DataDisk   int `json:"data_disk"`
		Database   int `json:"database"`
		Memory     int `json:"memory"`
		NetIn      int `json:"net_in"`
		NetOut     int `json:"net_out"`
	} `json:"config"`
	CPULimitMode bool   `json:"cpu_limit_mode"`
	Duration     int    `json:"duration"`    // 创建时长(月)
	EggTypeID    int    `json:"egg_type_id"` // 游戏类型
	NodeUUID     string `json:"node_uuid"`
	OnlineMode   bool   `json:"online_mode"`
	OsID         int    `json:"os_id"`
	PanelUser    string `json:"panel_user"` // 游戏云面板用户
	PayMode      string `json:"pay_mode"`
	PlanID       int    `json:"plan_id"`
	Subtype      string `json:"subtype"` // kvm/mcsm
	Try          bool   `json:"try"`
	WithCouponID int    `json:"with_coupon_id"`
	WithEipFlags string `json:"with_eip_flags"` // 是否开启高防，us_ddosip -> 美国高防，nb_ddosip -> 宁波高防
	WithEipNum   int    `json:"with_eip_num"`
	WithEipType  string `json:"with_eip_type"`
	Zone         string `json:"zone"`
}

// 创建游戏云响应
type CreateRgsResponse struct {
	Code int `json:"code"`
	Data struct {
		ExpDate         int    `json:"ExpDate"`      // 过期时间
		ExpireNotice    int    `json:"ExpireNotice"` // 到期提醒
		AutoRenew       bool   `json:"AutoRenew"`    // 自动续费
		UnsubscribeAble bool   `json:"UnsubscribeAble"`
		Try             bool   `json:"Try"` // 试用
		ID              int    `json:"ID"`
		UID             int    `json:"UID"`
		PlanID          int    `json:"PlanID"`     // 套餐ID
		CreateDate      int    `json:"CreateDate"` // 创建时间
		NodeUUID        string `json:"NodeUUID"`   // 节点UUID
		Node            struct {
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
		} `json:"Node"` // 节点信息
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
		NatPublicIP             string `json:"NatPublicIP"`     // Nat IP
		NatPublicDomain         string `json:"NatPublicDomain"` // Nat 域名
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
		Plan                    struct {
			ID               int    `json:"id"`
			Region           string `json:"region"`
			Subtype          string `json:"subtype"`
			PlanName         string `json:"plan_name"`
			Machine          string `json:"machine"`
			ChargeType       string `json:"charge_type"`
			Chinese          string `json:"chinese"`
			IsFree           bool   `json:"is_free"`
			PointRenewPrice  any    `json:"point_renew_price"`
			IsSelling        bool   `json:"is_selling"`
			StockDiscount    int    `json:"stock_discount"`
			EipStockDiscount int    `json:"eip_stock_discount"`
			IPPrices         any    `json:"ip_prices"`
			IPSelling        any    `json:"ip_selling"`
			CPUPrice         int    `json:"cpu_price"`
			MemoryPrice      int    `json:"memory_price"`
			NetInPrice       int    `json:"net_in_price"`
			NetOutPrice      int    `json:"net_out_price"`
			BaseDiskPrice    int    `json:"base_disk_price"`
			DataDiskPrice    int    `json:"data_disk_price"`
			Config           []struct {
				CPU         int `json:"cpu"`
				Memory      int `json:"memory"`
				NetIn       int `json:"net_in,omitempty"`
				CPUMax      int `json:"cpu_max,omitempty"`
				CPUMin      int `json:"cpu_min,omitempty"`
				NetOut      int `json:"net_out"`
				BaseDisk    int `json:"base_disk,omitempty"`
				DataDisk    int `json:"data_disk,omitempty"`
				BasePrice   int `json:"base_price"`
				MemoryMax   int `json:"memory_max"`
				MemoryMin   int `json:"memory_min"`
				NetInMax    int `json:"net_in_max,omitempty"`
				NetInMin    int `json:"net_in_min,omitempty"`
				NetOutMax   int `json:"net_out_max,omitempty"`
				NetOutMin   int `json:"net_out_min"`
				BaseDiskMax int `json:"base_disk_max,omitempty"`
				BaseDiskMin int `json:"base_disk_min,omitempty"`
				DataDiskMax int `json:"data_disk_max,omitempty"`
				DataDiskMin int `json:"data_disk_min,omitempty"`
			} `json:"config"` // 这里不知道为什么要返还一堆套餐列表
			AutoRestock      int     `json:"auto_restock"`
			AvailableStock   int     `json:"available_stock"`
			CPUPointDefault  int     `json:"cpu_point_default"`
			CPUPointConsume  int     `json:"cpu_point_consume"`
			CPUPointPrice    float64 `json:"cpu_point_price"`
			CPUBase          float64 `json:"cpu_base"`
			CPUMax           int     `json:"cpu_max"`
			EipPrice         int     `json:"eip_price"`
			DefencePrice     int     `json:"defence_price"`
			AllocationPrice  int     `json:"allocation_price"`
			DatabasePrice    int     `json:"database_price"`
			BackupPrice      int     `json:"backup_price"`
			DailyModeSupport bool    `json:"daily_mode_support"`
			DailyPriceScale  int     `json:"daily_price_scale"`
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
		} `json:"OsInfo"` // 系统信息
		CPU                   int    `json:"CPU"`
		Memory                int    `json:"Memory"`
		BaseDisk              int    `json:"BaseDisk"`
		DataDisk              int    `json:"DataDisk"`
		InitedDate            int    `json:"InitedDate"`
		Allocation            int    `json:"Allocation"`
		Database              int    `json:"Database"`
		Backup                int    `json:"Backup"`
		CPULimitMode          bool   `json:"CpuLimitMode"`
		CPULimitStatus        bool   `json:"CpuLimitStatus"`
		CPUPoint              int    `json:"CpuPoint"`
		DailyMode             bool   `json:"DailyMode"`
		RBSKeepLast           int    `json:"RBSKeepLast"`
		RBSAutoBackup         string `json:"RBSAutoBackup"`
		RBSLastAutoBackupDate int    `json:"RBSLastAutoBackupDate"`
		EggTypeID             int    `json:"EggTypeId"`
		EggType               any    `json:"EggType"`
		PteroUserName         string `json:"PteroUserName"`
		PteroUser             any    `json:"PteroUser"`
		ServerID              int    `json:"ServerID"`
		AllocationID          int    `json:"AllocationID"`
		McsmUserName          string `json:"McsmUserName"`
		McsmUser              any    `json:"McsmUser"` // Mcsm用户
		ServerUUID            string `json:"ServerUUID"`
		DaemonUUID            string `json:"DaemonUUID"`
		GameInfo              any    `json:"GameInfo"`
	} `json:"data"`
}

// 游戏云列表信息
// 由于响应实在过于庞大，我们只维护部分必要的响应，如有扩展需求，请在项目的data文件夹下寻找响应实例自行解码
type RgsList struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords int `json:"TotalRecords"`
		Records      []struct {
			ExpDate                 int    `json:"ExpDate"`
			ExpireNotice            int    `json:"ExpireNotice"`
			AutoRenew               bool   `json:"AutoRenew"`
			UnsubscribeAble         bool   `json:"UnsubscribeAble"`
			Try                     bool   `json:"Try"`
			ID                      int    `json:"ID"`
			UID                     int    `json:"UID"`
			PlanID                  int    `json:"PlanID"`
			CreateDate              int    `json:"CreateDate"`
			Status                  string `json:"Status"`
			StopReason              string `json:"StopReason"`
			RewardPointsToBeCollect int    `json:"RewardPointsToBeCollect"` // 待领取的积分
			Tag                     string `json:"Tag"`
			OsID                    int    `json:"OsID"`
			OsName                  string `json:"OsName"`
			HostName                string `json:"HostName"`
			DefaultPass             string `json:"DefaultPass"`
			MainIPv4                string `json:"MainIPv4"`
			IntIPv4                 string `json:"IntIPv4"`
			UsageData               struct {
				CPU        int `json:"CPU"`
				Mem        int `json:"Mem"`
				MemUsage   int `json:"MemUsage"`
				DiskRead   int `json:"DiskRead"`
				DiskWrite  int `json:"DiskWrite"`
				Disk       int `json:"Disk"`
				NetOut     int `json:"NetOut"`
				NetIn      int `json:"NetIn"`
				UpdateTime int `json:"UpdateTime"`
			} `json:"UsageData"`
			Zone                  string `json:"Zone"`
			NatPublicIP           string `json:"NatPublicIP"`
			NatPublicDomain       string `json:"NatPublicDomain"`
			NATSpareDomain        string `json:"NATSpareDomain"`
			NetIn                 int    `json:"NetIn"`
			NetOut                int    `json:"NetOut"`
			NowNetIn              int    `json:"NowNetIn"`
			NowNetOut             int    `json:"NowNetOut"`
			NetMode               string `json:"NetMode"`
			BridgeSyncing         bool   `json:"BridgeSyncing"`
			VnetID                int    `json:"VnetID"`
			UpdateTime            int    `json:"UpdateTime"`
			FwSyncTime            int    `json:"FwSyncTime"`
			FwMode                string `json:"FwMode"`
			AbCPULimit            int    `json:"AbCpuLimit"`
			AbNetLimit            int    `json:"AbNetLimit"`
			AbWhiteReason         string `json:"AbWhiteReason"`
			OsInfo                any    `json:"OsInfo"`
			CPU                   int    `json:"CPU"`
			Memory                int    `json:"Memory"`
			BaseDisk              int    `json:"BaseDisk"`
			DataDisk              int    `json:"DataDisk"`
			InitedDate            int    `json:"InitedDate"`
			Allocation            int    `json:"Allocation"`
			Database              int    `json:"Database"`
			CPULimitMode          bool   `json:"CpuLimitMode"`
			CPULimitStatus        bool   `json:"CpuLimitStatus"`
			CPUPoint              int    `json:"CpuPoint"`
			DailyMode             bool   `json:"DailyMode"` // 是否日付模式
			RBSKeepLast           int    `json:"RBSKeepLast"`
			RBSAutoBackup         string `json:"RBSAutoBackup"`
			RBSLastAutoBackupDate int    `json:"RBSLastAutoBackupDate"`
			McsmUserName          string `json:"McsmUserName"`
			McsmUser              struct {
				Name      string `json:"name"`
				Password  string `json:"password"`
				UserID    int    `json:"user_id"`
				PanelUUID string `json:"panel_uuid"`
			} `json:"McsmUser"`
		} `json:"Records"`
	} `json:"data"`
}

// 游戏云详情
// 由于响应实在过于庞大，我们只维护部分必要的响应，如有扩展需求，请在项目的data文件夹下寻找响应实例自行解码
type RgsDetail struct {
	Code int `json:"code"`
	Data struct {
		Data struct {
			ExpDate                 int    `json:"ExpDate"` // 到期时间
			ExpireNotice            int    `json:"ExpireNotice"`
			AutoRenew               bool   `json:"AutoRenew"` // 是否自动续费
			UnsubscribeAble         bool   `json:"UnsubscribeAble"`
			Try                     bool   `json:"Try"` // 是否试用
			ID                      int    `json:"ID"`
			UID                     int    `json:"UID"`
			PlanID                  int    `json:"PlanID"`
			CreateDate              int    `json:"CreateDate"`
			Status                  string `json:"Status"` // 状态
			StopReason              string `json:"StopReason"`
			Node struct {
				UUID string `json:"UUID"`
				AuthKey string `json:"AuthKey"`
				Region string `json:"Region"`
				IPRegion string `json:"IpRegion"`
				Machine string `json:"Machine"`
				Product string `json:"Product"`
				Subtype string `json:"Subtype"`
				ChineseName string `json:"ChineseName"`
				PhysicalNode string `json:"PhysicalNode"`
				Config string `json:"Config"`
				Stock interface{} `json:"Stock"`
				StatusData string `json:"StatusData"`
				ShowMonitorData string `json:"ShowMonitorData"`
				UpdateTime string `json:"UpdateTime"`
				GitRepositoryName string `json:"GitRepositoryName"`
				CertifyRequired bool `json:"CertifyRequired"`
				IsDisableBackup bool `json:"IsDisableBackup"`
				IsHidden bool `json:"IsHidden"`
				NodeName string `json:"NodeName"`
			} `json:"Node"`
			RewardPointsToBeCollect int    `json:"RewardPointsToBeCollect"` // 待领取的积分
			Tag                     string `json:"Tag"`
			OsID                    int    `json:"OsID"`
			OsName                  string `json:"OsName"`
			HostName                string `json:"HostName"`
			DefaultPass             string `json:"DefaultPass"` // 默认密码
			MainIPv4                string `json:"MainIPv4"`
			IntIPv4                 string `json:"IntIPv4"`
			UsageData               struct {
				CPU        int `json:"CPU"`
				Mem        int `json:"Mem"`
				MemUsage   int `json:"MemUsage"`
				DiskRead   int `json:"DiskRead"`
				DiskWrite  int `json:"DiskWrite"`
				Disk       int `json:"Disk"`
				NetOut     int `json:"NetOut"`
				NetIn      int `json:"NetIn"`
				UpdateTime int `json:"UpdateTime"`
			} `json:"UsageData"` // 使用情况
			Zone            string `json:"Zone"`
			NatPublicIP     string `json:"NatPublicIP"`
			NatPublicDomain string `json:"NatPublicDomain"`
			NATSpareDomain  string `json:"NATSpareDomain"`
			NetIn           int    `json:"NetIn"`
			NetOut          int    `json:"NetOut"`
			NowNetIn        int    `json:"NowNetIn"`
			NowNetOut       int    `json:"NowNetOut"`
			NetMode         string `json:"NetMode"`
			BridgeSyncing   bool   `json:"BridgeSyncing"`
			VnetID          int    `json:"VnetID"`
			UpdateTime      int    `json:"UpdateTime"`
			FwSyncTime      int    `json:"FwSyncTime"`
			FwMode          string `json:"FwMode"`
			AbCPULimit      int    `json:"AbCpuLimit"`
			AbNetLimit      int    `json:"AbNetLimit"`
			AbWhiteReason   string `json:"AbWhiteReason"`
			OsInfo          any    `json:"OsInfo"`
			CPU             int    `json:"CPU"`
			Memory          int    `json:"Memory"`
			BaseDisk        int    `json:"BaseDisk"`
			DataDisk        int    `json:"DataDisk"`
			InitedDate      int    `json:"InitedDate"`
			Backup          int    `json:"Backup"`
			CPULimitMode    bool   `json:"CpuLimitMode"`
			CPULimitStatus  bool   `json:"CpuLimitStatus"`
			CPUPoint        int    `json:"CpuPoint"`
			DailyMode       bool   `json:"DailyMode"` // 是否日付模式
			McsmUserName    string `json:"McsmUserName"`
			McsmUser        struct {
				Name      string `json:"name"`
				Password  string `json:"password"`
				UserID    int    `json:"user_id"`
				PanelUUID string `json:"panel_uuid"`
			} `json:"McsmUser"` // mcsm 用户信息
		} `json:"Data"`
		NatList []struct {
			ID       int    `json:"ID"`
			PortIn   int    `json:"PortIn"`
			PortOut  int    `json:"PortOut"`
			PortType string `json:"PortType"`
			Tag      string `json:"Tag"`
		} `json:"NatList"` // 端口映射列表
		EIPList         any `json:"EIPList"`
		ConfigPrice     int `json:"ConfigPrice"`
		RenewPointPrice struct {
			Num7  int `json:"7"`  // 积分续费七天
			Num31 int `json:"31"` // 积分续费31天
		} `json:"RenewPointPrice"` // 积分续费
	} `json:"data"`
}

// 游戏云CPU充电请求
type ChargeRgsCPURequest struct {
	Mode   string `json:"mode"`   // 支付方式(money/point)
	Money  int    `json:"money"`  // 消耗用户余额(支付方式为point时0)
	Points int    `json:"points"` // 消耗用户积分(支付方式为money时0)
}

// MCSM面板用户
type McsmUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// 游戏云配置信息
type RgsConfig struct {
	CPU        int `json:"cpu"`
	Memory     int `json:"memory"`
	NetOut     int `json:"net_out"`
	NetIn      int `json:"net_in"`
	BaseDisk   int `json:"base_disk"`
	DataDisk   int `json:"data_disk"`
	Allocation int `json:"allocation"`
	Database   int `json:"database"`
	Backup     int `json:"backup"`
}

// 游戏云升级价格
type RgsUpgradePrice struct {
	Code int `json:"code"`
	Data struct {
		Detail struct {
			Price        int  `json:"price"`
			AgentPrice   int  `json:"agent_price"`
			StockPrice   int  `json:"stock_price"`
			DefaultPrice int  `json:"default_price"`
			CouponValue  int  `json:"coupon_value"`
			SaleReward   int  `json:"sale_reward"`
			AgentReward  int  `json:"agent_reward"`
			AgentID      int  `json:"agent_id"`
			IgnoreAgent  bool `json:"ignore_agent"`
			PerScene     struct {
				Upgrade int `json:"upgrade"`
			} `json:"per_scene"`
		} `json:"detail"`
		Price int `json:"price"`
	} `json:"data"`
}

// 游戏云更换egg(游戏类型)请求
type ChangeRgsEggRequest struct {
	EggTypeID int      `json:"egg_type_id"` // 蛋ID
	SaveDirs  []string `json:"save_dirs"`   // 要保留的目录
}
