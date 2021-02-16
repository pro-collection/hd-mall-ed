package productController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/admin/models/productModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Delete(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &productModel.Product{}
	var err error

	params := &deleteParamsStruct{}

	err = c.ShouldBindJSON(params)

	if funk.IsEmpty(params.Id) {
		err = errors.New(e.GetMsg(e.NotFoundId))
		api.ResFail(e.NotFoundId)
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
