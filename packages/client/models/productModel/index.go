package productModel

import (
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/common/database"
	"hd-mall-ed/packages/common/database/tableModel"
)

type Product tableModel.Product

func (*Product) GetPrimaryCategoryAndProductList() *[]CategoryProductStruct {
	primaryCategoryList := &[]tableModel.Category{}
	database.DataBase.Model(&tableModel.Category{}).Where("type = ?", 2).Find(primaryCategoryList)
	var list []CategoryProductStruct

	for _, category := range *primaryCategoryList {
		productList := &[]tableModel.ProductBase{}

		// todo 这个地方可以使用 协程
		database.DataBase.Model(&Product{}).Where("category_id = ?", category.ID).Limit(8).Find(productList)

		categoryProductStruct := CategoryProductStruct{}
		_ = deepcopier.Copy(&category).To(&categoryProductStruct)
		categoryProductStruct.ProductList = productList

		list = append(list, categoryProductStruct)
	}

	return &list
}

// 获取限时折扣的商品
func (*Product) GetLimitDiscount() *[]tableModel.ProductBase {
	list := &[]tableModel.ProductBase{}
	database.DataBase.Where("tag = ", 6).Where("status = ", 1).Limit(12).Find(list)
	return list
}
