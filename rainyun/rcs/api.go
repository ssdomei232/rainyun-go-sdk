package rcs

import (
	"fmt"
	"net/url"

	"github.com/ssdomei232/rainyun-go-sdk/v2/rainyun/common"
)

// 设置RCS IP描述
//
// id: RCS ID
// ip: ip address
//
// desc: description
func (c *Client) SetRcsEipDescription(id int, ip string, desc string) (*common.BasicOperationResponse, error) {
	path := "/product/rcs/{id}/eip/description"

	req := SetRcsEipDescriptionRequest{
		Description: desc,
		IP:          ip,
	}

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 获取 RCS 列表
//
// options: RCS查询参数 可以用 EncodingStandardQueryParameters 获取.
func (c *Client) GetRcsList(options string) (*RcsListResponse, error) {
	path := fmt.Sprintf("/product/rcs?options=%s", options)

	var resp RcsListResponse
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 创建 RCS
func (c *Client) CreateRcs(req *CreateRcsRequest) (*CreateRcsResopnse, error) {
	path := "/product/rcs/"

	var resp CreateRcsResopnse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 获取RCS详情
//
// id: RCS ID
func (c *Client) GetRcsDetails(id int) (*RcsDetails, error) {
	path := fmt.Sprintf("/product/rcs/%d/", id)

	var resp RcsDetails
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// RCS创建备份
//
// id: RCS ID, label: 备份名称
func (c *Client) CreateRcsBackup(id int, label string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/", id)

	req := CreateRcsBackupRequest{
		Label: label,
	}

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, resp)

	return &resp, err
}

// RCS删除备份
//
// id: RCS ID, bid: 备份ID
func (c *Client) DeleteRcsBackup(id int, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/%d/", id, bid)

	var resp common.BasicOperationResponse
	err := c.DoRequest("DELETE", path, nil, &resp)

	return &resp, err
}

// RCS取消备份
//
// id: RCS ID, bid: 备份ID
func (c *Client) CancelRcsBackup(id int, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/%d/cancel", id, bid)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// RCS还原备份
//
// id: RCS ID, bid: 备份ID
func (c *Client) RestoreRcsBackup(id int, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/%d/restore", id, bid)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// RCS设置备份选项
//
// id: RCS ID
func (c *Client) EnableRcsAutoBackup(id int, req *RcsSetBackupOptionsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/setting", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, req, &resp)

	return &resp, err
}

// RCS重装系统
//
// id: RCS ID
func (c *Client) ReinstallRcs(id int, req *ReinstallRcsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/changeos", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// RCS管理弹性云盘
//
// id: RCS ID
func (c *Client) RcsManagesElasticCloudDisks(id int, req *RcsManagesElasticCloudDisksRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/edisk/", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 创建并绑定弹性IP到RCS
//
// id: RCS ID
func (c *Client) CreateAndBindElasticIpToRcs(id int, req *CreateAndBindIpToRcsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/eip/", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 更换RCS IP
//
// id: RCS ID
func (c *Client) ChangeRcsIP(id int, req *ChangeRcsIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/eip/change", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 放弃RCS IP
//
// id: RCS ID
func (c *Client) DisCardRcsIP(id int, req DisCardRcsIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/eip/discard", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 获取RCS防火墙规则列表
//
// id: RCS ID
//
// options: RCS查询参数 可以用 EncodingStandardQueryParameters 获取.
func (c *Client) GetRcsFirewallRules(id int, options string) (*RcsFirewallRuleList, error) {
	path := fmt.Sprintf("/product/rcs/%d/firewall/rule?options=%s", id, options)

	var resp RcsFirewallRuleList
	err := c.DoRequest("GET", path, options, &resp)

	return &resp, err
}

// 创建/设置RCS防火墙规则
//
// id: RCS ID
func (c *Client) SetRcsFirewallRule(id int, req *SetRcsFirewallRuleRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/firewall/rule", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 删除RCS防火墙规则
//
// id: RCS ID
func (c *Client) DeleteRcsFirewallRule(id int, ruleID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/firewall/rule/%d", id, ruleID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("DELETE", path, nil, &resp)

	return &resp, err
}

// 移动RCS防火墙规则优先级
//
// id: RCS ID
func (c *Client) MobileRcsFirewallRulePriority(id int, req MobileRcsFirewallRulePriorityRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%dfirewall/rule/{ruleId}/pos", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PUT", path, req, &resp)

	return &resp, err
}

// 释放RCS
//
// id: RCS ID
func (c *Client) FreeRcs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/free", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 获取RCS监控数据
//
// id: RCS ID
//
// startDate: 开始时间
//
// endDate: 结束时间
func (c *Client) GetRcsMonitorData(id int, startDate int, endDate int) (*RcsMonitoringData, error) {
	path := fmt.Sprintf("/product/rcs/%d/monitor?start_date=%d&end_date=%d", id, startDate, endDate)

	var resp RcsMonitoringData
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 添加RCS NAT端口映射
//
// id: RCS ID
func (c *Client) AddRcsNatPortMapping(id int, req *AddRcsNatPortMappingRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/nat", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 删除RCS NAT端口映射
//
// id: RCS ID
//
// natID: NAT规则 ID
func (c *Client) DeleteRcsNatPortMapping(id int, natID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/nat?nat_id=%d", id, natID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("DELETE", path, nil, &resp)

	return &resp, err
}

// 云服务器重启操作
//
// id: RCS ID
func (c *Client) RebootRcs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/reboot", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 获取 RCS 续费价格
//
// id: RCS ID, duration: 续费时长(月), coupon: 优惠券ID
func (c *Client) GetRcsRenewPrice(id int, duration int, couponID int) (*RCSRenewPrice, error) {
	path := fmt.Sprintf("/product/rcs/price?scene=renew&product_id=%d&duration=%d&with_coupon_id=%d&is_old=true", id, duration, couponID)

	var resp RCSRenewPrice
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// RCS续费
//
// id: RCS ID
func (c *Client) RenewRcs(id int, req RenewRcsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/renew/", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// RCS自动续费选项
//
// id: RCS ID
func (c *Client) EnableRcsAutoRenew(id int, req EnableRcsAutoRenewRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/renew/option", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// RCS重置密码
//
// id: RCS ID
//
// newPass: 新密码,留空则自动生成
func (c *Client) ResetRcsPassword(id int, newPass string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/reset-password", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, ResetRcsPasswordRequest{Password: newPass}, &resp)

	return &resp, err
}

// RCS开机
//
// id: RCS ID
func (c *Client) StartRcs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/start", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// RCS关机
//
// id: RCS ID
func (c *Client) StopRcs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/stop", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 设置RCS标签
//
// id: RCS ID
//
// tag: 标签
func (c *Client) SetRcsTag(id int, tag string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/tag", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, SetRcsTagRequest{TagName: tag}, &resp)

	return &resp, err
}

// RCS充流量
//
// id: RCS ID
//
// count: 充多少(单位G)
func (c *Client) ChargeRcsTrafic(id int, count int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/traffic/charge", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// RCS限流
//
// id: RCS ID
//
// threshold: 日流量阈值(G)
//
// limit: 限制带宽(M)
func (c *Client) LimitRcsTrafic(id int, threshold int, limit int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/traffic/limit", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, LimitRcsTrafficRequest{DayTrafficInGb: threshold, TrafficLimit: limit}, &resp)

	return &resp, err
}

// RCS升级
//
// id: RCS ID
//
// plan: 升级到的套餐ID
//
// coupon: 优惠券ID,默认为0
func (c *Client) UpgradeRcs(id int, plan int, coupon int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/upgrade", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, UpgradeRcsRequest{DestPlan: plan, WithCouponID: coupon}, &resp)

	return &resp, err
}

// RCS连接VNC
//
// id: RCS ID
//
// consoleType: 控制台类型,可选值: novnc,xtermjs
func (c *Client) GetRcsVnc(id int, consoleType string) (*common.VncConnectionInfo, error) {
	path := fmt.Sprintf("/product/rcs/%d/vnc?console_type=%s", id, consoleType)

	var resp common.VncConnectionInfo
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取RCS对应的PVE地址
//
// id: RCS ID
func (c *Client) GetRcsPveAddress(id int) (pveAddress string, err error) {
	v, err := c.GetRcsVnc(id, "novnc")
	if err != nil {
		return "", err
	}

	parsedURL, err := url.Parse(v.Data.RequestURL)
	if err != nil {
		return "", err
	}

	return parsedURL.Hostname(), err
}
