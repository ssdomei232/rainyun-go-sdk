package rainyun_go_sdk

import "fmt"

/* =============用户部分============= */

// 获取用户信息.
func (c *Client) GetUserInfo() (*UserInfo, error) {
	path := "/user/"

	var resp UserInfo
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取可兑换积分产品列表.
func (c *Client) GetUserRewardPruducts() (*UserRewardProducts, error) {
	path := "/user/reward/products"

	var resp UserRewardProducts
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 发布优惠券给下级用户
func (c *Client) PublishCouponsToLowerLevelUsers(req *PublishCouponsToLowerLevelUsersRequest) (*BasicOperationResponse, error) {
	path := "/user/vip/coupon"

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 发布优惠券到积分商城(供下级领取)
func (c *Client) PostCouponsToPointsMall(req *PostCouponsToPointsMallRequest) (*BasicOperationResponse, error) {
	path := "/user/vip/coupon"

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 获取积分商城商品列表
func (c *Client) GetPointsMallItems() (*PointsMallItemsResponse, error) {
	path := "/user/reward/items"

	var resp PointsMallItemsResponse
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 兑换积分物品
func (c *Client) RedeemPointsForItem(id int) (*BasicOperationResponse, error) {
	path := "/user/reward/items"

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, RedeemPointsForItemRequest{ItemID: id}, &resp)

	return &resp, err
}

// 请求二次验证
// verificationMethod: 验证方式(sms/totp/email)
func (c *Client) Request2FA(verificationMethod string) (*BasicOperationResponse, error) {
	path := "/user/mfa/request"

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, Request2FARequest{Type: verificationMethod}, &resp)

	return &resp, err

}

// 验证二次验证结果
// authCode: 验证码
func (c *Client) Verify2FAResult(authCode int) (*BasicOperationResponse, error) {
	path := "/user/mfa/verify"

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, Verify2FAResultRequest{AuthCode: authCode}, &resp)

	return &resp, err
}

/* ==============RCS============== */

// 设置RCS IP描述
// id: RCS ID, ip: ip address, desc: description
func (c *Client) SetRcsEipDescription(id int, ip string, desc string) (*BasicOperationResponse, error) {
	path := "/product/rcs/{id}/eip/description"

	req := SetRcsEipDescriptionRequest{
		Description: desc,
		IP:          ip,
	}

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 获取 RCS 列表
// options: RCS查询参数 可以用 MarshalRCSQueryParameters 获取.
func (c *Client) GetRcsList(options string) (*RcsListResponse, error) {
	path := "/product/rcs"

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
func (c *Client) GetRcsDetails(id int) (*RcsDetails, error) {
	path := fmt.Sprintf("/product/rcs/%d/", id)

	var resp RcsDetails
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// RCS创建备份
// id: RCS ID, label: 备份名称
func (c *Client) CreateRcsBackup(id int, label string) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/", id)

	req := CreateRcsBackupRequest{
		Label: label,
	}

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, req, resp)

	return &resp, err
}

// RCS删除备份
// id: RCS ID, bid: 备份ID
func (c *Client) DeleteRcsBackup(id int, bid int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/%d/", id, bid)

	var resp BasicOperationResponse
	err := c.DoRequest("DELETE", path, nil, &resp)

	return &resp, err
}

// RCS取消备份
// id: RCS ID, bid: 备份ID
func (c *Client) CancelRcsBackup(id int, bid int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/%d/cancel", id, bid)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// RCS还原备份
// id: RCS ID, bid: 备份ID
func (c *Client) RestoreRcsBackup(id int, bid int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/%d/restore", id, bid)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// RCS开启每日自动备份
// id: RCS ID
func (c *Client) EnableRcsAutoBackup(id int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/setting", id)

	var resp BasicOperationResponse
	err := c.DoRequest("PATCH", path, nil, &resp)

	return &resp, err
}

// RCS重装系统
// id: RCS ID
func (c *Client) ReinstallRcs(id int, req *ReinstallRcsRequest) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/changeos", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// RCS管理弹性云盘
// id: RCS ID
func (c *Client) RcsManagesElasticCloudDisks(id int, req *RcsManagesElasticCloudDisksRequest) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/edisk/", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 创建并绑定弹性IP到RCS
// id: RCS ID
func (c *Client) CreateAndBindElasticIpToRcs(id int, req *CreateAndBindIpToRcsRequest) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/eip/", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 更换IP
// id: RCS ID
func (c *Client) ChangeIP(id int, req *ChangeRcsIPRequest) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/eip/change", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 放弃IP
// id: RCS ID
func (c *Client) DisCardIP(id int, req DisCardRcsIPRequest) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/eip/discard", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 获取RCS防火墙规则列表
// id: RCS ID
// options: RCS查询参数 可以用 MarshalRCSQueryParameters 获取.
func (c *Client) GetFirewallRules(id int, options string) (*RcsFirewallRuleList, error) {
	path := fmt.Sprintf("/product/rcs/%d/firewall/rule", id)

	var resp RcsFirewallRuleList
	err := c.DoRequest("GET", path, options, &resp)

	return &resp, err
}

// 创建/设置RCS防火墙规则
// id: RCS ID
func (c *Client) SetFirewallRule(id int, req *SetRcsFirewallRuleRequest) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/firewall/rule", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 删除RCS防火墙规则
// id: RCS ID
func (c *Client) DeleteFirewallRule(id int, ruleID int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/firewall/rule/%d", id, ruleID)

	var resp BasicOperationResponse
	err := c.DoRequest("DELETE", path, nil, &resp)

	return &resp, err
}

// 移动RCS防火墙规则优先级
// id: RCS ID
func (c *Client) MobileFirewallRulePriority(id int, req MobileRcsFirewallRulePriorityRequest) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%dfirewall/rule/{ruleId}/pos", id)

	var resp BasicOperationResponse
	err := c.DoRequest("PUT", path, req, &resp)

	return &resp, err
}

// 释放RCS
// id: RCS ID
func (c *Client) FreeRcs(id int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/free", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 获取RCS监控数据
// id: RCS ID
// startDate: 开始时间
// endDate: 结束时间
func (c *Client) GetRcsMonitorData(id int, startDate int, endDate int) (*RcsMonitoringData, error) {
	path := fmt.Sprintf("/product/rcs/%d/monitor?start_date=%d&end_date=%d", id, startDate, endDate)

	var resp RcsMonitoringData
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 添加NAT端口映射
// id: RCS ID
func (c *Client) AddRcsNatPortMapping(id int, req *AddRcsNatPortMappingRequest) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/nat", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 删除NAT端口映射
// id: RCS ID
// natID: NAT规则 ID
func (c *Client) DeleteRcsNatPortMapping(id int, natID int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/nat?nat_id=%d", id, natID)

	var resp BasicOperationResponse
	err := c.DoRequest("DELETE", path, nil, &resp)

	return &resp, err
}

// 云服务器重启操作
// id: RCS ID
func (c *Client) RebootRcs(id int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/reboot", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 获取续费价格
// id: RCS ID
func (c *Client) GetRenewPrice(id int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/renew/", id)

	var resp BasicOperationResponse
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// RCS续费
// id: RCS ID
func (c *Client) RenewRcs(id int, req RenewRcsRequest) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/renew/", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// RCS自动续费选项
// id: RCS ID
func (c *Client) EnableRcsAutoRenew(id int, req EnableRcsAutoRenewRequest) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/renew/option", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// RCS重置密码
// id: RCS ID
// newPass: 新密码,留空则自动生成
func (c *Client) ResetRcsPassword(id int, newPass string) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/reset-password", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, ResetRcsPasswordRequest{Password: newPass}, &resp)

	return &resp, err
}

// RCS开机
// id: RCS ID
func (c *Client) StartRcs(id int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/start", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// RCS关机
// id: RCS ID
func (c *Client) StopRcs(id int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/stop", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 设置RCS标签
// id: RCS ID
// tag: 标签
func (c *Client) SetRcsTag(id int, tag string) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/tag", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, SetRcsTagRequest{TagName: tag}, &resp)

	return &resp, err
}

// RCS充流量
// id: RCS ID
// count: 充多少(单位G)
func (c *Client) ChargeRcsTrafic(id int, count int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/traffic/charge", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// RCS限流
// id: RCS ID
// threshold: 日流量阈值(G)
// limit: 限制带宽(M)
func (c *Client) LimitRcsTrafic(id int, threshold int, limit int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/traffic/limit", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, LimitRcsTrafficRequest{DayTrafficInGb: threshold, TrafficLimit: limit}, &resp)

	return &resp, err
}

// RCS升级
// id: RCS ID
// plan: 升级到的套餐ID
// coupon: 优惠券ID,默认为0
func (c *Client) UpgradeRcs(id int, plan int, coupon int) (*BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/upgrade", id)

	var resp BasicOperationResponse
	err := c.DoRequest("POST", path, UpgradeRcsRequest{DestPlan: plan, WithCouponID: coupon}, &resp)

	return &resp, err
}

// 获取RCS使用情况列表
func (c *Client) GetRcsUsageList() (*RcsUsageList, error) {
	path := "/product/rcs/usage"

	var resp RcsUsageList
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}
