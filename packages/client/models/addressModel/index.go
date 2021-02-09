package addressModel

import (
	"errors"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/client/database"
	"hd-mall-ed/packages/client/database/tableModel"
	"hd-mall-ed/packages/client/models/userModel"
)

type Address tableModel.Address

// 通过 id 找地址
func (address *Address) FindAddressById(addressId, userId uint) error {
	// 主见查询
	err := database.DataBase.Where("user_id = ?", userId).First(&address, addressId).Error
	if err != nil {
		return err
	}
	return nil
}

// 通过 userid 查找关联的 地址列表
func (address *Address) FindAddressByUserId(userId uint) (*[]tableModel.AddressBase, error) {
	var addressList []tableModel.AddressBase
	if funk.IsEmpty(userId) {
		return nil, errors.New("user_id 不存在")
	}

	err := database.DataBase.Model(&Address{}).Where("user_id = ?", userId).Find(&addressList).Error
	if err != nil {
		return &addressList, err
	}
	return &addressList, nil
}

// 创建用户收件地址
func (address *Address) CreateAddress() error {
	// 必须要要求有 user_id
	if funk.IsEmpty(address.UserId) {
		return errors.New("user_id 不存在")
	}
	return database.DataBase.Create(address).Error
}

// 更新默认地址
// 独立的 service 方法
func UpdateDefaultAddressId(userId, addressId int) error {
	err := database.DataBase.Model(&userModel.User{}).Where("id = ?", userId).Update("default_address_id", addressId).Error
	return err
}
