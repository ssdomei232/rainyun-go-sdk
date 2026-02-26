package product

// 产品数量列表
type ProductCountList struct {
	Code int `json:"code"`
	Data struct {
		Ros struct {
			TotalCount int `json:"TotalCount"`
		} `json:"ros"`
		Rvh struct {
			TotalCount int `json:"TotalCount"`
		} `json:"rvh"`
		Rcs struct {
			TotalCount int `json:"TotalCount"`
		} `json:"rcs"`
		Rgpu struct {
			TotalCount int `json:"TotalCount"`
		} `json:"rgpu"`
		Rgs struct {
			TotalCount int `json:"TotalCount"`
		} `json:"rgs"`
		Ssl struct {
			TotalCount int `json:"TotalCount"`
		} `json:"ssl"`
		Rbm struct {
			TotalCount    int         `json:"TotalCount"`
			AboutToExpire interface{} `json:"AboutToExpire"`
		} `json:"rbm"`
		Domain struct {
			TotalCount int `json:"TotalCount"`
		} `json:"domain"`
		Rcdn struct {
			TotalCount int `json:"TotalCount"`
		} `json:"rcdn"`
		Rca struct {
			TotalCount int `json:"TotalCount"`
		} `json:"rca"`
	} `json:"data"`
}

// 产品ID列表
type ProductIDList struct {
	Code int `json:"code"`
	Data struct {
		Rca    []int `json:"rca"`
		Rcs    []int `json:"rcs"`
		Ros    []int `json:"ros"`
		Rvh    []int `json:"rvh"`
		Rgpu   []int `json:"rgpu"`
		Rgs    []int `json:"rgs"`
		Ssl    []int `json:"ssl"`
		Rbm    []int `json:"rbm"`
		Domain []int `json:"domain"`
		Rcdn   []int `json:"rcdn"`
	} `json:"data"`
}

// 增减面板用户产品请求
type ModifyPanelUserProductRequests struct {
	Action      string `json:"action"`       // 操作：add/del
	ProductID   int    `json:"product_id"`   // 产品ID
	ProductType string `json:"product_type"` // 产品类型: rcs/rgs/ros/rbm 只支持这四种产品
	Name        string `json:"name"`         // 子用户名
}

// 面板用户列表
type PanelUserList struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords int `json:"TotalRecords"` // 面板用户总数
		Records      []struct {
			Name       string `json:"name"`        // 面板用户名
			Password   string `json:"password"`    // 面板用户密码
			UserID     int    `json:"user_id"`     // 你的UID
			CreateDate int    `json:"create_date"` // 创建时间
			Products   []struct {
				Name        string `json:"name"`         // 自用户名
				ProductType string `json:"product_type"` // 产品类型: rcs/rgs/ros/rbm 只支持这四种产品
				ProductID   int    `json:"product_id"`   // 产品ID
			} `json:"products"` // 产品列表
		} `json:"Records"`
	} `json:"data"`
}

// 面板用户请求
type PanelUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// 积分续费请求
type PointRenewalRequest struct {
	DurationDay int    `json:"duration_day"` // 续费天数
	ProductID   int    `json:"product_id"`   // 产品ID
	ProductType string `json:"product_type"` // 产品类型: rcs/rgs/ros/rvh/rcdn
}
