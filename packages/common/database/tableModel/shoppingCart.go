package tableModel

/*
type 类型
	1 - 购物车
	2 - pre confirm order
	3 - 购物车创建临时订单场景
*/
type ShoppingCartBase struct {
	ID        uint `json:"id" gorm:"primarykey"`
	UserId    uint `json:"user_id" gorm:"index;comment:关联的UserId"`
	ProductId uint `json:"product_id" gorm:"index;comment:关联的productId"`
	Type      int  `json:"type" gorm:"not null;comment:类型" valid:"required"`
	Count     int  `json:"count" gorm:"not null;comment:数量;default:1"`
}

type ShoppingCart struct {
	ShoppingCartBase
	ModelStruct
}
