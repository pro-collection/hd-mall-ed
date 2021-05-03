package shoppingCartController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/shoppingCartModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"strconv"
)

// 参数 type
func GetList(c *gin.Context) {
	api := adminApp.ApiInit(c)

	// 必须要 type 参数
	model := &shoppingCartModel.ShoppingCart{}

	typeString := c.Query("type")
	typeNumber, _ := strconv.Atoi(typeString)
	model.Type = typeNumber

	model.UserId = uint(api.GetUserId())
	list, err := model.GetList()
	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	api.Response(list)
}
