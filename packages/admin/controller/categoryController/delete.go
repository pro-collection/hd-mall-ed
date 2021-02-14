package categoryController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/categoryModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Delete(c *gin.Context) {
	api := adminApp.ApiInit(c)
	var err error
	model := &categoryModel.Category{}

	params := &deleteParamsStruct{}
	err = c.ShouldBindJSON(params)

	if params.Id <= 0 {
		api.ResFail(e.InvalidParams)
		return
	}

	model.ID = uint(params.Id)
	err = model.Delete()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
