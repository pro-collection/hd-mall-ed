package categoryModel

import (
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/client/database"
	"hd-mall-ed/packages/client/database/tableModel"
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
	baseModel := tableModel.CategoryBase{}
	_ = deepcopier.Copy(category).To(&baseModel)
	return database.DataBase.Updates(baseModel).Error
}

// 删除
func (category *Category) Delete() error {
	return database.DataBase.Delete(category).Error
}
