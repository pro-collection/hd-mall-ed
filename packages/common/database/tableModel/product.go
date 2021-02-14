package tableModel

type ProductBase struct {
	ID           uint    `json:"id" gorm:"primarykey"`
	Name         string  `json:"name" gorm:"not null;comment:'商品名称'"`
	Title        string  `json:"title" gorm:"not null;comment:'title'"`
	Desc         string  `json:"desc" gorm:"comment:'描述'"`
	OriginalCost float32 `json:"original_cost" gorm:"comment:'原价'"`
	Price        float32 `json:"price" gorm:"not null;comment:'实际价格'"`
	Inventory    int     `json:"inventory" gorm:"comment:'库存'"`
	PrimaryImage string  `json:"primary_image" gorm:"comment:'主图'"`
	Sales        int     `json:"sales" gorm:"comment:'销量'"`
	Params       string  `json:"params" gorm:"type:longtext;comment:商品参数，json 格式"`
	Tag          string  `json:"tag" gorm:"comment:'标签 tag'"`
}

type Product struct {
	ProductBase
	ModelStruct
}
