package productModel

import (
	"hd-mall-ed/packages/common/database"
	"hd-mall-ed/packages/common/database/tableModel"
	"hd-mall-ed/packages/common/models/modelHelper"
	"strconv"
)

type Product tableModel.Product

// 创建
func (product *Product) Create() error {
	return database.DataBase.Create(product).Error
}

// 条件查询商品信息
func (product *Product) GetListByQuery(query map[string]string) ([]*Product, error) {
	var productList []*Product

	page, err := strconv.Atoi(query["page"])
	pageSize, err := strconv.Atoi(query["pageSize"])

	err = database.DataBase.
		Model(&Product{}).
		Scopes(modelHelper.Paginate(page, pageSize)).
		Where(query).
		Find(&productList).Error
	if err != nil {
		return productList, err
	}
	return productList, nil
}


