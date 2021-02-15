package productCategoryController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/admin/models/productCategoryModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"strconv"
)

func GetList(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &productCategoryModel.ProductCategory{}
	idString := c.DefaultQuery("id", "")
	id, err := strconv.Atoi(idString)

	if funk.IsEmpty(id) {
		err = errors.New(e.GetMsg(e.NotFoundId))
	}

	// 查询逻辑
	list, err := model.GetList()

	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.Response(list)
}
