package productController

type getListQueryStruct struct {
	Page int `json:"page" valid:"required~缺少page"`
	PageSize int `json:"page_size" valid:"required~缺少page_size"`
	CategoryId int `json:"category_id" valid:"required"`
}
