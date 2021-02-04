package tableModal

import "gorm.io/gorm"

type BaseUser struct {
	Name             string `gorm:"not null"`
	Integral         uint   `gorm:"comment:积分"`
	DefaultAddressId uint   `gorm:"comment:默认地址"`
	Avatar           string `gorm:"comment:头像"`
	Email            string `gorm:"comment:邮箱"`
	Role             string `gorm:"comment:角色"`
}

type User struct {
	gorm.Model
	Password string `gorm:"not null"`
}
