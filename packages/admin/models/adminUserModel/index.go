package adminUserModel

import (
	"hd-mall-ed/packages/common/database"
	"hd-mall-ed/packages/common/database/tableModel"
)

type AdminUser tableModel.AdminUser

// 创建用户
func (user *AdminUser) CreateUser() (err error) {
	err = database.DataBase.Create(&user).Error
	return
}

// 查找user
func (user *AdminUser) FindUser() AdminUser {
	database.DataBase.First(user)
	return *user
}

// 通过 id 查找user
func (*AdminUser) FindUserById(id uint) (AdminUser, error) {
	resultUser := AdminUser{}
	err := database.DataBase.First(&resultUser, id).Error
	if err != nil {
		return resultUser, err
	}
	return resultUser, nil
}

// 通过用户名查找
func (user *AdminUser) FindUserByName(name string) (AdminUser, error) {
	resultUser := AdminUser{}
	err := database.DataBase.Where("name = ?", name).Not("id = ?", user.ID).First(&resultUser).Error
	if err != nil {
		return resultUser, err
	}
	return resultUser, nil
}

// 更新方法
func (user *AdminUser) Update() error {
	err := database.DataBase.Model(&AdminUser{}).Where("id = ?", user.ID).Updates(&user).Error
	return err
}
