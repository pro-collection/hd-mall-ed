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

	paginationBase := database.DataBase.
		Model(&Product{}).
		Scopes(modelHelper.Paginate(query.Page, query.PageSize))

	queryMap := make(map[string]interface{})

	if queryBase.CategoryId != 0 {
		queryMap["category_id"] = queryBase.CategoryId
	}

	if queryBase.Max != 0 {
		paginationBase.Where("price < ?", queryBase.Max)
	}

	if queryBase.Min != 0 {
		paginationBase.Where("price > ?", queryBase.Min)
	}

	if queryBase.Query != "" {
		paginationBase.Where("name like ?", "%"+queryBase.Query+"%")
		paginationBase.Where("title like ?", "%"+queryBase.Query+"%")
	}

	//if queryBase.SortType != "" {
	//	switch queryBase.SortType {
	//	case "1":
	//		{
	//
	//			break
	//		}
	//	}
	//}

	err := paginationBase.
		Where(queryMap).
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
