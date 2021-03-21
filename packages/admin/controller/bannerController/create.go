package bannerController

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/admin/models/staticModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

// 单个创建
func Create(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &staticModel.Static{}
	err := c.ShouldBindJSON(model)

	// 验证参数
	if funk.IsEmpty(model.Link) || funk.IsEmpty(model.Url) {
		api.ResFail(e.InvalidParams)
		return
	}

	// 创建入库
	err = model.Create()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
