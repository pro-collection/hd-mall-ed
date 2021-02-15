package productCategoryController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/productCategoryModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Update(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &productCategoryModel.ProductCategory{}
	var err error

	err = c.ShouldBindJSON(model)
	if api.ValidateHasError(model) {
		return
	}

	err = model.Update()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
