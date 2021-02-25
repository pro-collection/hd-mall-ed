package productModel

import (
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
		database.DataBase.Model(&Product{}).Where("category_id = ?", category.ID).Limit(6).Find(productList)

		categoryProductStruct := CategoryProductStruct{}
		categoryProductStruct.CategoryId = int(category.ID)
		categoryProductStruct.ProductList = productList

		list = append(list, categoryProductStruct)
	}

	return &list
}
