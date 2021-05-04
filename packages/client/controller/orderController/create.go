package orderController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/orderModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
)

func Create(c *gin.Context) {
	api := app.ApiFunction{C: c}
	orderMapper := &orderModel.Order{}

	err := c.ShouldBindJSON(orderMapper)
	orderMapper.UserId = uint(api.GetUserId())
	orderMapper.Status = 1

	if api.ValidateHasError(orderMapper) {
		return
	}

	err = orderMapper.CreateOrder()
	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	api.ResponseNoData()
}
