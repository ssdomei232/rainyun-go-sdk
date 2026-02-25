package rca

import (
	"fmt"

	"github.com/ssdomei232/rainyun-go-sdk/v2/rainyun/common"
)

// 云应用获取区域信息
func (c *Client) GetRcaRegionInfo() (*RcaRegionInfo, error) {
	path := "/product/rca/region"

	var resp RcaRegionInfo
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 云应用获取雨点余额使用情况
func (c *Client) GetRcaRaindropUsage() (*RcaRaindropUsage, error) {
	path := "/product/rca/raindrop/usage"

	var resp RcaRaindropUsage
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 云应用获取雨点套餐列表
func (c *Client) GetRcaRaindropPlansList() (*RaindropPlansList, error) {
	path := "/product/rca/raindrop/plans"

	var resp RaindropPlansList
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 云应用获取雨点消费历史
//
// options: 查询参数 可以用 EncodingStandardQueryParameters 获取.
func (c *Client) GetRaindropConsumeLog(options string) (*RaindropConsumeLog, error) {
	path := fmt.Sprintf("/product/rca/raindrop/consume_log?options=%s", options)

	var resp RaindropConsumeLog
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 云应用购买雨点
//
// planID: 雨点套餐ID
//
// couponID: 优惠券ID
func (c *Client) BuyRaindrop(planID int, couponID int) (*common.BasicOperationResponse, error) {
	path := "/product/rca/raindrop"

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, BuyRaindropRequest{PlanID: planID, WithCouponID: couponID}, &resp)

	return &resp, err
}

// 开通云应用产品
//
// regionID: 地域ID
func (c *Client) ActivateRca(regionID int) (*common.BasicOperationResponse, error) {
	path := "/product/rca/activate"

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, ActivateRcaRequest{RegionID: regionID}, &resp)

	return &resp, err
}

// 云应用获取雨点余额
func (c *Client) GetRcaRaindropBalance() (*RaindropBalance, error) {
	path := "/product/rca/raindrop"

	var resp RaindropBalance
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 创建云应用项目
func (c *Client) CreateRcaProject(req *CreateRcaProjectRequest) (*CreateRcaProjectResponse, error) {
	path := "/product/rca/project/"

	var resp CreateRcaProjectResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 云应用列出项目
//
// options: 查询参数 可以用 EncodingStandardQueryParameters 获取.
func (c *Client) ListRcaProjects(options string) (*RcaProjectList, error) {
	path := fmt.Sprintf("/product/rca/project/?no_metrics=false&options=%s", options) // no_metrics 含义不明

	var resp RcaProjectList
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 销毁云应用项目
//
// id: RCA项目ID
func (c *Client) DestroyRcaProject(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("DELETE", path, nil, &resp)

	return &resp, err
}

// 获取Rca项目的指标信息
//
// id: RCA项目ID
//
// starttime: 开始时间(timestamp)
//
// endtime: 结束时间(timestamp)
func (c *Client) GetRcaProjectMetrics(id int, startTime int, endTime int) (*RcaProjectMetrics, error) {
	path := fmt.Sprintf("/product/rca/project/%d/metrics?start_time=%d&end_time=%d", id, startTime, endTime)

	var resp RcaProjectMetrics
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取云应用项目详情
//
// id: RCA项目ID
func (c *Client) GetRcaProjectDetail(id int) (*RcaProjectDetails, error) {
	path := fmt.Sprintf("/product/rca/project/%d/", id)

	var resp RcaProjectDetails
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 云应用项目设置备份目标
//
// id: RCA项目ID
func (c *Client) SetRcaProjectBackupTarget(id int, req *SetRcaProjectBackupTargetRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/backup_target", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, req, &resp)

	return &resp, err
}

// 云应用项目磁盘扩容
//
// id: RCA项目ID
func (c *Client) ExpandRcaProjectDisk(id int, req *RcaProjectDiskExpansionRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/disk_expand", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 云应用增加IP地址
//
// id: RCA项目ID
func (c *Client) AddRcsProjectIP(id int, req *RcaAddsIpAddressRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/eip", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 云应用移除IP地址
//
// id: RCA项目ID
func (c *Client) RemoveRcaProjectIP(id int, req *RcaRemoveIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/eip", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("DELETE", path, req, &resp)

	return &resp, err
}

// 云应用项目修改SFTP设置
//
// id: RCA项目ID
//
// password: 密码
//
// username: 用户名
func (c *Client) SetRcaProjectSFTPConfig(id int, password string, username string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/sftp", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, RcaProjectSetSftpConfigRequest{Password: password, Username: username}, &resp)

	return &resp, err
}

// 云应用项目列出IP地址
//
// id: RCA项目ID
func (c *Client) ListRcaProjectIPs(id int) (*RcaIPInfo, error) {
	path := fmt.Sprintf("/product/rca/project/%d/eip", id)

	var resp RcaIPInfo
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 云应用创建App模板
func (c *Client) CreateRcaAppTemplate(req *RcaCreateAppTemplateRequest) (*RcaCreateAppTemplateResponse, error) {
	path := "/product/rca/appstore/"

	var resp RcaCreateAppTemplateResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 创建App模板版本
//
// id: APP模板ID
func (c *Client) CreateRcaAppTemplateVersion(id int, req *CreateAppTemplateVersionRequest) (*CreateAppTemplateVersionResponse, error) {
	path := fmt.Sprintf("/product/rca/appstore/%d/release", id)

	var resp CreateAppTemplateVersionResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}
