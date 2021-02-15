package productCategoryController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/admin/models/productCategoryModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Delete(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &productCategoryModel.ProductCategory{}
	var err error
	err = c.ShouldBindJSON(model)

	if funk.IsEmpty(model.ID) {
		err = errors.New(e.GetMsg(e.NotFoundId))
	}

	err = model.Delete()

	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
