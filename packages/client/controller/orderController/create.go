package orderController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/orderModel"
	"hd-mall-ed/packages/client/models/shoppingCartModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
	"time"
)

/*
请求参数 tableModel.OrderBase
*/
func Create(c *gin.Context) {
	api := app.ApiFunction{C: c}
	orderMapper := &orderModel.Order{}
	shoppingCartMapper := &shoppingCartModel.ShoppingCart{}

	err := c.ShouldBindJSON(orderMapper)
	shoppingCartMapper.UserId = uint(api.GetUserId())

	orderMapper.UserId = uint(api.GetUserId())
	orderMapper.Status = 1 // 设置状态 - 确认订单

	shoppingCartUpdateMap := make(map[string]interface{})
	shoppingCartUpdateMap["type"] = 3

	go shoppingCartMapper.StandUpdate("temp_order_id = ?", orderMapper.OrderId, &shoppingCartUpdateMap)

	orderMapper.SendTime = time.Unix(0, 0)
	orderMapper.ConfirmTime = time.Unix(0, 0)
	orderMapper.CompleteTime = time.Unix(0, 0)
	err = orderMapper.CreateOrder()
	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	api.Response(&createRes{
		Id:      orderMapper.ID,
		OrderId: orderMapper.OrderId,
	})
}
