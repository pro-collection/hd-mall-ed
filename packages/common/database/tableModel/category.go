package tableModel

// @title
// @description
// @auth            晴小篆  331393627@qq.com
// @param
// 	type:
//		1 - 没有类型设置
//		2 - primary
// @return
type CategoryBase struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Name     string `json:"name" gorm:"not null;comment:'商品分类'" valid:"required~缺少商品类别名称"`
	Type     int    `json:"type" gorm:"comment:'商品分类的类型'"`
	ParentId uint   `json:"parent_id" gorm:"index;comment:'父节点 id , 可以为空'"`
	Avatar   string `json:"avatar" gorm:"comment:缩略图"`
	BannerImage string `json:"banner_image" gorm:"content:banner图片;type:longtext"`
	BannerProductId uint `json:"banner_product_id" gorm:"content:banner对应的商品id, 可以实现跳转"`
	BannerLink string `json:"banner_link" gorm:"content:banner对应商品的跳转链接"`
}

type Category struct {
	ModelStruct
	CategoryBase
}
