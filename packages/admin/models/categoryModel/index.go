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

// 获取所有的列表数据
func (category *Category) Get() ([]*tableModel.CategoryBase, error) {
	var categoryList []*tableModel.CategoryBase
	if err := database.DataBase.Model(&Category{}).Find(&categoryList).Error; err != nil {
		return categoryList, err
	}
	return categoryList, nil
}

// 查询 primary 分类的个数
func (category *Category) GetPrimaryCount() (int, error) {
	var count int64
	err := database.DataBase.Model(&Category{}).Where("parent_id != ?", 0).Count(&count).Error
	if err != nil {
		return int(count), err
	}
	return int(count), nil
}

// 通过类别名称查找是否重复
func (category *Category) FindByName() (*Category, error) {
	findCategory := &Category{}
	if err := database.DataBase.Model(&Category{}).Where("name = ?", category.Name).First(findCategory).Error; err != nil {
		return findCategory, err
	}
	return findCategory, nil
}

// 更新
func (category *Category) Update() error {
	return database.DataBase.Model(&Category{}).Where("id = ?", category.ID).Updates(&category).Error
}

// 删除
func (category *Category) Delete() error {
	return database.DataBase.Delete(category).Error
}
