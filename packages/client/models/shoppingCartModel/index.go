package shoppingCartModel

import (
	"hd-mall-ed/packages/common/database"
	"hd-mall-ed/packages/common/database/tableModel"
)

type ShoppingCart tableModel.ShoppingCart

// 插入多个
func (shoppingCart *ShoppingCart) CreateShoppingCartInfo() error {
	return database.DataBase.Create(shoppingCart).Error
}

// 通过 id 拿到信息
func (shoppingCart *ShoppingCart) GetDetailById() error {
	return database.DataBase.Where("id = ?", shoppingCart.ID).First(shoppingCart).Error
}

func (shoppingCart *ShoppingCart) GetDetailByQuery(query *map[string]interface{}) (*ShoppingCart, error) {
	cart := &ShoppingCart{}
	err := database.DataBase.Where(*query).First(cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

// 通过 id 拿到信息
func (shoppingCart *ShoppingCart) GetListByTempOrderId() (*[]ShoppingCart, error) {
	cartList := &[]ShoppingCart{}
	err := database.DataBase.Where("temp_order_id = ?", shoppingCart.TempOrderId).Find(cartList).Error
	if err != nil {
		return cartList, err
	}
	return cartList, nil
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
func (shoppingCart *ShoppingCart) Delete(idList *[]uint) error {
	return database.DataBase.Unscoped().Delete(&ShoppingCart{}, idList).Error
}

// 更新
func (shoppingCart *ShoppingCart) Update() error {
	return database.DataBase.Model(&ShoppingCart{}).Where("id = ?", shoppingCart.ID).
		Where("user_id", shoppingCart.UserId).
		Update("count", shoppingCart.Count).Error
}

// 批量更新
func (shoppingCart *ShoppingCart) Updates(queryMap *[]uint, updateMap *map[string]interface{}) error {
	return database.DataBase.Model(&ShoppingCart{}).
		Where("user_id", shoppingCart.UserId).
		Where("id in ?", *queryMap).
		Updates(*updateMap).Error
}

// 可扩展更新方法
func (shoppingCart *ShoppingCart) StandUpdate(queryString string, queryMap interface{}, updateMap *map[string]interface{}) error {
	return database.DataBase.Model(&ShoppingCart{}).
		Where("user_id", shoppingCart.UserId).
		Where(queryString, queryMap).
		Updates(updateMap).Error
}

// 可扩展更新方法 - 已经限定了 order_id
func (shoppingCart *ShoppingCart) StandUpdateDefault(queryString string, queryMap interface{}, updateMap *map[string]interface{}) error {
	return database.DataBase.Model(&ShoppingCart{}).
		Where("user_id", shoppingCart.UserId).
		Where("temp_order_id = ?", shoppingCart.TempOrderId).
		Where(queryString, queryMap).
		Updates(updateMap).Error
}
