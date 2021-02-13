package categoryController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/categoryModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Create(c *gin.Context) {
	api := adminApp.ApiInit(c)
	var err error
	model := &categoryModel.Category{}
	err = c.ShouldBindJSON(model)
	if err != nil {
		api.ResFail(e.FailBindJson)
		return
	}

	if api.ValidateHasError(model) {
		return
	}

	// 验证是否重复
	if handleCheckUserNameIsExistHelper(model, api) {
		return
	}

	// 创建
	err = model.Create()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	// 创建成功
	api.ResponseNoData()
}
