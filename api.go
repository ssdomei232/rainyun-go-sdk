package rainyun_go_sdk

// GetUserInfo 获取用户信息
func (c *Client) GetUserInfo() (*UserInfo, error) {
	uri := "/user/"

	var resp UserInfo
	err := c.DoRequest("GET", uri, nil, &resp)

	return &resp, err
}

func (c *Client) GetUserRewardPruducts() (*UserRewardProducts, error) {
	uri := "/user/reward/products"

	var resp UserRewardProducts
	err := c.DoRequest("GET", uri, nil, &resp)

	return &resp, err
}
