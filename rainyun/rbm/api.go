package rbm

import (
	"fmt"

	"github.com/ssdomei232/rainyun-go-sdk/v2/rainyun/common"
)

// 获取RBM套餐列表
func (c *Client) GetRBMPlanList() (*RBMPlanList, error) {
	path := "/product/rbm/models"

	var resp RBMPlanList
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// RBM实例更换系统
func (c *Client) ChangeRBMOS(id int, req *RBMChangeOSRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/changeos", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 创建并绑定弹性IP到RBM
func (c *Client) AssociateEIP(id int, req *RBMAssociateEIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/eip", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 设置IP描述
func (c *Client) SetIPDescription(id int, req *RBMSetIPDescriptionRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/eip/description", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 放弃IP
func (c *Client) ReleaseIP(id int, req *RBMReleaseIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/eip/discard", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// RBM实例启动KVM代理
func (c *Client) StartKVMAgent(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/kvm-proxy", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// RBM重新启动KVM
func (c *Client) RestartKVM(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/kvm-reboot", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// RBM实例关机
func (c *Client) ShutdownRBM(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/poweroff", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// RBM实例开机
func (c *Client) StartRBM(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/poweron", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 充流量
func (c *Client) ChargeTraffic(id int, req *RBMChargeTrafficRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/traffic/charge", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// TODO: 列出RBM实例,创建RBM实例,裸金属刷bios,RBM清点配置,获取监控数据,重置RBM实例IPMI密码,限流,切换流量套餐,获取使用情况列表
