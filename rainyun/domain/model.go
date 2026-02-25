package domain

// 域名过白请求
type AddDomainToWhiteListRequest struct {
	Domain string `json:"domain"` // 域名
	Region string `json:"region"` // 区域：cn-sq1/cn-nb1/cn-xy1/cn-cq1
}

// 域名白名单列表
type DomainWhitelist struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords int `json:"TotalRecords"`
		Records      []struct {
			ID          int    `json:"id"` // ID
			UID         int    `json:"uid"`
			Domain      string `json:"domain"`       // 域名
			SiteLicense string `json:"site_license"` // 备案号（貌似只有早期过白的会有这个）
			Region      string `json:"region"`       // 区域
			AddTime     int    `json:"add_time"`     // 时间
			Status      int    `json:"status"`       //	状态: 1:已处理
		} `json:"Records"`
	} `json:"data"`
}

// 已验证域名列表
type VerifiedDomainList struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords int `json:"TotalRecords"`
		Records      []struct {
			ID      int    `json:"id"` // id
			UID     int    `json:"uid"`
			Domain  string `json:"domain"`   // 域名
			AddTime int    `json:"add_time"` // 时间
		} `json:"Records"`
	} `json:"data"`
}

// 域名验证信息
type DomainVerificationInfo struct {
	Code int `json:"code"`
	Data struct {
		Rr        string `json:"rr"`         // 主机名
		TopDomain string `json:"top_domain"` // 主域名
		Record    string `json:"record"`     // 记录值
	} `json:"data"`
}

// 添加域名认证请求
type AddDomainVerificationRequest struct {
	Domain string `json:"domain"`
}
