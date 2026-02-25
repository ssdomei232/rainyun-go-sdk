package public

import (
	"github.com/ssdomei232/rainyun-go-sdk/v2/rainyun/common"
)

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
	client := common.NewClient("")

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
