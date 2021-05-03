package tableModel


/*
type 类型
	1 - 购物车
	2 - pre confirm order
*/
type ShoppingCartBase struct {
	ID        uint `json:"id" gorm:"primarykey"`
	ProductId uint `json:"product_id" gorm:"index;comment:关联的productId"`
	Type int `json:"type" gorm:"not null;comment:图片类型" valid:"required"`
	Count int `json:"count" gorm:"not null;comment:数量;default:1" valid:"required"`
}

type ShoppingCart struct {
	ModelStruct
}
