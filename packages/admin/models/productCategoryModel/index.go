package productCategoryModel

import (
	"hd-mall-ed/packages/common/database"
	"hd-mall-ed/packages/common/database/tableModel"
)

// 售卖的商品分类
type ProductCategory tableModel.ProductCategory

func (category *ProductCategory) GetList() (*[]ProductCategory, error) {
	list := &[]ProductCategory{}

	err := database.DataBase.Model(&ProductCategory{}).Where("product_id = ?", category.ProductId).Find(&list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (category *ProductCategory) Create() error {
	return database.DataBase.Create(category).Error
}

func (category *ProductCategory) Update() error {
	return database.DataBase.Model(&ProductCategory{}).Updates(*category).Error
}

func (category *ProductCategory) Delete() error {
	return database.DataBase.Model(&ProductCategory{}).Delete(category).Error
}
