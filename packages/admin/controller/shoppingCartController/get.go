package shoppingCartController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/shoppingCartModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func GetList(c *gin.Context) {
	api := adminApp.ApiInit(c)

	// 必须要 type 参数
	model := &shoppingCartModel.ShoppingCart{}

	_ = c.ShouldBindJSON(model)

	if api.ValidateHasError(model) {
		return
	}

	model.UserId = uint(api.GetUserId())
	list, err := model.GetList()
	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	api.Response(list)
}
