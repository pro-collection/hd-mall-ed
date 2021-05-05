package orderController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/orderModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
	"time"
)

func Update(c *gin.Context) {
	api := app.ApiFunction{C: c}

	orderMapper := &orderModel.Order{}
	orderMapper.UserId = uint(api.GetUserId())

	err := c.ShouldBindJSON(orderMapper)

	queryMap := make(map[string]interface{})

	if orderMapper.Status > 0 && orderMapper.Status <= 4 {
		queryMap["status"] = orderMapper.Status
		if orderMapper.Status == 2 {
			// 发货
			queryMap["send_time"] = time.Now()
		} else if orderMapper.Status == 3 {
			// 确认收货
			queryMap["confirm_time"] = time.Now()
		} else if orderMapper.Status == 4 {
			queryMap["complete_time"] = time.Now()
		}
	} else {
		// 缺少订单状态
		api.ResFail(e.Fail)
		return
	}

	// 更新
	err = orderMapper.UpdateStatus(&queryMap)
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
