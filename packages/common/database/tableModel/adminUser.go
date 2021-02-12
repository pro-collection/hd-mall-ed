package tableModel

// @title           基础 admin user
// @description
// @auth            晴小篆  331393627@qq.com
// @param
// 	role
// 		1 - 超管用户
// 		2 - 普通用户
// @return
type AdminUserBase struct {
	ID     uint   `json:"id" gorm:"primarykey"`
	Name   string `json:"name" gorm:"not null"`
	Avatar string `json:"avatar" gorm:"comment:头像"`
	Email  string `json:"email" gorm:"comment:邮箱"`
	Role   string `json:"role" gorm:"comment:角色"`
}

type AdminUser struct {
	ModelStruct
	Password string `json:"password" gorm:"not null"`
	AdminUserBase
}
