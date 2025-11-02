package rainyun_go_sdk

import "encoding/json"

// RCS查询参数
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

// 工单查询参数
type WorkerorderQueryParameters struct {
	ColumnFilters struct {
		Title string `json:"Title"`
	} `json:"columnFilters"`
	Sort    []any `json:"sort"`
	Page    int   `json:"page"`
	PerPage int   `json:"perPage"`
}

// 编码工单查询参数
func MarshalWorkerorderQueryParameters(title string, page int, perPage int) string {
	queryParameters := WorkerorderQueryParameters{
		ColumnFilters: struct {
			Title string `json:"Title"`
		}{
			Title: title,
		},
		Sort:    []any{},
		Page:    page,
		PerPage: perPage,
	}

	result, _ := json.Marshal(&queryParameters)

	return string(result)
}
