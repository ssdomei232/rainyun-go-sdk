package user

import "github.com/ssdomei232/rainyun-go-sdk/v2/rainyun/common"

// 获取用户信息.
func (c *Client) GetUserInfo() (*UserInfo, error) {
	path := "/user/"

	var resp UserInfo
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 刷新 api 密钥.
//
// ⚠️不要轻易使用此方法
func (c *Client) RefreshApikey() (*common.BasicOperationResponse, error) {
	path := "/user/"

	options := `{"option":"apikey"}`

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, options, &resp)

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
func (c *Client) PublishCouponsToLowerLevelUsers(req *PublishCouponsToLowerLevelUsersRequest) (*common.BasicOperationResponse, error) {
	path := "/user/vip/coupon"

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, req, &resp)

	return &resp, err
}

// 发布优惠券到积分商城(供下级领取)
func (c *Client) PostCouponsToPointsMall(req *PostCouponsToPointsMallRequest) (*common.BasicOperationResponse, error) {
	path := "/user/vip/coupon"

	var resp common.BasicOperationResponse
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
func (c *Client) RedeemPointsForItem(id int) (*common.BasicOperationResponse, error) {
	path := "/user/reward/items"

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, RedeemPointsForItemRequest{ItemID: id}, &resp)

	return &resp, err
}

// 请求二次验证
//
// verificationMethod: 验证方式(sms/totp/email)
func (c *Client) Request2FA(verificationMethod string) (*common.BasicOperationResponse, error) {
	path := "/user/mfa/request"

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, Request2FARequest{Type: verificationMethod}, &resp)

	return &resp, err

}

// 验证二次验证结果
//
// authCode: 验证码
func (c *Client) Verify2FAResult(authCode int) (*common.BasicOperationResponse, error) {
	path := "/user/mfa/verify"

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, Verify2FAResultRequest{AuthCode: authCode}, &resp)

	return &resp, err
}
