package ros

// 创建对象存储桶响应
type CreateRosBucketResponse struct {
	Code int `json:"code"`
	Data struct {
		ID                           int    `json:"id"`   // 存储桶ID
		Name                         string `json:"name"` // 存储桶名称
		UID                          int    `json:"uid"`  // 用户ID
		CreateDate                   int    `json:"create_date"`
		AccessKey                    string `json:"access_key"`
		SecretKey                    string `json:"secret_key"`
		IsPublicAccess               bool   `json:"is_public_access"`
		InstanceID                   int    `json:"instance_id"`
		Instance                     any    `json:"Instance"`
		StopReason                   string `json:"stop_reason"`
		DomainList                   any    `json:"DomainList"`
		SslIsEnable                  bool   `json:"ssl_is_enable"`
		SSLCertList                  any    `json:"SSLCertList"`
		SslAutoRedirect              bool   `json:"ssl_auto_redirect"`
		WafIsEnable                  bool   `json:"waf_is_enable"`
		WafGlobalQPS                 int    `json:"waf_global_qps"`
		WafPerIPQPS                  int    `json:"waf_per_ip_qps"`
		WafBlockTime                 int    `json:"waf_block_time"`
		WafGlobalJsCheck             bool   `json:"waf_global_js_check"`
		WafPerIPJsCheck              bool   `json:"waf_per_ip_js_check"`
		RefererRestrictIsEnable      bool   `json:"referer_restrict_is_enable"`
		RefererRestrictWhitelist     any    `json:"RefererRestrictWhitelist"`
		RefererRestrictBlacklist     any    `json:"RefererRestrictBlacklist"`
		RefererRestrictBypassMissing bool   `json:"referer_restrict_bypass_missing"`
		IPRestrictIsEnable           bool   `json:"ip_restrict_is_enable"`
		IPRestrictWhitelist          any    `json:"IPRestrictWhitelist"`
		IPRestrictBlacklist          any    `json:"IPRestrictBlacklist"`
		GzipIsEnable                 bool   `json:"gzip_is_enable"`
		GzipCompressLevel            int    `json:"gzip_compress_level"`
		ReadyDelete                  bool   `json:"ready_delete"`
	} `json:"data"`
}

// 创建对象存储桶请求
type CreateRosBucketRequest struct {
	BucketName string `json:"bucket_name"` // 存储桶名称
	InstanceID int    `json:"instance_id"` // 实例ID
}

// 对象存储桶详情
type RosBucketDetail struct {
	Code int `json:"code"`
	Data struct {
		Data struct {
			ID             int    `json:"id"`
			Name           string `json:"name"`
			UID            int    `json:"uid"`
			CreateDate     int    `json:"create_date"`
			AccessKey      string `json:"access_key"`       // bucket的AK
			SecretKey      string `json:"secret_key"`       // bucket的SK
			IsPublicAccess bool   `json:"is_public_access"` // 是否公开
			InstanceID     int    `json:"instance_id"`      // 实例ID
			Instance       struct {
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
				Status                  string `json:"Status"`
				StopReason              string `json:"StopReason"`
				RewardPointsToBeCollect int    `json:"RewardPointsToBeCollect"`
				Tag                     string `json:"Tag"`
				Plan                    struct {
					ID                 int     `json:"id"`
					Region             string  `json:"region"`
					Subtype            string  `json:"subtype"`
					PlanName           string  `json:"plan_name"`
					Machine            string  `json:"machine"`
					ChargeType         string  `json:"charge_type"`
					Chinese            string  `json:"chinese"`
					IsFree             bool    `json:"is_free"`
					PointRenewPrice    any     `json:"point_renew_price"`
					IsSelling          bool    `json:"is_selling"`
					Price              int     `json:"price"`
					StorageSize        int     `json:"storage_size"`
					Bandwidth          int     `json:"bandwidth"`
					ExtraTransferPrice float64 `json:"extra_transfer_price"`
					ExtraStoragePrice  float64 `json:"extra_storage_price"`
				} `json:"Plan"`
				AccessKey               string `json:"access_key"` // 实例的AK
				SecretKey               string `json:"secret_key"` // 实例的SK
				IsPublicAccess          bool   `json:"is_public_access"`
				IsEnableExtraAccounting bool   `json:"is_enable_extra_accounting"`
				LastResetDate           int    `json:"last_reset_date"`
				NextResetDate           int    `json:"next_reset_date"`
				ExpDate                 int    `json:"ExpDate"`
				ExpireNotice            int    `json:"ExpireNotice"`
				AutoRenew               bool   `json:"AutoRenew"`
				UnsubscribeAble         bool   `json:"UnsubscribeAble"`
				Try                     bool   `json:"Try"`
				PublicAPIURL            string `json:"public_api_url"`
				Buckets                 any    `json:"buckets"`
				ExtraPayTime            int    `json:"extra_pay_time"`
				Restrictions            string `json:"restrictions"`
				PublicAccessRestricted  bool   `json:"public_access_restricted"`
			} `json:"Instance"`
			StopReason                   string `json:"stop_reason"`
			DomainList                   any    `json:"DomainList"`
			SslIsEnable                  bool   `json:"ssl_is_enable"`
			SSLCertList                  any    `json:"SSLCertList"`
			SslAutoRedirect              bool   `json:"ssl_auto_redirect"`
			WafIsEnable                  bool   `json:"waf_is_enable"`
			WafGlobalQPS                 int    `json:"waf_global_qps"`
			WafPerIPQPS                  int    `json:"waf_per_ip_qps"`
			WafBlockTime                 int    `json:"waf_block_time"`
			WafGlobalJsCheck             bool   `json:"waf_global_js_check"`
			WafPerIPJsCheck              bool   `json:"waf_per_ip_js_check"`
			RefererRestrictIsEnable      bool   `json:"referer_restrict_is_enable"`
			RefererRestrictWhitelist     any    `json:"RefererRestrictWhitelist"`
			RefererRestrictBlacklist     any    `json:"RefererRestrictBlacklist"`
			RefererRestrictBypassMissing bool   `json:"referer_restrict_bypass_missing"`
			IPRestrictIsEnable           bool   `json:"ip_restrict_is_enable"`
			IPRestrictWhitelist          any    `json:"IPRestrictWhitelist"`
			IPRestrictBlacklist          any    `json:"IPRestrictBlacklist"`
			GzipIsEnable                 bool   `json:"gzip_is_enable"`
			GzipCompressLevel            int    `json:"gzip_compress_level"`
			UsageData                    struct {
				MetricsUsageTotalBytes int `json:"metrics_usage_total_bytes"`
				MetricsReceivedBytes   int `json:"metrics_received_bytes"`
				MetricsSentBytes       int `json:"metrics_sent_bytes"`
				MetricsRequests        int `json:"metrics_requests"`
				SpeedReceivedBytes     int `json:"speed_received_bytes"`
				SpeedSentBytes         int `json:"speed_sent_bytes"`
				SpeedRequests          int `json:"speed_requests"`
				UpdateTime             int `json:"UpdateTime"`
			} `json:"UsageData"`
			ReadyDelete bool `json:"ready_delete"`
		} `json:"Data"`
	} `json:"data"`
}

// 对象存储监控（桶和实例通用）
type RosMonitorData struct {
	Code int `json:"code"`
	Data struct {
		Columns []string `json:"Columns"`
		Values  [][]any  `json:"Values"`
	} `json:"data"`
}

// 对象存储实例列表
type RosInstanceList struct {
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
			Plan                    struct {
				ID                 int         `json:"id"`
				Region             string      `json:"region"`
				Subtype            string      `json:"subtype"`
				PlanName           string      `json:"plan_name"`
				Machine            string      `json:"machine"`
				ChargeType         string      `json:"charge_type"`
				Chinese            string      `json:"chinese"`
				IsFree             bool        `json:"is_free"`
				PointRenewPrice    interface{} `json:"point_renew_price"`
				IsSelling          bool        `json:"is_selling"`
				Price              int         `json:"price"`
				StorageSize        int         `json:"storage_size"`
				Bandwidth          int         `json:"bandwidth"`
				ExtraTransferPrice float64     `json:"extra_transfer_price"`
				ExtraStoragePrice  float64     `json:"extra_storage_price"`
			} `json:"Plan"`
			AccessKey               string      `json:"access_key"`       // ⚠️空
			SecretKey               string      `json:"secret_key"`       // ⚠️空
			IsPublicAccess          bool        `json:"is_public_access"` // 是否开启公共访问
			IsEnableExtraAccounting bool        `json:"is_enable_extra_accounting"`
			LastResetDate           int         `json:"last_reset_date"`
			NextResetDate           int         `json:"next_reset_date"`
			ExpDate                 int         `json:"ExpDate"`
			ExpireNotice            int         `json:"ExpireNotice"`
			AutoRenew               bool        `json:"AutoRenew"`
			UnsubscribeAble         bool        `json:"UnsubscribeAble"`
			Try                     bool        `json:"Try"`
			PublicAPIURL            string      `json:"public_api_url"`
			Buckets                 interface{} `json:"buckets"`
			ExtraPayTime            int         `json:"extra_pay_time"`
			UsageData               struct {
				MetricsUsageTotalBytes int `json:"metrics_usage_total_bytes"` // 存储使用量
				MetricsReceivedBytes   int `json:"metrics_received_bytes"`    // 接收字节数
				MetricsSentBytes       int `json:"metrics_sent_bytes"`        // 发送字节数
				MetricsRequests        int `json:"metrics_requests"`          // 请求数
				SpeedReceivedBytes     int `json:"speed_received_bytes"`      // 接收字节速度
				SpeedSentBytes         int `json:"speed_sent_bytes"`          // 发送字节速度
				SpeedRequests          int `json:"speed_requests"`            // 请求速度
				UpdateTime             int `json:"UpdateTime"`                // 更新时间
			} `json:"UsageData"`
			Restrictions           string `json:"restrictions"`
			PublicAccessRestricted bool   `json:"public_access_restricted"`
		} `json:"Records"`
	} `json:"data"`
}

// 修改存储桶Proxy设置请求
// 一个请求只能设置一项设置，要设置的设置项需要放在 Setting 中，只需要写需要修改的设置项，其他不填
type ModifyRosBucketProxySettingsRequest struct {
	DomainList   []string `json:"domain_list,omitempty"` // 域名列表
	GzipSettings struct {
		CompressLevel int      `json:"compress_level,omitempty"` // 压缩等级（1-5,默认为2）
		FileTypes     []string `json:"file_types,omitempty"`     // 文件MIME类型，默认["*"]
		IsEnable      bool     `json:"is_enable,omitempty"`      // 是否启用GZIP压缩，默认启用
	} `json:"gzip_settings,omitempty"` // Gzip设置
	IPRestrictSettings struct {
		Blacklist []string `json:"blacklist,omitempty"` // 黑名单IP或CIDR范围列表，和白名单列表只能二选一不能共存
		IsEnable  bool     `json:"is_enable,omitempty"` // 是否启用Referer防盗链
		Whitelist []string `json:"whitelist,omitempty"` // 白名单IP或CIDR范围列表，和黑名单列表只能二选一不能共存
	} `json:"ip_restrict_settings,omitempty"`
	RefererRestrictSettings struct {
		Blacklist     []string `json:"blacklist,omitempty"`      // 黑名单域名列表，和白名单列表只能二选一不能共存
		BypassMissing bool     `json:"bypass_missing,omitempty"` // 是否允许Referer请求头不存在或者格式有误
		IsEnable      bool     `json:"is_enable,omitempty"`      // 是否启用Referer防盗链
		Whitelist     []string `json:"whitelist,omitempty"`      // 白名单域名列表，和黑名单列表只能二选一不能共存
	} `json:"referer_restrict_settings,omitempty"`
	Setting     string `json:"setting"` // 必选项，要设置的选项(domain/ssl/waf/gzip/ip_restrict/referer_restrict)
	SslSettings struct {
		IsEnable bool `json:"is_enable,omitempty"` // 是否启用SSL
		IsForce  bool `json:"is_force,omitempty"`  // 是否强制使用SSL
	} `json:"ssl_settings,omitempty"` // SSL设置，SSL证书会自动匹配
	WafSettings struct {
		BlockTime     int  `json:"block_time,omitempty"`      // 未通过JS验证时的封禁时间，秒
		GlobalJsCheck bool `json:"global_js_check,omitempty"` // 是否启用全局JS检查，只有超过阈值才会生效
		GlobalQPS     int  `json:"global_qps,omitempty"`      // 全局访问速率限制，超过会被拒绝
		IsEnable      bool `json:"is_enable,omitempty"`       // 是否启用WAF防火墙
		PerIPJsCheck  bool `json:"per_ip_js_check,omitempty"` // 是否启用单IP JS检查，只有超过阈值才会生效
		PerIPQPS      int  `json:"per_ip_qps,omitempty"`      // 个IP访问速率限制，超过会被拒绝
	} `json:"waf_settings,omitempty"`
}

// ROS实例续费请求
type RenewRosInstanceRequest struct {
	Duration     int `json:"duration"`       // 续费时长（月）
	WithCouponID int `json:"with_coupon_id"` // 优惠券ID
}

// ROS实例自动续费选项
type RosInstanceAutoRenewOption struct {
	AutoRenewOption bool `json:"auto_renew_option"`
}

// ROS实例缩放请求
type ScaleRosInstanceRequest struct {
	DestPlan     int `json:"dest_plan"`      // 目标套餐ID
	WithCouponID int `json:"with_coupon_id"` // 优惠券ID
}

// 设置对象存储实例标签请求
type SetRosInstnceTagRequest struct {
	TagName string `json:"tag_name"`
}

// 开关对象存储实例的弹性计费选项请求
type ToggleRosInstanceExtraAccountingRequest struct {
	IsEnable bool `json:"is_enable"`
}

// 创建对象存储实例请求
type CreateRosInstanceRequest struct {
	Duration     int `json:"duration"`       // 购买时长（月）
	PlanID       int `json:"plan_id"`        // 套餐ID
	WithCouponID int `json:"with_coupon_id"` // 优惠券ID
}

// 创建对象存储实例响应
type CreateRosInstanceResponse struct {
	Code int `json:"code"`
	Data struct {
		ID         int    `json:"ID"`
		UID        int    `json:"UID"`
		PlanID     int    `json:"PlanID"`
		CreateDate int    `json:"CreateDate"` // 创建时间
		NodeUUID   string `json:"NodeUUID"`   // 节点UUID
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
		} `json:"Node"` // 节点信息
		Status                  string `json:"Status"`
		StopReason              string `json:"StopReason"`
		RewardPointsToBeCollect int    `json:"RewardPointsToBeCollect"` // 待领取积分奖励
		Tag                     string `json:"Tag"`
		Plan                    any    `json:"Plan"`
		AccessKey               string `json:"access_key"`       // AK
		SecretKey               string `json:"secret_key"`       // SK
		IsPublicAccess          bool   `json:"is_public_access"` // 是否允许公共访问
		IsEnableExtraAccounting bool   `json:"is_enable_extra_accounting"`
		LastResetDate           int    `json:"last_reset_date"`
		NextResetDate           int    `json:"next_reset_date"`
		ExpDate                 int    `json:"ExpDate"`
		ExpireNotice            int    `json:"ExpireNotice"`
		AutoRenew               bool   `json:"AutoRenew"` // 自动续费
		UnsubscribeAble         bool   `json:"UnsubscribeAble"`
		Try                     bool   `json:"Try"`
		PublicAPIURL            string `json:"public_api_url"`
		Buckets                 any    `json:"buckets"`
		ExtraPayTime            int    `json:"extra_pay_time"`
		Restrictions            string `json:"restrictions"`
		PublicAccessRestricted  bool   `json:"public_access_restricted"`
	} `json:"data"`
}

// 对象存储实例详情
type RosInstanceDetail struct {
	Code int `json:"code"`
	Data struct {
		Data struct {
			ID         int    `json:"ID"`
			UID        int    `json:"UID"`
			PlanID     int    `json:"PlanID"`
			CreateDate int    `json:"CreateDate"` // 创建时间
			NodeUUID   string `json:"NodeUUID"`   // 节点UUID
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
			} `json:"Node"` // 节点信息
			Status                  string `json:"Status"`
			StopReason              string `json:"StopReason"`
			RewardPointsToBeCollect int    `json:"RewardPointsToBeCollect"` // 待领取积分奖励
			Tag                     string `json:"Tag"`
			Plan                    struct {
				ID                 int     `json:"id"`
				Region             string  `json:"region"`
				Subtype            string  `json:"subtype"`
				PlanName           string  `json:"plan_name"`
				Machine            string  `json:"machine"`
				ChargeType         string  `json:"charge_type"`
				Chinese            string  `json:"chinese"`
				IsFree             bool    `json:"is_free"`
				PointRenewPrice    any     `json:"point_renew_price"`
				IsSelling          bool    `json:"is_selling"`
				Price              int     `json:"price"`
				StorageSize        int     `json:"storage_size"`
				Bandwidth          int     `json:"bandwidth"`
				ExtraTransferPrice float64 `json:"extra_transfer_price"`
				ExtraStoragePrice  float64 `json:"extra_storage_price"`
			} `json:"Plan"`
			AccessKey               string `json:"access_key"`                 // AK
			SecretKey               string `json:"secret_key"`                 // SK
			IsPublicAccess          bool   `json:"is_public_access"`           // 是否允许公共访问
			IsEnableExtraAccounting bool   `json:"is_enable_extra_accounting"` // 是否启用额外计费
			LastResetDate           int    `json:"last_reset_date"`            // 最后重置时间
			NextResetDate           int    `json:"next_reset_date"`            // 下次重置时间
			ExpDate                 int    `json:"ExpDate"`                    // 到期时间
			ExpireNotice            int    `json:"ExpireNotice"`
			AutoRenew               bool   `json:"AutoRenew"` // 是否自动续费
			UnsubscribeAble         bool   `json:"UnsubscribeAble"`
			Try                     bool   `json:"Try"` // 是否试用
			PublicAPIURL            string `json:"public_api_url"`
			Buckets                 any    `json:"buckets"`
			ExtraPayTime            int    `json:"extra_pay_time"`
			UsageData               struct {
				MetricsUsageTotalBytes int `json:"metrics_usage_total_bytes"`
				MetricsReceivedBytes   int `json:"metrics_received_bytes"`
				MetricsSentBytes       int `json:"metrics_sent_bytes"`
				MetricsRequests        int `json:"metrics_requests"`
				SpeedReceivedBytes     int `json:"speed_received_bytes"`
				SpeedSentBytes         int `json:"speed_sent_bytes"`
				SpeedRequests          int `json:"speed_requests"`
				UpdateTime             int `json:"UpdateTime"`
			} `json:"UsageData"`
			Restrictions           string `json:"restrictions"`
			PublicAccessRestricted bool   `json:"public_access_restricted"`
		} `json:"Data"`
		RenewPointPrice struct {
			Num7  int `json:"7"`  // 积分续费七天
			Num31 int `json:"31"` // 积分续费一个月
		} `json:"RenewPointPrice"` // 积分续费
	} `json:"data"`
}

// 对象存储存储桶列表
//
// ⚠️注意这个接口是拿不到AK和SK的，响应里面的AK和SK都是空的
type RosBucketList struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords int `json:"TotalRecords"`
		Records      []struct {
			ID             int    `json:"id"`
			Name           string `json:"name"`
			UID            int    `json:"uid"`
			CreateDate     int    `json:"create_date"`
			AccessKey      string `json:"access_key"`
			SecretKey      string `json:"secret_key"`
			IsPublicAccess bool   `json:"is_public_access"`
			InstanceID     int    `json:"instance_id"`
			Instance       struct {
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
				Status                  string `json:"Status"`
				StopReason              string `json:"StopReason"`
				RewardPointsToBeCollect int    `json:"RewardPointsToBeCollect"`
				Tag                     string `json:"Tag"`
				Plan                    struct {
					ID                 int     `json:"id"`
					Region             string  `json:"region"`
					Subtype            string  `json:"subtype"`
					PlanName           string  `json:"plan_name"`
					Machine            string  `json:"machine"`
					ChargeType         string  `json:"charge_type"`
					Chinese            string  `json:"chinese"`
					IsFree             bool    `json:"is_free"`
					PointRenewPrice    any     `json:"point_renew_price"`
					IsSelling          bool    `json:"is_selling"`
					Price              int     `json:"price"`
					StorageSize        int     `json:"storage_size"`
					Bandwidth          int     `json:"bandwidth"`
					ExtraTransferPrice float64 `json:"extra_transfer_price"`
					ExtraStoragePrice  float64 `json:"extra_storage_price"`
				} `json:"Plan"`
				AccessKey               string `json:"access_key"` // ⚠️空
				SecretKey               string `json:"secret_key"` // ⚠️空
				IsPublicAccess          bool   `json:"is_public_access"`
				IsEnableExtraAccounting bool   `json:"is_enable_extra_accounting"`
				LastResetDate           int    `json:"last_reset_date"`
				NextResetDate           int    `json:"next_reset_date"`
				ExpDate                 int    `json:"ExpDate"`
				ExpireNotice            int    `json:"ExpireNotice"`
				AutoRenew               bool   `json:"AutoRenew"`
				UnsubscribeAble         bool   `json:"UnsubscribeAble"`
				Try                     bool   `json:"Try"`
				PublicAPIURL            string `json:"public_api_url"`
				Buckets                 any    `json:"buckets"`
				ExtraPayTime            int    `json:"extra_pay_time"`
				Restrictions            string `json:"restrictions"`
				PublicAccessRestricted  bool   `json:"public_access_restricted"`
			} `json:"Instance"`
			StopReason                   string `json:"stop_reason"`
			DomainList                   any    `json:"DomainList"`
			SslIsEnable                  bool   `json:"ssl_is_enable"`
			SSLCertList                  any    `json:"SSLCertList"`
			SslAutoRedirect              bool   `json:"ssl_auto_redirect"`
			WafIsEnable                  bool   `json:"waf_is_enable"`
			WafGlobalQPS                 int    `json:"waf_global_qps"`
			WafPerIPQPS                  int    `json:"waf_per_ip_qps"`
			WafBlockTime                 int    `json:"waf_block_time"`
			WafGlobalJsCheck             bool   `json:"waf_global_js_check"`
			WafPerIPJsCheck              bool   `json:"waf_per_ip_js_check"`
			RefererRestrictIsEnable      bool   `json:"referer_restrict_is_enable"`
			RefererRestrictWhitelist     any    `json:"RefererRestrictWhitelist"`
			RefererRestrictBlacklist     any    `json:"RefererRestrictBlacklist"`
			RefererRestrictBypassMissing bool   `json:"referer_restrict_bypass_missing"`
			IPRestrictIsEnable           bool   `json:"ip_restrict_is_enable"`
			IPRestrictWhitelist          any    `json:"IPRestrictWhitelist"`
			IPRestrictBlacklist          any    `json:"IPRestrictBlacklist"`
			GzipIsEnable                 bool   `json:"gzip_is_enable"`
			GzipCompressLevel            int    `json:"gzip_compress_level"`
			UsageData                    struct {
				MetricsUsageTotalBytes int `json:"metrics_usage_total_bytes"`
				MetricsReceivedBytes   int `json:"metrics_received_bytes"`
				MetricsSentBytes       int `json:"metrics_sent_bytes"`
				MetricsRequests        int `json:"metrics_requests"`
				SpeedReceivedBytes     int `json:"speed_received_bytes"`
				SpeedSentBytes         int `json:"speed_sent_bytes"`
				SpeedRequests          int `json:"speed_requests"`
				UpdateTime             int `json:"UpdateTime"`
			} `json:"UsageData"`
			ReadyDelete bool `json:"ready_delete"`
		} `json:"Records"`
	} `json:"data"`
}
