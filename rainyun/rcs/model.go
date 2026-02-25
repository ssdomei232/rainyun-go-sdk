package rcs

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
				DiskRead    int     `json:"DiskRead"`
				DiskWrite   int     `json:"DiskWrite"`
				NetOut      float64 `json:"NetOut"`
				NetIn       int     `json:"NetIn"`
				SmartHealth any     `json:"SmartHealth"`
				SmartTemp   int     `json:"SmartTemp"`
				UpdateTime  int     `json:"UpdateTime"`
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
				ID              int    `json:"id"`
				Region          string `json:"region"`
				Subtype         string `json:"subtype"`
				PlanName        string `json:"plan_name"`
				Machine         string `json:"machine"`
				ChargeType      string `json:"charge_type"`
				Chinese         string `json:"chinese"`
				IsFree          bool   `json:"is_free"`
				PointRenewPrice any    `json:"point_renew_price"`
				IsSelling       bool   `json:"is_selling"`
				Price           int    `json:"price"`
				TrafficBaseGb   int    `json:"traffic_base_gb"`
				TrafficPrice    struct {
					Num300  int `json:"300"`
					Num1024 int `json:"1024"`
					Num2048 int `json:"2048"`
				} `json:"traffic_price"`
				CPU            int    `json:"cpu"`
				Memory         int    `json:"memory"`
				NetIn          int    `json:"net_in"`
				NetOut         int    `json:"net_out"`
				IPPrices       any    `json:"ip_prices"`
				IPSelling      any    `json:"ip_selling"`
				AutoRestock    int    `json:"auto_restock"`
				AvailableStock int    `json:"available_stock"`
				GpuMemorySize  int    `json:"gpu_memory_size"`
				DgpuDevType    string `json:"dgpu_dev_type"`
				WebbarConfig   any    `json:"webbar_config"`
				NoBackup       bool   `json:"no_backup"`
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
	Duration     int    `json:"duration"`       // 创建时长(月)
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
			ID              int    `json:"id"`
			Region          string `json:"region"`
			Subtype         string `json:"subtype"`
			PlanName        string `json:"plan_name"`
			Machine         string `json:"machine"`
			ChargeType      string `json:"charge_type"`
			Chinese         string `json:"chinese"`
			IsFree          bool   `json:"is_free"`
			PointRenewPrice any    `json:"point_renew_price"`
			IsSelling       bool   `json:"is_selling"`
			Price           int    `json:"price"`
			TrafficBaseGb   int    `json:"traffic_base_gb"`
			TrafficPrice    struct {
				Num200  int `json:"200"`
				Num1024 int `json:"1024"`
			} `json:"traffic_price"`
			CPU            int    `json:"cpu"`
			Memory         int    `json:"memory"`
			NetIn          int    `json:"net_in"`
			NetOut         int    `json:"net_out"`
			IPPrices       any    `json:"ip_prices"`
			IPSelling      any    `json:"ip_selling"`
			AutoRestock    int    `json:"auto_restock"`
			AvailableStock int    `json:"available_stock"`
			GpuMemorySize  int    `json:"gpu_memory_size"`
			DgpuDevType    string `json:"dgpu_dev_type"`
			WebbarConfig   any    `json:"webbar_config"`
			NoBackup       bool   `json:"no_backup"`
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
				DiskRead    int     `json:"DiskRead"`
				DiskWrite   float64 `json:"DiskWrite"`
				NetOut      float64 `json:"NetOut"`
				NetIn       float64 `json:"NetIn"`
				SmartHealth any     `json:"SmartHealth"`
				SmartTemp   int     `json:"SmartTemp"`
				UpdateTime  int     `json:"UpdateTime"`
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
				ID              int    `json:"id"`
				Region          string `json:"region"`
				Subtype         string `json:"subtype"`
				PlanName        string `json:"plan_name"`
				Machine         string `json:"machine"`
				ChargeType      string `json:"charge_type"`
				Chinese         string `json:"chinese"`
				IsFree          bool   `json:"is_free"`
				PointRenewPrice any    `json:"point_renew_price"`
				IsSelling       bool   `json:"is_selling"`
				Price           int    `json:"price"`
				TrafficBaseGb   int    `json:"traffic_base_gb"`
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
				IPSelling      any    `json:"ip_selling"`
				AutoRestock    int    `json:"auto_restock"`
				AvailableStock int    `json:"available_stock"`
				GpuMemorySize  int    `json:"gpu_memory_size"`
				DgpuDevType    string `json:"dgpu_dev_type"`
				WebbarConfig   any    `json:"webbar_config"`
				NoBackup       bool   `json:"no_backup"`
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
			ID              int    `json:"id"`
			Region          string `json:"region"`
			Subtype         string `json:"subtype"`
			PlanName        string `json:"plan_name"`
			Machine         string `json:"machine"`
			ChargeType      string `json:"charge_type"`
			Chinese         string `json:"chinese"`
			IsFree          bool   `json:"is_free"`
			PointRenewPrice any    `json:"point_renew_price"`
			IsSelling       bool   `json:"is_selling"`
			Price           int    `json:"price"`
			TrafficBaseGb   int    `json:"traffic_base_gb"`
			TrafficPrice    any    `json:"traffic_price"`
			CPU             int    `json:"cpu"`
			Memory          int    `json:"memory"`
			NetIn           int    `json:"net_in"`
			NetOut          int    `json:"net_out"`
			IPPrices        any    `json:"ip_prices"`
			IPSelling       any    `json:"ip_selling"`
			AutoRestock     int    `json:"auto_restock"`
			AvailableStock  int    `json:"available_stock"`
			GpuMemorySize   int    `json:"gpu_memory_size"`
			DgpuDevType     string `json:"dgpu_dev_type"`
			WebbarConfig    any    `json:"webbar_config"`
			NoBackup        bool   `json:"no_backup"`
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
		NatList   []any `json:"NatList"`
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

// RCS设置备份选项请求
type RcsSetBackupOptionsRequest struct {
	AutoBackupHour   int `json:"auto_backup_hour"`   // 自动备份时间的小时
	AutoBackupMinute int `json:"auto_backup_minute"` // 自动备份时间的分钟
	KeepLast         int `json:"keep_last"`          // 保留份数(1/3/7)
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
	Duration     int `json:"duration"`       // 续费时长(月)
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

// RCS续费价格
type RCSRenewPrice struct {
	Code int `json:"code"`
	Data struct {
		Detail struct {
			Price        float64 `json:"price"`         // 价格
			AgentPrice   float64 `json:"agent_price"`   // 价格，不知道为啥和上面一样
			StockPrice   float64 `json:"stock_price"`   // 价格，不知道为啥和上面一样
			DefaultPrice int     `json:"default_price"` // unknown
			CouponValue  int     `json:"coupon_value"`
			SaleReward   int     `json:"sale_reward"`
			AgentReward  int     `json:"agent_reward"`
			AgentID      int     `json:"agent_id"`
			IgnoreAgent  bool    `json:"ignore_agent"`
			PerScene     struct {
				Eip      int     `json:"eip"`
				Renew    float64 `json:"renew"`     // 配置价格
				RenewEip float64 `json:"renew_eip"` // IP价格
			} `json:"per_scene"`
		} `json:"detail"`
		Price float64 `json:"price"`
	} `json:"data"`
}
