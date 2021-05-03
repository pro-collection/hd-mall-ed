package shoppingCartController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/shoppingCartModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Delete(c *gin.Context) {
	api := adminApp.ApiInit(c)

	model := &shoppingCartModel.ShoppingCart{}

	_ = c.ShouldBindJSON(model)

	err := model.Delete()
	if err != nil {
		api.ResFail(e.Fail)
		return
	}
	api.ResponseNoData()
}
