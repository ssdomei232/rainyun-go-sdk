package rgs

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ssdomei232/rainyun-go-sdk/rainyun/common"
	"github.com/ssdomei232/rainyun-go-sdk/rainyun/rcs"
)

// 获取游戏云列表
//
// options: 查询参数 可以用 EncodingStandardQueryParameters 获取
func (c *Client) GetRgsList(options string) (*RgsList, error) {
	path := "/product/rgs/"

	var resp RgsList
	err := c.DoRequest("GET", path, options, &resp)

	return &resp, err
}

// 创建游戏云
func (c *Client) CreateRgs(req *CreateRgsRequest) (*CreateRgsResponse, error) {
	path := "/product/rgs/"

	var resp CreateRgsResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 获取游戏云详情
//
// id： 游戏云 ID
func (c *Client) GetRgsDetails(id int) (*RgsDetail, error) {
	path := fmt.Sprintf("/product/rgs/%d/", id)

	var resp RgsDetail
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 游戏云创建备份
//
// id： 游戏云 ID
//
// label： 备份标签
func (c *Client) CreateRgsBackup(id int, label string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/backup/", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, rcs.CreateRcsBackupRequest{Label: label}, &resp)

	return &resp, err
}

// 游戏云删除备份
//
// id: 游戏云 ID, bid: 备份ID
func (c *Client) DeleteRgsBackup(id int, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/backup/%d/", id, bid)

	var resp common.BasicOperationResponse
	err := c.DoRequest("DELETE", path, nil, &resp)

	return &resp, err
}

// 游戏云取消备份
//
// id: RGS ID, bid: 备份ID
func (c *Client) CancelRgsBackup(id int, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/backup/%d/cancel", id, bid)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 游戏云还原备份
//
// id: 游戏云 ID, bid: 备份ID
func (c *Client) RestoreRgsBackup(id int, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/backup/%d/restore", id, bid)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 游戏云设置备份选项
//
// id: RGS ID
//
// 没错就是Rcs,这俩是一样的
func (c *Client) EnableRgsAutoBackup(id int, req *rcs.RcsSetBackupOptionsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/backup/setting", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, req, &resp)

	return &resp, err
}

// 游戏云重装系统
//
// id: RGS ID
func (c *Client) Reinstallgs(id int, req *rcs.ReinstallRcsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/changeos", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 游戏云限制模式(是否用余额结算)切换
//
// id: 游戏云 ID
//
// useMoney: 是否用余额结算CPU电量
func (c *Client) SwitchRgsBalanceMode(id int, useMoney bool) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/cpu-limit-mode", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, common.SwitchModeRequest{Mode: useMoney}, &resp)

	return &resp, err
}

// 游戏云CPU充电
//
// id: 游戏云 ID
func (c *Client) ChargeRgsCPU(id int, req *rcs.ChangeRcsIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/cpu-charge", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 游戏云日付模式开关
//
// id： 游戏云 ID
//
// dailyMode: true: 开启日付模式，false: 关闭日付模式
func (c *Client) SwitchRgsDailyMode(id int, dailyMode bool) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/daily-mode", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, common.SwitchModeRequest{Mode: dailyMode}, &resp)

	return &resp, err
}

// 创建并绑定弹性IP到游戏云
//
// id: 游戏云 ID
func (c *Client) CreateAndBindElasticIpToRgs(id int, req *rcs.CreateAndBindIpToRcsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/eip/", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 更换游戏云IP
//
// id: RCS ID
func (c *Client) ChangeRgsIP(id int, req *rcs.ChangeRcsIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/eip/change", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 放弃游戏云IP
//
// id: 游戏云ID
func (c *Client) DisCardRgsIP(id int, req rcs.DisCardRcsIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/eip/discard", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 获取游戏云防火墙规则列表
//
// id: 游戏云 ID
//
// options: 游戏云查询参数 可以用 EncodingStandardQueryParameters 获取.
func (c *Client) GetRgsFirewallRules(id int, options string) (*rcs.RcsFirewallRuleList, error) {
	path := fmt.Sprintf("/product/rgs/%d/firewall/rule?options=%s", id, options)

	var resp rcs.RcsFirewallRuleList
	err := c.DoRequest("GET", path, options, &resp)

	return &resp, err
}

// 创建/设置游戏云防火墙规则
//
// id: 游戏云 ID
func (c *Client) SetRgsFirewallRule(id int, req *rcs.SetRcsFirewallRuleRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/firewall/rule", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 删除游戏云防火墙规则
//
// id: 游戏云 ID
func (c *Client) DeleteRgsFirewallRule(id int, ruleID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/firewall/rule/%d", id, ruleID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("DELETE", path, nil, &resp)

	return &resp, err
}

// 移动游戏云防火墙规则优先级
//
// id: 游戏云 ID
func (c *Client) MobileRgsFirewallRulePriority(id int, req rcs.MobileRcsFirewallRulePriorityRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%dfirewall/rule/{ruleId}/pos", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PUT", path, req, &resp)

	return &resp, err
}

// 释放游戏云
//
// id: 游戏云ID
func (c *Client) FreeRgs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/free", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 获取游戏云监控数据
//
// id: 游戏云ID
//
// startDate: 开始时间
//
// endDate: 结束时间
func (c *Client) GetRgsMonitorData(id int, startDate int, endDate int) (*rcs.RcsMonitoringData, error) {
	path := fmt.Sprintf("/product/rgs/%d/monitor?start_date=%d&end_date=%d", id, startDate, endDate)

	var resp rcs.RcsMonitoringData
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 添加游戏云NAT端口映射
//
// id: 游戏云 ID
func (c *Client) AddRgsNatPortMapping(id int, req *rcs.AddRcsNatPortMappingRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/nat", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 游戏云重启操作
//
// id: RCS ID
func (c *Client) RebootRgs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/reboot", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 获取游戏云续费价格
//
// id: RCS ID, duration: 续费时长(月), coupon: 优惠券ID
func (c *Client) GetRgsRenewPrice(id int, duration int, couponID int) (*rcs.RCSRenewPrice, error) {
	path := fmt.Sprintf("/product/rgs/price?scene=renew&product_id=%d&duration=%d&with_coupon_id=%d&is_old=true", id, duration, couponID)

	var resp rcs.RCSRenewPrice
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 游戏云续费
//
// id: 游戏云ID
func (c *Client) RenewRgs(id int, req rcs.RenewRcsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/renew/", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 游戏云自动续费选项
//
// id: 游戏云ID
func (c *Client) EnableRgsAutoRenew(id int, req rcs.EnableRcsAutoRenewRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/renew/option", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 游戏云重置密码
//
// id: 游戏云ID
//
// newPass: 新密码,留空则自动生成
func (c *Client) ResetRgsPassword(id int, newPass string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/reset-password", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, rcs.ResetRcsPasswordRequest{Password: newPass}, &resp)

	return &resp, err
}

// 游戏云开机
//
// id: 游戏云ID
func (c *Client) StartRgs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/start", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 游戏云关机
//
// id: 游戏云ID
func (c *Client) StopRgs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/stop", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 设置游戏云标签
//
// id: 游戏云ID
//
// tag: 标签
func (c *Client) SetRgsTag(id int, tag string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/tag", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, rcs.SetRcsTagRequest{TagName: tag}, &resp)

	return &resp, err
}

// 游戏云升级
//
// id: 游戏云ID
//
// plan: 升级到的套餐ID
//
// coupon: 优惠券ID,默认为0
func (c *Client) UpgradeRgs(id int, plan int, coupon int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/upgrade", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, rcs.UpgradeRcsRequest{DestPlan: plan, WithCouponID: coupon}, &resp)

	return &resp, err
}

// 游戏云连接VNC
//
// id: 游戏云ID
//
// consoleType: 控制台类型,可选值: novnc,xtermjs
func (c *Client) GetRgsVnc(id int, consoleType string) (*common.VncConnectionInfo, error) {
	path := fmt.Sprintf("/product/rgs/%d/vnc?console_type=%s", id, consoleType)

	var resp common.VncConnectionInfo
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取游戏云对应的PVE地址
//
// id: 游戏云ID
func (c *Client) GetRgsPveAddress(id int) (pveAddress string, err error) {
	v, err := c.GetRgsVnc(id, "novnc")
	if err != nil {
		return "", err
	}

	parsedURL, err := url.Parse(v.Data.RequestURL)
	if err != nil {
		return "", err
	}

	return parsedURL.Hostname(), err
}

// 创建MCSM面板用户
//
// name: 用户名
//
// password: 密码
func (c *Client) CreateMcsmUser(name string, password string) (*common.BasicOperationResponse, error) {
	path := "/product/rgs/mcsm/panel_user/"

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, McsmUser{Name: name, Password: password}, &resp)

	return &resp, err
}

// 编辑MCSM面板用户
//
// name: 用户名
//
// password: 密码
func (c *Client) EditMcsmUser(name string, password string) (*common.BasicOperationResponse, error) {
	path := "/product/rgs/mcsm/panel_user/"

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, McsmUser{Name: name, Password: password}, &resp)

	return &resp, err
}

// 删除MCSM面板用户
//
// name: 用户名
func (c *Client) DeleteMcsmUser(name string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/mcsm/panel_user/%s", name)

	var resp common.BasicOperationResponse
	err := c.DoRequest("DELETE", path, nil, &resp)

	return &resp, err
}

// 获取游戏云升级价格
//
// id: RCS ID, duration: 续费时长(月), coupon: 优惠券ID
func (c *Client) GetRgsUpgradePrice(id int, duration int, couponID int, config *RgsConfig) (*RgsUpgradePrice, error) {
	configByte, err := json.Marshal(config)
	configString := string(configByte)
	path := fmt.Sprintf("/product/rgs/price?scene=upgrade&product_id=%d&duration=%d&with_coupon_id=%d&is_old=true&config=%s", id, duration, couponID, configString)

	var resp RgsUpgradePrice
	err = c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 游戏云更换egg(游戏类型)
//
// 此操作需要二步验证
//
// eggTypeID: 游戏类型ID
//
// saveDirs: 要保留的目录
func (c *Client) ChangeRgsEgg(id int, eggTypeID int, saveDirs []string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/change-egg", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, ChangeRgsEggRequest{EggTypeID: eggTypeID, SaveDirs: saveDirs}, &resp)

	return &resp, err
}
