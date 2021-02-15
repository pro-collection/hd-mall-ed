package productModel

import (
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
func (product *Product) GetListByQuery(query map[string]interface{}) ([]*Product, error) {
	var productList []*Product

	page := query["page"].(int)
	pageSize := query["pageSize"].(int)

	err := database.DataBase.
		Model(&Product{}).
		Scopes(modelHelper.Paginate(page, pageSize)).
		Where(query).
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
