package productModel

type GetListQueryBaseStruct struct {
	CategoryId int `json:"category_id"`
}

type GetListQueryStruct struct {
	Page     int `json:"page" valid:"required~缺少page"`
	PageSize int `json:"page_size" valid:"required~缺少page_size"`
	GetListQueryBaseStruct
}
