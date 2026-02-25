package product

import (
	"fmt"

	"github.com/ssdomei232/rainyun-go-sdk/v2/rainyun/common"
)

// 获取各产品数量
func (c *Client) GetProductCountList() (*ProductCountList, error) {
	path := "/product/"

	var resp ProductCountList
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取产品ID列表
func (c *Client) GetProductIDList() (*ProductIDList, error) {
	path := "/product/id_list"

	var resp ProductIDList
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取面板用户列表
//
// options: 查询参数 可以用 EncodingStandardQueryParameters 获取.
func (c *Client) GetPanelUserList(options string) (*PanelUserList, error) {
	path := "/product/panel_users/?options=" + options

	var resp PanelUserList
	err := c.DoRequest("GET", path, options, &resp)

	return &resp, err
}

// 增减面板用户产品
func (c *Client) ModifyPanelUserProduct(req *ModifyPanelUserProductRequests) (*common.BasicOperationResponse, error) {
	path := "/product/panel_users/"

	var resp common.BasicOperationResponse
	err := c.DoRequest("PUT", path, req, &resp)

	return &resp, err

}

// 创建面板用户
//
// name： 子用户用户名
//
// pass： 子用户密码
func (c *Client) CreatePanelUser(name string, pass string) (*common.BasicOperationResponse, error) {
	path := "/product/panel_users/"

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, PanelUserRequest{Name: name, Password: pass}, &resp)

	return &resp, err
}

// 面板用户改密
//
// name： 子用户用户名
//
// pass： 子用户新密码
func (c *Client) ChangePanelUserPassword(name string, pass string) (*common.BasicOperationResponse, error) {
	path := "/product/panel_users/"

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, PanelUserRequest{Name: name, Password: pass}, &resp)

	return &resp, err
}

// 删除面板用户
//
// name： 子用户用户名
func (c *Client) DeletePanelUser(name string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/panel_users/%s", name)

	var resp common.BasicOperationResponse
	err := c.DoRequest("DELETE", path, nil, &resp)

	return &resp, err
}
