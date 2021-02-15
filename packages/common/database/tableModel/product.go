package tableModel

// @title
// @description
// @auth            晴小篆  331393627@qq.com
// @param
// 	Status
//		1 - 上架
//		2 - 下架
// @return
type ProductBase struct {
	ID           uint    `json:"id" gorm:"primarykey"`
	CategoryId   uint    `json:"category_id" gorm:"index;comment:关联种类id" valid:"required"`
	Name         string  `json:"name" gorm:"not null;comment:商品名称" valid:"required"`
	Title        string  `json:"title" gorm:"not null;comment:title" valid:"required"`
	Desc         string  `json:"desc" gorm:"comment:描述"`
	OriginalCost float32 `json:"original_cost" gorm:"comment:原价" valid:"required"`
	Price        float32 `json:"price" gorm:"not null;comment:实际价格" valid:"required"`
	Inventory    int     `json:"inventory" gorm:"comment:库存" valid:"required"`
	PrimaryImage string  `json:"primary_image" gorm:"comment:主图" valid:"required"`
	Sales        int     `json:"sales" gorm:"comment:销量"`
	Params       string  `json:"params" gorm:"type:longtext;comment:商品参数，json 格式"`
	Tag          string  `json:"tag" gorm:"comment:标签tag"`
	Status       int     `json:"status" gorm:"comment:商品状态"`
}

type Product struct {
	ProductBase
	ModelStruct
}
