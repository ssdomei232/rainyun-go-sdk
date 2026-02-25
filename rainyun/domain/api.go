package domain

import (
	"fmt"

	"github.com/ssdomei232/rainyun-go-sdk/rainyun/common"
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
