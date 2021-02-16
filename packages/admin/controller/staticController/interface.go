package staticController

import "hd-mall-ed/packages/admin/models/staticModel"

type createStaticsParams struct {
	List []staticModel.Static `json:"list" valid:"required"`
}

type GetListByQueryParams struct {
	ID        uint   `json:"id"`
	ProductId uint   `json:"product_id"`
	Type      int    `json:"type"`
	FileName  string `json:"file_name"`
	Url       string `json:"url"`
	Link      string `json:"link"`
}
