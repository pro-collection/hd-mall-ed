package addressModel

import (
	"hd-mall-ed/packages/client/database"
	"hd-mall-ed/packages/client/database/tableModel"
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
