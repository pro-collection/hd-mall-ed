package shoppingCartController

import (
	"hd-mall-ed/packages/admin/models/productModel"
	"hd-mall-ed/packages/common/database/tableModel"
)

type deleteParams struct {
	ProductInfo productModel.Product `json:"product_info"`
	tableModel.ShoppingCart
}
