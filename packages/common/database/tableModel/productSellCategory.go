package tableModel

type ProductCategoryBase struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	ProductId uint   `json:"product_id" gorm:"not null;index;comment:'关联的productId'" valid:"reqruied"`
	Title     string `json:"title" gorm:"comment:'类别名称'" valid:"required"`
	Image     string `json:"image" gorm:"type:longtext;comment:'类别展示图片';"`
}

type ProductCategory struct {
	ProductCategoryBase
	ModelStruct
}
