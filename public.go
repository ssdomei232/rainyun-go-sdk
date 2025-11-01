package rainyun_go_sdk

import "time"

// 页面信息
type AppConfig struct {
	Code int `json:"code"`
	Data []struct {
		AgentID int    `json:"agent_id"`
		Type    string `json:"type"`
		Value   []struct {
			SenderID      int    `json:"sender_id"`
			Content       string `json:"content"`
			Title         string `json:"title"`
			Name          string `json:"name"`
			Order         int    `json:"order"`
			Page          string `json:"page"`
			VgtID         int    `json:"vgt_id"`
			OriginalIndex int    `json:"originalIndex"`
		} `json:"value"`
		Date int `json:"date"`
	} `json:"data"`
}

// 论坛公告
type RainyunForumNews struct {
	Code int `json:"code"`
	Data []struct {
		Type      string    `json:"Type"`
		Title     string    `json:"Title"`
		TimeStamp time.Time `json:"TimeStamp"`
		URL       string    `json:"URL"`
	} `json:"data"`
}

// 节点状态
type NodeStatus struct {
	Code int `json:"code"`
	Data struct {
		TotalRecords         int           `json:"TotalRecords"`
		OngoingServiceEvents []interface{} `json:"OngoingServiceEvents"`
		Records              []struct {
			UUID        string  `json:"UUID"`
			ChineseName string  `json:"ChineseName"`
			Product     string  `json:"Product"`
			CPU         float64 `json:"CPU"`
			Memory      float64 `json:"Memory"`
			NetOut      int     `json:"NetOut"`
			UpdateTime  string  `json:"UpdateTime"`
			Status      string  `json:"Status"`
			Data        string  `json:"Data"`
		} `json:"Records"`
	} `json:"data"`
}

// GetAppConfig 获取页面信息.
func GetAppConfig() (*AppConfig, error) {
	path := "/app_config"

	var resp AppConfig
	err := publicDoRequest("GET", path, nil, &resp)

	return &resp, err
}

// GetRainyunForumNews 获取论坛公告.
func GetRainyunForumNews() (*RainyunForumNews, error) {
	path := "/news"

	var resp RainyunForumNews
	err := publicDoRequest("GET", path, nil, &resp)

	return &resp, err
}

// GetNodeStatus 获取节点状态.
// options: 标准查询参数 可以用 MarshalStandardQueryParameters 获取.
func GetNodeStatus(options string) (*NodeStatus, error) {
	path := "/status"

	var resp NodeStatus
	err := publicDoRequest("GET", path, options, &resp)

	return &resp, err
}

// publicDoRequest 发送不需要登录的请求.
func publicDoRequest(method, endpoint string, reqData any, respData any) error {
	client := NewClient("")

	return client.DoRequest(method, endpoint, reqData, respData)
}
