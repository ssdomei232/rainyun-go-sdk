// TODO: 统一处理标准查询参数
package rainyun_go_sdk

import (
	"encoding/json"
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
