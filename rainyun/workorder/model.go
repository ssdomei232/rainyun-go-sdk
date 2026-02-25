package workorder

// 工单列表
type WorkorderList struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords int `json:"TotalRecords"`
		Records      []struct {
			ID                 int    `json:"ID"`  // 工单ID
			UID                int    `json:"UID"` // 用户ID
			ExtUserInfo        string `json:"ExtUserInfo"`
			UserName           string `json:"UserName"` // 用户名
			UserEmail          string `json:"UserEmail"`
			UserVip            string `json:"UserVip"`            // 用户vip等级
			UserIcon           string `json:"UserIcon"`           // 用户头像
			Title              string `json:"Title"`              // 工单标题
			Type               string `json:"Type"`               // 工单类型
			RelatedProductType string `json:"RelatedProductType"` // 关联产品类型
			RelatedProductID   int    `json:"RelatedProductID"`   // 关联产品id
			IsUrgent           int    `json:"IsUrgent"`           // 是否为紧急工单
			Status             string `json:"Status"`             // 状态(finished/answered/waiting)
			Time               int    `json:"Time"`               // 工单创建时间
			LastTime           int    `json:"LastTime"`
			WaitBeginTime      int    `json:"WaitBeginTime"`
			AssistID           int    `json:"AssistID"` // 客服id
			AuthStatus         string `json:"AuthStatus"`
			AuthTime           int    `json:"AuthTime"`
			AuthID             int    `json:"AuthID"`
		} `json:"Records"`
	} `json:"data"`
}

// 工单详情
type WorkorderDetail struct {
	Code int `json:"code"`
	Data struct {
		ID                 int    `json:"ID"`
		UID                int    `json:"UID"`
		ExtUserInfo        string `json:"ExtUserInfo"`
		UserName           string `json:"UserName"`
		UserEmail          string `json:"UserEmail"`
		UserVip            string `json:"UserVip"`
		UserIcon           string `json:"UserIcon"`
		Title              string `json:"Title"`
		Content            string `json:"Content"`
		Type               string `json:"Type"`
		RelatedProductType string `json:"RelatedProductType"`
		RelatedProductID   int    `json:"RelatedProductID"`
		IsUrgent           int    `json:"IsUrgent"`
		Status             string `json:"Status"`
		Time               int    `json:"Time"`
		LastTime           int    `json:"LastTime"`
		WaitBeginTime      int    `json:"WaitBeginTime"`
		AssistID           int    `json:"AssistID"`
		AuthStatus         string `json:"AuthStatus"`
		AuthTime           int    `json:"AuthTime"`
		AuthID             int    `json:"AuthID"`
		Discuss            []struct {
			ID             int    `json:"ID"`
			IsAssist       bool   `json:"IsAssist"`
			UID            int    `json:"UID"`
			UserName       string `json:"UserName"`
			UserEmail      string `json:"UserEmail"`
			UserVip        string `json:"UserVip"`
			UserIcon       string `json:"UserIcon"`
			Content        string `json:"Content"`
			Time           int    `json:"Time"`
			WaitTime       int    `json:"WaitTime"`
			LastEditedTime int    `json:"LastEditedTime"`
			IsScored       bool   `json:"IsScored"`
		} `json:"Discuss"`
	} `json:"data"`
}

// 创建工单
type CreateWorkerorderRequest struct {
	Content            string `json:"content"`
	IsAuthed           bool   `json:"is_authed"`            // 是否需要授权(可选)
	IsUrgent           int    `json:"is_urgent"`            // 是否为紧急工单
	RelatedProductID   int    `json:"related_product_id"`   // 关联产品类型(可选)
	RelatedProductType string `json:"related_product_type"` // 关联产品id(可选)
	Title              string `json:"title"`
	Type               string `json:"type"` // 工单类型
}

// 创建工单响应
type CreateWorkerorderResponse struct {
	Code int `json:"code"`
	Data struct {
		ID                 int    `json:"ID"`
		UID                int    `json:"UID"`
		ExtUserInfo        string `json:"ExtUserInfo"`
		UserName           string `json:"UserName"`
		UserEmail          string `json:"UserEmail"`
		UserVip            string `json:"UserVip"`
		UserIcon           string `json:"UserIcon"`
		Title              string `json:"Title"`
		Content            string `json:"Content"`
		Type               string `json:"Type"`
		RelatedProductType string `json:"RelatedProductType"`
		RelatedProductID   int    `json:"RelatedProductID"`
		IsUrgent           int    `json:"IsUrgent"`
		Status             string `json:"Status"`
		Time               int    `json:"Time"`
		LastTime           int    `json:"LastTime"`
		WaitBeginTime      int    `json:"WaitBeginTime"`
		AssistID           int    `json:"AssistID"`
		AuthStatus         string `json:"AuthStatus"`
		AuthTime           int    `json:"AuthTime"`
		AuthID             int    `json:"AuthID"`
	} `json:"data"`
}

// 产品授权请求
type ProductAuthRequest struct {
	ProductID   int    `json:"product_id"`   // 产品ID
	ProductType string `json:"product_type"` // 产品类型(rvh/rcs/rgs/rbm/ros)
}

// 回复工单请求
type ReplyWorkerorderRequest struct {
	Content string `json:"content"` // 回复内容
}

// 工单打分
type ScoreWorkerorderRequest struct {
	Score     int    `json:"score"`      // 分数(1-5)
	Reason    string `json:"reason"`     // (可选)
	IsSolved  bool   `json:"is_solved"`  // 是否解决(可选)
	DiscussID int    `json:"discuss_id"` // 回复ID(可选)
	Aid       int    `json:"aid"`        // 客服id(可选)
}

// 工单打分详情
type ScoreWorkerorderDetail struct {
	Code int `json:"code"`
	Data struct {
		UID         int    `json:"uid"`
		Aid         int    `json:"aid"`
		OrderID     int    `json:"order_id"`
		DiscussID   int    `json:"discuss_id"`
		Score       int    `json:"score"`
		Reason      string `json:"reason"`
		IsSolved    bool   `json:"is_solved"`
		DiscussTime int    `json:"DiscussTime"`
		ScoreTime   int    `json:"score_time"`
	} `json:"data"`
}

// 工单状态
type WorkorderStatus struct {
	Code int `json:"code"`
	Data struct {
		LastTime int    `json:"LastTime"`
		Status   string `json:"Status"`
	} `json:"data"`
}

// 设置工单状态请求
type SetWorkorderStatusRequest struct {
	Status string `json:"status"`
}
