package rainyun_go_sdk

import "encoding/json"

type RCSQueryParameters struct {
	ColumnFilters struct {
		RcsID string `json:"rcs.ID"`
	} `json:"columnFilters"`
	Sort    []interface{} `json:"sort"`
	Page    int           `json:"page"`
	PerPage int           `json:"perPage"`
}

// 编码RCS查询参数
func MarshalRCSQueryParameters(id int, page int, perPage int) string {
	queryParameters := RCSQueryParameters{
		ColumnFilters: struct {
			RcsID string `json:"rcs.ID"`
		}{
			RcsID: string(rune(id)),
		},
		Sort:    []interface{}{},
		Page:    page,
		PerPage: perPage,
	}

	result, _ := json.Marshal(&queryParameters)

	return string(result)
}
