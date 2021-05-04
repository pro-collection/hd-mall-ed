package orderController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/orderModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
	"strconv"
)

func GetDetail(c *gin.Context) {
	api := app.ApiFunction{C: c}
	orderMapper := &orderModel.Order{}

	id := c.DefaultQuery("id", "")
	orderId := c.DefaultQuery("order_id", "")

	queryMapper := make(map[string]interface{})

	// 绑定 user 信息
	queryMapper["user_id"] = uint(api.GetUserId())

	if id != "" {
		idNumber, _ := strconv.Atoi(id)
		queryMapper["id"] = idNumber
	}

	if orderId != "" {
		orderIdNumber, _ := strconv.Atoi(orderId)
		queryMapper["order_id"] = orderIdNumber
	}

	if id == "" || orderId == "" {
		api.ResFail(e.Fail)
		return
	}

	// 详情
	err := orderMapper.GetOrderDetailByQuery(&queryMapper)
	if err != nil {
		api.ResFail(e.Fail)
		return
	}
	api.Response(queryMapper)
}

func GetList(c *gin.Context) {
	api := app.ApiFunction{C: c}
	orderMapper := &orderModel.Order{}

	queryMap := make(map[string]interface{})

	status := c.DefaultQuery("status", "")
	if status != "" {

	}

	queryMap["user_id"] = api.GetUserId()

	list, err := orderMapper.GetOrderListByQuery(&queryMap)
	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	api.Response(list)
}
