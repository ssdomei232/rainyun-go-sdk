// TODO: 统一处理标准查询参数
package rainyun_go_sdk

import (
	"encoding/json"
	"strconv"
)

// 标准查询参数
type StandQueryParameters struct {
	ColumnFilters struct{} `json:"columnFilters"`
	Sort          []any    `json:"sort"`
	Page          int      `json:"page"`
	PerPage       int      `json:"perPage"`
}

// 编码标准查询参数
func EncodingStandardQueryParameters(page int, perPage int) string {
	queryParameters := StandQueryParameters{
		ColumnFilters: struct{}{},
		Sort:          []any{},
		Page:          page,
		PerPage:       perPage,
	}

	result, _ := json.Marshal(&queryParameters)

	return string(result)
}

// 获取VNC连接URL
func GetVncConnectURL(v *VncConnectionInfo) (string, error) {
	decoded, err := strconv.Unquote(v.Data.VNCProxyURL)
	if err != nil {
		return "", err
	}

	return decoded, nil
}
