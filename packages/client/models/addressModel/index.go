package addressModel

import (
	"hd-mall-ed/packages/client/database"
	"hd-mall-ed/packages/client/database/tableModel"
	"hd-mall-ed/packages/client/models/userModel"
)

type Address tableModel.Address

// 通过 id 找地址
func (address *Address) FindAddressById(id uint) (tableModel.BaseAddress, error) {
	resultAddress := tableModel.BaseAddress{}
	// 主见查询
	err := database.DataBase.First(&resultAddress, id).Error
	if err != nil {
		return resultAddress, err
	}
	return resultAddress, nil
}

// 通过 userid 查找关联的 地址列表
func (address *Address) FindAddressByUserId(userId uint) ([]tableModel.BaseAddress, error) {
	var addressList []tableModel.BaseAddress

	err := database.DataBase.Model(&Address{}).Where("user_id = ?", userId).Find(addressList).Error
	if err != nil {
		return addressList, err
	}
	return addressList, nil
}

// 更新默认地址
// 独立的 service 方法
func UpdateDefaultAddressId(userId, addressId int) error {
	err := database.DataBase.Model(&userModel.User{}).Where("id = ?", userId).Update("default_address_id", addressId).Error
	return err
}
