package productController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/productModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"hd-mall-ed/packages/common/pkg/utils"
)

func GetListByQuery(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &productModel.Product{}
	var err error
	var list []*productModel.Product
	query := &getListQueryStruct{}
	err = c.ShouldBindJSON(query)
	queryMap := utils.Struct2Map(query)

	list, err = model.GetListByQuery(queryMap)
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.Response(list)
}
