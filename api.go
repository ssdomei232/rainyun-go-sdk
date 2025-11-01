package rainyun_go_sdk

/* =============用户信息部分============= */

// GetUserInfo 获取用户信息.
func (c *Client) GetUserInfo() (*UserInfo, error) {
	path := "/user/"

	var resp UserInfo
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// GetUserRewardPruducts 获取可兑换积分产品列表.
func (c *Client) GetUserRewardPruducts() (*UserRewardProducts, error) {
	path := "/user/reward/products"

	var resp UserRewardProducts
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// GetUserLogs 获取用户日志.
// options: 标准查询参数 可以用 MarshalStandardQueryParameters 获取.
// log_type: 日志类型: "user/": 账号变动日志, "consume/": 消费记录, "user/expense/unsubscribe": 退订记录.
func (c *Client) GetUserLogs(options string, log_type string) (*UserLogsResponse, error) {
	path := "/user/logs"

	var resp UserLogsResponse
	err := c.DoRequest("GET", path, map[string]string{
		"options":  options,
		"log_type": log_type,
	}, &resp)

	return &resp, err
}
