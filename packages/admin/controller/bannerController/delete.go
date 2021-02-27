package bannerController

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/admin/models/staticModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Delete(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &staticModel.Static{}

	err := c.ShouldBindJSON(model)
	if funk.IsEmpty(model.ID) {
		api.ResFail(e.InvalidParams)
		return
	}

	// 删除
	err = model.Delete()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
