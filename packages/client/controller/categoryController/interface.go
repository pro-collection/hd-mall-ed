package categoryController

import "hd-mall-ed/packages/common/database/tableModel"

type listItemResultStruct struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Name     string `json:"name" gorm:"not null;comment:'商品分类'" valid:"required~缺少商品类别名称"`
	Type     int    `json:"type" gorm:"comment:'商品分类的类型'"`
	Avatar   string `json:"avatar" gorm:"comment:'缩略图''"`
	Children []*tableModel.CategoryBase
}
