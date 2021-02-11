package tableModel

type CategoryBase struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Name     string `json:"name" gorm:"not null;comment:'商品分类'" valid:"required~缺少商品类别名称"`
	ParentId uint   `json:"parent_id" gorm:"index;comment:'父节点 id , 可以为空'"`
	Avatar   string `json:"avatar" gorm:"comment:'缩略图''"`
}

type Category struct {
	ModelStruct
	CategoryBase
}
