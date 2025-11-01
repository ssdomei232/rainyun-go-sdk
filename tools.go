package rainyun_go_sdk

import "encoding/json"

type StandardQueryParameters struct {
	ColumnFilters struct {
	} `json:"columnFilters"`
	Sort    []StandardQueryParametersSort `json:"sort"`
	Page    int                           `json:"page"`    // 页码
	PerPage int                           `json:"perPage"` // 每页显示的行数
}

type StandardQueryParametersSort struct {
	Field string `json:"field"`
	Type  string `json:"type"`
}

// 编码标准查询参数,sort为可选参数,不需要排序就直接传nil
// sort 示例: [{"field":"Time","type":"asc"}]
func MarshalStandardQueryParameters(id int, sort []StandardQueryParametersSort, page int, perPage int, start int, end int) string {
	if sort == nil {
		sort = []StandardQueryParametersSort{}
	}

	var queryParameters StandardQueryParameters = StandardQueryParameters{
		ColumnFilters: struct {
		}{},
		Sort:    sort,
		Page:    page,
		PerPage: perPage,
	}

	result, _ := json.Marshal(&queryParameters)

	return string(result)
}
