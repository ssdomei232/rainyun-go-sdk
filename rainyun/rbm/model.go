package rbm

// RBM套餐列表
type RBMPlanList struct {
	Code int `json:"code"`
	Data []struct {
		ID              int    `json:"id"`                // 套餐ID
		Region          string `json:"region"`            // 地域
		Line            string `json:"line"`              // 线路 3c/single/bgp
		Subtype         string `json:"subtype"`           // rbm
		PlanName        string `json:"plan_name"`         // 套餐名称
		Machine         string `json:"machine"`           // 意义不明 generic/generic_2
		ChargeType      string `json:"charge_type"`       // 计费方式 package(流量不限)/package_traffic(流量叠加)
		Chinese         string `json:"chinese"`           // 中文名
		IsFree          bool   `json:"is_free"`           // 是否免费(为什么会有这种参数)
		PointRenewPrice any    `json:"point_renew_price"` // 积分续费价格，全是 null
		IsSelling       bool   `json:"is_selling"`        // 是否在售
		Price           int    `json:"price"`             // 价格
		TrafficBaseGb   int    `json:"traffic_base_gb"`   // 流量基数，单位GB
		TrafficPrice    any    `json:"traffic_price"`     // 流量价格，单位GB，package_traffic 计费方式才有值
		CPU             int    `json:"cpu"`               // cpu核数
		Memory          int    `json:"memory"`            // 内存，单位GB
		Storage         int    `json:"storage"`           // 存储，单位GB
		NetIn           int    `json:"net_in"`            // 入网带宽，单位Mbps
		NetOut          int    `json:"net_out"`           // 出网带宽，单位Mbps
		AvailableStock  int    `json:"available_stock"`   // 可用库存
		IPPrices        any    `json:"ip_prices"`         // IP价格，单位元/月
	} `json:"data"`
}

// RBM实例更换系统请求
type RBMChangeOSRequest struct {
	OsID int `json:"os_id"`
}

// 创建并绑定弹性IP到RBM请求
type RBMAssociateEIPRequest struct {
	WithFlags  string `json:"with_flags"`   // IP特征(可选)
	WithIPNum  int    `json:"with_ip_num"`  // 数量
	WithIPType string `json:"with_ip_type"` // IPv4/IPv6
}

// 更换IP请求
type RBMChangeIPRequest struct {
	DisableOldIPReason string `json:"disable_old_ip_reason"` // 禁用旧IP的原因(可选)
	IP                 string `json:"ip"`                    // 旧IP地址
	ToIP               string `json:"to_ip"`                 // 不明，可选
}

// 设置IP描述请求
type RBMSetIPDescriptionRequest struct {
	Description string `json:"description"`
	IP          string `json:"ip"`
}

// 放弃IP请求
type RBMReleaseIPRequest struct {
	IP string `json:"ip"`
}

// 充流量请求
type RBMChargeTrafficRequest struct {
	TrafficInGb int `json:"traffic_in_gb"`
}
