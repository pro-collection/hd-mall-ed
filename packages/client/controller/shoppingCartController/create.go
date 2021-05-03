package shoppingCartController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/shoppingCartModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
	"time"
)

/*
数组形式的参数 tableModel.ShoppingCartBase
*/
func Create(c *gin.Context) {
	api := app.ApiFunction{C: c}

	model := &shoppingCartModel.ShoppingCart{}

	_ = c.ShouldBind(&model)

	model.UserId = uint(api.GetUserId())

	if model.Type == 2 {
		model.TempOrderId = uint(time.Now().UnixNano() / 1e6)
	}

	if model.Type == 1 {
		// 加入购物车场景
		query := make(map[string]interface{})
		query["product_id"] = model.ProductId
		query["type"] = 1
		query["user_id"] = api.GetUserId()
		cart, _ := model.GetDetailByQuery(&query)
		if cart.ID != 0 {
			api.ResFailMessage(e.Fail, "该商品已经加入购物车, 请勿重复添加")
			return
		}
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
