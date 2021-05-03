package shoppingCartController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/shoppingCartModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

/*
数组形式的参数 tableModel.ShoppingCartBase
*/
func Create(c *gin.Context) {
	api := adminApp.ApiInit(c)

	model := &shoppingCartModel.ShoppingCart{}

	var listParams []shoppingCartModel.ShoppingCart

	_ = c.ShouldBind(&listParams)

	for i, _ := range listParams {
		listParams[i].UserId = uint(api.GetUserId())
	}

	err := model.CreateShoppingCartInfo(&listParams)
	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	isTempOrder := false
	var tempOrderId uint
	for index, cart := range listParams {
		if index == 0 {
			if cart.Type == 2 {
				isTempOrder = true
				tempOrderId = cart.ID
			}
		}
	}
	if isTempOrder {
		api.Response(tempOrderId)
		return
	}
	api.ResponseNoData()
}
