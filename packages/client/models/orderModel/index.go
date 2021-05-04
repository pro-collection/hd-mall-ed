package orderModel

import (
	"hd-mall-ed/packages/common/database"
	"hd-mall-ed/packages/common/database/tableModel"
)

type Order tableModel.Order

func (order *Order) CreateOrder() error {
	return database.DataBase.Create(order).Error
}

func (order *Order) GetOrderDetailByQuery(query *map[string]interface{}) error {
	return database.DataBase.Where(*query).First(order).Error
}

// 获取列表信息
func (order *Order) GetOrderListByQuery(query *map[string]interface{}) (*[]Order, error) {
	list := &[]Order{}

	err := database.DataBase.Where(query).Find(list).Error

	if err != nil {
		return list, err
	}
	return list, nil
}

// 删除
func (order *Order) DeleteOrder() error {
	return database.DataBase.Delete(order).Error
}

// 更新
func (order *Order) Update() error {
	return database.DataBase.Updates(order).Error
}
