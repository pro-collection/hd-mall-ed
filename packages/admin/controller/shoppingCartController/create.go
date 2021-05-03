package shoppingCartController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/shoppingCartModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"time"
)

/*
数组形式的参数 tableModel.ShoppingCartBase
*/
func Create(c *gin.Context) {
	api := adminApp.ApiInit(c)

	model := &shoppingCartModel.ShoppingCart{}

	_ = c.ShouldBind(&model)

	model.UserId = uint(api.GetUserId())

	if model.Type == 2 {
		model.TempOrderId = uint(time.Now().UnixNano() / 1e6)
	}

	err := model.CreateShoppingCartInfo()
	if err != nil {
		api.ResFail(e.Fail)
		return
	}


	if model.Type == 2 {
		api.Response(model.TempOrderId)
		return
	}

	api.ResponseNoData()
}
