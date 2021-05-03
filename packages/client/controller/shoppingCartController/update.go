package shoppingCartController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/shoppingCartModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
	"time"
)

/*
参数 tableModel.ShoppingCartBase
*/
func Update(c *gin.Context) {
	api := app.ApiFunction{C: c}

	model := &shoppingCartModel.ShoppingCart{}

	_ = c.ShouldBindJSON(model)

	model.UserId = uint(api.GetUserId())

	err := model.Update()
	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	api.ResponseNoData()
}

func SettleAccounts(c *gin.Context) {
	api := app.ApiFunction{C: c}
	model := &shoppingCartModel.ShoppingCart{}

	model.UserId = uint(api.GetUserId())

	var list []shoppingCartModel.ShoppingCart

	_ = c.ShouldBind(&list)

	tempOrderId := uint(time.Now().UnixNano() / 1e6)

	var queryMap []uint
	updateMap := make(map[string]interface{})

	for _, cart := range list {
		//updateMap["type"] = cart.Type
		updateMap["temp_order_id"] = tempOrderId
		queryMap = append(queryMap, cart.ID)
	}

	err := model.Updates(&queryMap, &updateMap)
	if err != nil {
		api.ResFail(e.Fail)
		return
	}
	api.Response(tempOrderId)
}
