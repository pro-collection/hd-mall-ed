package productController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/productModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func GetListByQuery(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &productModel.Product{}
	var err error
	var list []*productModel.Product
	query := &productModel.GetListQueryStruct{}
	err = c.ShouldBindJSON(query)

	if api.ValidateHasError(query) {
		return
	}

	list, err = model.GetListByQuery(query)
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.Response(list)
}
