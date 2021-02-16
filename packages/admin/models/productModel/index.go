package productModel

import (
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/common/database"
	"hd-mall-ed/packages/common/database/tableModel"
	"hd-mall-ed/packages/common/models/modelHelper"
)

type Product tableModel.Product

// 创建
func (product *Product) Create() error {
	return database.DataBase.Create(product).Error
}

// 条件查询商品信息
func (product *Product) GetListByQuery(query *GetListQueryStruct) ([]*Product, error) {
	var productList []*Product

	queryBase := &GetListQueryBaseStruct{}
	_ = deepcopier.Copy(query).To(queryBase)

	err := database.DataBase.
		Model(&Product{}).
		Scopes(modelHelper.Paginate(query.Page, query.PageSize)).
		Where(queryBase).
		Find(&productList).Error
	if err != nil {
		return productList, err
	}
	return productList, nil
}

func (product *Product) Update() error {
	return database.DataBase.Where("id = ?", product.ID).Updates(*product).Error
}

func (product *Product) Delete() error {
	return database.DataBase.Delete(product).Error
}
