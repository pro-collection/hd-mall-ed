package productModel

import "hd-mall-ed/packages/common/database/tableModel"

type CategoryProductStruct struct {
	tableModel.CategoryBase
	ProductList *[]tableModel.ProductBase `json:"product_list"`
}
