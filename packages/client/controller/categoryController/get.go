package categoryController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/categoryModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
)

func Get(c *gin.Context) {
	api := &app.ApiFunction{C: c}
	model := categoryModel.Category{}

	list, err := model.Get()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.Response(handleListHelper(list))
}
