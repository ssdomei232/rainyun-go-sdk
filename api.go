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

/* ==============RCS============== */

// 设置IP描述
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
