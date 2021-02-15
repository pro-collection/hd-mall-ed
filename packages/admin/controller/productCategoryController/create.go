package productCategoryController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/productCategoryModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Create(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &productCategoryModel.ProductCategory{}
	var err error

	// 绑定参数
	err = c.ShouldBindJSON(model)
	if api.ValidateHasError(model) {
		return
	}

	// todo 创建之前校验是是否有重复商品售卖类型
	err = model.Create()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
