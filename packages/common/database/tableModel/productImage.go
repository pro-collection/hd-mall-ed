package tableModel


// @title           商品图片库
// @description     全都是商品图片
// @auth            晴小篆  331393627@qq.com
// @param
// 	type
//		1 - 商品主图
//		2 - 商品详情图
// @return
type ProductImageBase struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	ProductId uint   `json:"product_id" gorm:"not null;index;comment:'关联的productId'"`
	Type      int    `json:"type" gorm:"not null;comment:'图片类型'"`
	Url       string `json:"url" gorm:"not null;comment:'图片链接';type:longtext"`
}

type ProductImage struct {
	ProductImageBase
	ModelStruct
}
