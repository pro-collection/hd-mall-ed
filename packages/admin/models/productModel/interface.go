package productModel

type GetListQueryBaseStruct struct {
	CategoryId int `json:"category_id"`
	Max int64 `json:"max"`
	Min int64 `json:"min"`
	Query string `json:"query"`
	SortType string `json:"sort_type"`
	Type string `json:"type"`
}

type GetListQueryStruct struct {
	Page     int `json:"page" valid:"required~缺少page"`
	PageSize int `json:"page_size" valid:"required~缺少page_size"`
	GetListQueryBaseStruct
}
