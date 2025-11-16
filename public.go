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

// 蛋(游戏)列表
type EggList struct {
	Code int `json:"code"`
	Data []struct {
		Name     string   `json:"name"`      // 名称
		EggGroup string   `json:"egg_group"` // 分组
		Title    string   `json:"title"`     // 标题
		Desc     string   `json:"desc"`      // 描述
		IconURL  string   `json:"icon_url"`  // 图标
		Order    int      `json:"order"`     // 排序
		IsHidden bool     `json:"is_hidden"` // 是否隐藏
		SaveDirs []string `json:"save_dirs"` // 重装时建议的需要保留的目录
	} `json:"data"`
}

// 蛋(游戏类型)类型列表
type EggTypeList struct {
	Code int `json:"code"`
	Data []struct {
		ID      int    `json:"id"`       // 蛋(游戏类型)ID
		EggName string `json:"egg_name"` // 蛋(游戏类型)名称
		Egg     struct {
			Name     string   `json:"name"`
			EggGroup string   `json:"egg_group"`
			Title    string   `json:"title"`
			Desc     string   `json:"desc"`
			IconURL  string   `json:"icon_url"`
			Order    int      `json:"order"`
			IsHidden bool     `json:"is_hidden"`
			SaveDirs []string `json:"save_dirs"`
		} `json:"egg"`
		Docker     string `json:"docker"` // docker镜像地址
		McsmDocker string `json:"mcsm_docker"`
		Env        struct {
			ServerType    string `json:"SERVER_TYPE"`
			ServerVersion string `json:"SERVER_VERSION"`
			Subver        string `json:"SUBVER"`
			Latest        bool   `json:"LATEST"`
			Prerelease    bool   `json:"PRERELEASE"`
		} `json:"env"`
		Order      int  `json:"order"`
		IsHidden   bool `json:"is_hidden"`
		UpdateTime int  `json:"update_time"`
		AutoUpdate bool `json:"auto_update"`
	} `json:"data"`
}

// 游戏云系统列表
type RgsOSList struct {
	Code int `json:"code"`
	Data []struct {
		ID             int    `json:"id"`
		Region         string `json:"region"`
		Subtype        string `json:"subtype"`
		Machine        string `json:"machine"`
		Name           string `json:"name"`
		Version        string `json:"version"`
		SyncStatus     string `json:"sync_status"`
		OsType         string `json:"os_type"`
		ChineseName    string `json:"chinese_name"`
		Icon           string `json:"icon"`
		IsWithBbr      bool   `json:"is_with_bbr"`
		IsEol          bool   `json:"is_eol"`
		IsAvailable    bool   `json:"is_available"`
		Order          int    `json:"order"`
		LatestFilename string `json:"latest_filename"`
		NoVMAgent      bool   `json:"no_vm_agent"`
	} `json:"data"`
}

// 对象存储套餐列表
type RosPlanList struct {
	Code int `json:"code"`
	Data []struct {
		ID                 int     `json:"id"`
		Region             string  `json:"region"`
		Subtype            string  `json:"subtype"`
		PlanName           string  `json:"plan_name"`
		Machine            string  `json:"machine"`
		ChargeType         string  `json:"charge_type"`
		Chinese            string  `json:"chinese"`
		IsFree             bool    `json:"is_free"`
		PointRenewPrice    any     `json:"point_renew_price"` // 积分续费价格
		IsSelling          bool    `json:"is_selling"`
		Price              int     `json:"price"`
		StorageSize        int     `json:"storage_size"`
		Bandwidth          int     `json:"bandwidth"`
		ExtraTransferPrice float64 `json:"extra_transfer_price"`
		ExtraStoragePrice  float64 `json:"extra_storage_price"`
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

// 获取RCS操作系统列表
func GetRcsOSList() (*RcsOSList, error) {
	path := "/product/rcs/os"

	var resp RcsOSList
	err := publicDoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取蛋(游戏)列表
func GetEggList() (*EggList, error) {
	path := "/product/rgs/egg"

	var resp EggList
	err := publicDoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取蛋(游戏类型)类型列表
func GetEggTypeList() (*EggTypeList, error) {
	path := "/product/rgs/egg_type"

	var resp EggTypeList
	err := publicDoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取游戏云系统列表
func GetRgsOSList() (*RgsOSList, error) {
	path := "/product/rgs/os-templates"

	var resp RgsOSList
	err := publicDoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取对象存储套餐列表
func GetRosPlanList() (*RosPlanList, error) {
	path := "/product/ros/plans"

	var resp RosPlanList
	err := publicDoRequest("GET", path, nil, &resp)

	return &resp, err
}
