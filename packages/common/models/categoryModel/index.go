package categoryModel

import (
	"hd-mall-ed/packages/common/database"
	"hd-mall-ed/packages/common/database/tableModel"
)

type Category tableModel.Category

// 创建
func (category *Category) Create() error {
	return database.DataBase.Create(category).Error
}

// 查询
func (category *Category) Get() ([]*tableModel.CategoryBase, error) {
	var categoryList []*tableModel.CategoryBase
	if err := database.DataBase.Model(&Category{}).Find(categoryList).Error; err != nil {
		return categoryList, err
	}
	return categoryList, nil
}

// 更新
func (category *Category) Update() error {
	return database.DataBase.Model(&Category{}).Updates(&category).Error
}

// 删除
func (category *Category) Delete() error {
	return database.DataBase.Delete(category).Error
}
