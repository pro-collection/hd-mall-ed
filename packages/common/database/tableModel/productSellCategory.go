package tableModel

type ProductSellCategoryBase struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	ProductId uint   `json:"product_id" gorm:"not null;index;comment:'关联的productId'"`
	Title     string `json:"title" gorm:"comment:'类别名称'"`
	Image     string `json:"image" gorm:"type:longtext;comment:'类别展示图片';"`
}

type ProductSellCategory struct {
	ProductSellCategoryBase
	ModelStruct
}
