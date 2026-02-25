package rbm

// 获取RBM套餐列表
func (c *Client) GetRBMPlanList() (*RBMPlanList, error) {
	path := "/product/rbm/models"

	var resp RBMPlanList
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}
