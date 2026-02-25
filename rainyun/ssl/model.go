package ssl

// SSL证书
type SslCertificate struct {
	Cert string `json:"cert"` // 证书
	Key  string `json:"key"`  // 私钥
}

// SSL证书列表
type SslCertificateList struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords int `json:"TotalRecords"`
		Records      []struct {
			ID            int         `json:"ID"`
			UID           int         `json:"UID"`
			Domain        string      `json:"Domain"`        // 域名(逗号分割)
			Issuer        string      `json:"Issuer"`        // 品牌
			StartDate     int         `json:"StartDate"`     // 开始时间
			ExpDate       int         `json:"ExpDate"`       // 结束时间
			UploadTime    int         `json:"UploadTime"`    // 上传时间
			NginxErr      string      `json:"NginxErr"`      // ？
			BaishanCertID int         `json:"BaishanCertID"` // 白山云证书ID
			BindDomains   interface{} `json:"BindDomains"`   // 绑定的域名
		} `json:"Records"`
	} `json:"data"`
}

// SSL证书详情
type SslDetail struct {
	Code int `json:"code"`
	Data struct {
		Cert       string `json:"Cert"`       // 证书
		Key        string `json:"Key"`        // 私钥
		DomainName string `json:"DomainName"` // 域名(逗号分割)
		Issuer     string `json:"Issuer"`     // 品牌
		StartDate  int    `json:"StartDate"`  // 开始时间
		ExpDate    int    `json:"ExpDate"`    // 结束时间
		RemainDays int    `json:"RemainDays"` // 剩余天数
	} `json:"data"`
}
