package workorder

import (
	"fmt"

	"github.com/ssdomei232/rainyun-go-sdk/rainyun/common"
)

// 获取工单列表
//
// options: 查询参数 可以用 EncodingStandardQueryParameters 获取.
func (c *Client) GetWorkOrderList(options string) (*WorkorderList, error) {
	path := fmt.Sprintf("/workorder/?options=%s", options)

	var resp WorkorderList
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取工单详情
func (c *Client) GetWorkOrderDetail(id int) (*WorkorderDetail, error) {
	path := fmt.Sprintf("/workorder/%d", id)

	var resp WorkorderDetail
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 创建工单
func (c *Client) CreateWorkOrder(req *CreateWorkerorderRequest) (*CreateWorkerorderResponse, error) {
	path := "/workorder/"

	var resp CreateWorkerorderResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 产品授权
//
// id:工单ID
//
// productID:产品ID
//
// productType:产品类型(rvh/rcs/rgs/rbm/ros)
func (c *Client) ProductAuth(id int, productID int, productType string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/workorder/%d/auth", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, ProductAuthRequest{ProductID: productID, ProductType: productType}, &resp)

	return &resp, err
}

// 回复工单
//
// id: 工单ID
//
// content: 回复内容
func (c *Client) ReplyWorkorder(id int, content string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/workorder/%d/reply_order", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, ReplyWorkerorderRequest{Content: content}, &resp)

	return &resp, err
}

// 编辑回复工单
//
// id: 工单ID
//
// replyID: 回复ID
//
// content: 编辑后的内容
func (c *Client) EditWorkorderReply(id int, replyID int, content string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/workorder/%d/reply_order/%d", id, replyID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, ReplyWorkerorderRequest{Content: content}, &resp)

	return &resp, err
}

// 工单打分
//
// id: 工单ID
func (c *Client) ScoreWorkorder(id int, req *ScoreWorkerorderRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/workorder/%d/score", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 编辑工单打分
//
// id: 工单ID
func (c *Client) EditScoreWorkorder(id int, req *ScoreWorkerorderRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/workorder/%d/score", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, req, &resp)

	return &resp, err
}

// 获取工单打分
//
// id: 工单ID
//
// discussID: 客服回复ID
func (c *Client) GetScoreWorkorder(id int, discussID int) (*ScoreWorkerorderDetail, error) {
	path := fmt.Sprintf("/workorder/%d/score/%d", id, discussID)

	var resp ScoreWorkerorderDetail
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取工单状态
//
// id: 工单ID
func (c *Client) GetWorkorderStatus(id int) (*WorkorderStatus, error) {
	path := fmt.Sprintf("/workorder/%d/status", id)

	var resp WorkorderStatus
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 设置工单状态
func (c *Client) SetWorkorderStatus(id int, status string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/workorder/%d/status", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, SetWorkorderStatusRequest{Status: status}, &resp)

	return &resp, err
}
