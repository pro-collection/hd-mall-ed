package staticController

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/admin/models/staticModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

// @title           更新文件
// @description
// @auth            晴小篆  331393627@qq.com
// @param
// @return
func Update(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &staticModel.Static{}

	err := c.ShouldBindJSON(model)
	if api.ValidateHasError(model) {
		return
	}

	if funk.IsEmpty(model.ID) {
		api.ResFail(e.NotFoundId)
		return
	}

	err = model.Update()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
