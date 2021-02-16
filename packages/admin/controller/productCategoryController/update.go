package productCategoryController

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/admin/models/productCategoryModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Update(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &productCategoryModel.ProductCategory{}
	var err error

	err = c.ShouldBindJSON(model)
	if funk.IsEmpty(model.ID) {
		api.ResFail(e.NotFoundId)
		return
	}

	if api.ValidateHasError(model) {
		return
	}

	// todo 更新之前校验是是否有重复商品售卖类型
	err = model.Update()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
