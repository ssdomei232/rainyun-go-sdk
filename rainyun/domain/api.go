package domain

import (
	"fmt"

	"github.com/ssdomei232/rainyun-go-sdk/v2/rainyun/common"
)

// 备案域名过白
//
// domain： 域名
//
// region： 区域：cn-sq1/cn-nb1/cn-xy1/cn-cq1
func (c *Client) AddDomainToWhiteList(domain string, region string) (*common.BasicOperationResponse, error) {
	path := "/product/domain_white_list"

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, &AddDomainToWhiteListRequest{
		Domain: domain,
		Region: region,
	}, &resp)

	return &resp, err
}

// 获取域名白名单列表
//
// options: 查询参数 可以用 EncodingStandardQueryParameters 获取.
func (c *Client) GetDomainWhiteList(options string) (*DomainWhitelist, error) {
	path := fmt.Sprintf("/product/domain/whitelist?options=%s", options)

	var resp DomainWhitelist
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取已验证域名列表
//
// options: 查询参数 可以用 EncodingStandardQueryParameters 获取.
func (c *Client) GetVerifiedDomainList(options string) (*VerifiedDomainList, error) {
	path := fmt.Sprintf("/product/domain/certify?options=%s", options)

	var resp VerifiedDomainList
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 添加域名认证
//
// domain: 域名
func (c *Client) AddDomainCertify(domain string) (*DomainVerificationInfo, error) {
	path := "/product/domain/certify"

	var resp DomainVerificationInfo
	err := c.DoRequest("POST", path, AddDomainVerificationRequest{Domain: domain}, &resp)

	return &resp, err
}

// 域名认证校验
//
// domain: 域名
func (c *Client) VerifyDomainCertify(domain string) (*common.BasicOperationResponse, error) {
	path := "/product/domain/certify/verify"

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, AddDomainVerificationRequest{Domain: domain}, &resp)

	return &resp, err
}

// 添加域名DNS解析
//
// id: 域名ID
func (c *Client) AddDomainDNSRecord(id int, req *AddDomainDNSRecordRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/dns", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 修改域名DNS解析
func (c *Client) UpdateDomainDNSRecord(id int, req *AddDomainDNSRecordRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/dns", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, req, &resp)

	return &resp, err
}

// 删除域名DNS解析
func (c *Client) DeleteDomainDNSRecord(id int, recordID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/dns", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("DELETE", path, DeleteDomainDNSRecordRequest{RecordID: recordID}, &resp)

	return &resp, err
}

// 添加域名DNSSEC
func (c *Client) AddDomainDNSSEC(id int, req *AddDomainDNSSECRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/dnssec", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 删除域名DNSSEC
func (c *Client) DeleteDomainDNSSEC(id int, req *DeleteDomainDNSSECRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/dnssec/delete", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 同步域名DNSSEC
func (c *Client) SyncDomainDNSSEC(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/dnssec/sync", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 关闭域名锁定
func (c *Client) UnlockDomain(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/lock/disable", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PUT", path, nil, &resp)

	return &resp, err
}

// 开启域名锁定
func (c *Client) LockDomain(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/lock/enable", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PUT", path, nil, &resp)

	return &resp, err
}

// 修改域名NS服务器
func (c *Client) UpdateDomainNS(id int, nss []string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/nameservers", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PUT", path, UpdateDomainNSRequest{NameServers: nss}, &resp)

	return &resp, err
}

// 重置域名NS服务器
func (c *Client) ResetDomainNS(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/nameservers/reset", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 续费域名
func (c *Client) RenewDomain(id int, years int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/renew", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, map[string]int{"years": years}, &resp)

	return &resp, err
}

// 域名过户
func (c *Client) TransferDomain(id int, req *DomainTransferRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/transfer", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 查询域名模板列表
//
// options: 查询参数 可以用 EncodingStandardQueryParameters 获取.
func (c *Client) GetDomainTemplateList(options string) (*DomainTemplateList, error) {
	path := fmt.Sprintf("/product/domain/template/?options=%s", options)

	var resp DomainTemplateList
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 删除域名信息模板
func (c *Client) DeleteDomainTemplate(sysID string) (*common.BasicOperationResponse, error) {
	path := "/product/domain/template/"

	var resp common.BasicOperationResponse
	err := c.DoRequest("DELETE", path, DeleteDomainTemplateRequest{SysID: sysID}, &resp)

	return &resp, err
}

// 更新域名管理密码
func (c *Client) UpdateDomainPassword(id int, password string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/password", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, UpdateDomainPasswordRequest{Password: password}, &resp)

	return &resp, err
}

// 编辑域名模板
func (c *Client) EditDomainTemplate(sysID string, req *EditDomainTemplateRequest) (*common.BasicOperationResponse, error) {
	path := "/product/domain/template"

	var resp common.BasicOperationResponse
	err := c.DoRequest("PUT", path, req, &resp)

	return &resp, err
}

// 获取域名模板详情
func (c *Client) GetDomainTemplateDetail(sysID string) (*GetDomainTemplateDetailResponse, error) {
	path := fmt.Sprintf("/product/domain/template/detail/?sys_id=%s", sysID)

	var resp GetDomainTemplateDetailResponse
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取域名whois信息
func (c *Client) GetDomainWhoisInfo(domain string) (*DomainWhoisInfo, error) {
	path := fmt.Sprintf("/product/domain/whois?domain=%s", domain)

	var resp DomainWhoisInfo
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 检查域名能否注册
func (c *Client) CheckDomainAvailability(req *CheckDomainAvailableRequest) (*CheckDomainAvailableResponse, error) {
	path := "/product/domain/check"

	var resp CheckDomainAvailableResponse
	err := c.DoRequest("GET", path, req, &resp)

	return &resp, err
}

// TODO: 列出域名列表,获取域名详情,下载域名证书，获取域名DNS解析记录列表，获取域名DNSSEC详情,获取域名续费价格,域名注册,获取域名管理密码
