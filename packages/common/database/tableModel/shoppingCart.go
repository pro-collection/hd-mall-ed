package tableModel

/*
type 类型
	1 - 购物车
	2 - pre confirm order
	3 - 订单场景
*/
type ShoppingCartBase struct {
	ID           uint    `json:"id" gorm:"primarykey"`
	UserId       uint    `json:"user_id" gorm:"index;comment:关联的UserId"`
	ProductId    uint    `json:"product_id" gorm:"index;comment:关联的productId"`
	TempOrderId  uint    `json:"temp_order_id" gorm:"comment:临时订单id"`
	Type         int     `json:"type" gorm:"not null;comment:类型" valid:"required"`
	Count        int     `json:"count" gorm:"not null;comment:数量;default:1"`
	Name         string  `json:"name" gorm:"not null;comment:商品名称" valid:"required"`
	Title        string  `json:"title" gorm:"not null;comment:title" valid:"required"`
	Desc         string  `json:"desc" gorm:"comment:描述"`
	Price        float32 `json:"price" gorm:"not null;comment:实际价格" valid:"required"`
	PrimaryImage string  `json:"primary_image" gorm:"comment:主图;type:longtext" valid:"required"`
}

type ShoppingCart struct {
	ShoppingCartBase
	ModelStruct
}
