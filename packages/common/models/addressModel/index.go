package addressModel

import (
	"errors"
	"github.com/thoas/go-funk"
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/common/database"
	"hd-mall-ed/packages/common/database/tableModel"
	"hd-mall-ed/packages/common/models/userModel"
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
func (*Address) FindAddressByUserId(userId uint) (*[]tableModel.AddressBase, error) {
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

// 更新用户地址
func (*Address) UpdateAddress(updateParams *UpdateRequestParamsStruct) error {
	addressParams := &Address{}
	_ = deepcopier.Copy(updateParams).To(addressParams)

	// updates 这个方法， 一定要和 model 是一个结构体才行
	return database.DataBase.Model(&Address{}).Where("id = ?", addressParams.ID).Updates(addressParams).Error
}

// 删除用户地址
func (address *Address) DeleteAddress() error {
	return database.DataBase.Delete(address).Error
}

// 更新默认地址
// 独立的 service 方法
func (address *Address) UpdateDefaultAddressId(userId int) error {
	err := database.DataBase.
		Model(&userModel.User{}).
		Where("id = ?", userId).
		Update("default_address_id", address.ID).Error
	return err
}
