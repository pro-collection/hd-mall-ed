package bannerController

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/admin/models/staticModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Update(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &staticModel.Static{}

	// 参数绑定
	err := c.ShouldBindJSON(model)
	if funk.IsEmpty(model.ID) {
		api.ResFail(e.NotFoundId)
		return
	}
	if api.ValidateHasError(model) {
		api.ResFail(e.InvalidParams)
		return
	}

	// 更新
	err = model.Update()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}
	api.ResponseNoData()
}
