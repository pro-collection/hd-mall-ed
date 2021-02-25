package productModel

import "hd-mall-ed/packages/common/database/tableModel"

type CategoryProductStruct struct {
	CategoryId int `json:"category_id"`
	ProductList *[]tableModel.ProductBase `json:"product_list"`
}
