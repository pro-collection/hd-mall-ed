package categoryController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/categoryModel"
	"hd-mall-ed/packages/common/controller/categoryHelper"
	"hd-mall-ed/packages/common/database/tableModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func GetList(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &categoryModel.Category{}
	var list []*tableModel.CategoryBase
	var err error

	list, err = model.Get()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.Response(categoryHelper.HandleListHelper(list))
}
