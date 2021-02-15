package productController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/productModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Create(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &productModel.Product{}
	var err error

	// 参数绑定
	err = c.ShouldBindJSON(model)
	if api.ValidateHasError(model) {
		return
	}

	err = model.Create()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
