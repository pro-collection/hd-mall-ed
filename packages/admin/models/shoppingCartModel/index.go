package shoppingCartModel

import (
	"hd-mall-ed/packages/common/database"
	"hd-mall-ed/packages/common/database/tableModel"
)

type ShoppingCart tableModel.ShoppingCart

// 插入多个
func (*ShoppingCart) CreateShoppingCartInfo(infoList *[]ShoppingCart) error {
	return database.DataBase.Model(&ShoppingCart{}).Create(infoList).Error
}

// 获取所有的
func (shoppingCart *ShoppingCart) GetList() (*[]ShoppingCart, error) {
	list := &[]ShoppingCart{}
	err := database.DataBase.Model(&ShoppingCart{}).
		Where("user_id = ?", shoppingCart.UserId).
		Where("type = ?", shoppingCart.Type).
		Find(list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

// 删除
func (shoppingCart *ShoppingCart) Delete() error {
	return database.DataBase.Delete(shoppingCart).Error
}

// 更新
func (shoppingCart *ShoppingCart) Update() error {
	return database.DataBase.Where("id = ?", shoppingCart.ID).
		Where("user_id", shoppingCart.UserId).
		Update("count", shoppingCart.Count).Error
}
