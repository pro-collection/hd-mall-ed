package tableModel

// @title           图片等静态资源
// @description
// @auth            晴小篆  331393627@qq.com
// @param
//	type int
//		1 - 商品主图
//		2 - 商品详情图
// 		3 - banner 主图
//		4 - banner 附图
// @return
type StaticBase struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	ProductId uint   `json:"product_id" gorm:"index;comment:关联的productId"`
	Type      int    `json:"type" gorm:"not null;comment:图片类型" valid:"required"`
	FileName  string `json:"file_name" gorm:"comment:文件名字" valid:"required"`
	Url       string `json:"url" gorm:"not null;comment:图片链接;type:longtext" valid:"required"`
	Link      string `json:"link" gorm:"comment:图片的点击跳转链接"`
}

type Static struct {
	StaticBase
	ModelStruct
}
