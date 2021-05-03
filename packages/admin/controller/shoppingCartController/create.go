package shoppingCartController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/shoppingCartModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Create(c *gin.Context) {
	api := adminApp.ApiInit(c)

	model := &shoppingCartModel.ShoppingCart{}

	listParams := &[]shoppingCartModel.ShoppingCart{}

	_ = c.ShouldBindJSON(model)

	err := model.CreateShoppingCartInfo(listParams)
	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	api.ResponseNoData()
}
