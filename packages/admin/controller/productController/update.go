package productController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/admin/models/productModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Update(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &productModel.Product{}
	var err error

	// 参数绑定与校验
	err = c.ShouldBindJSON(model)
	if api.ValidateHasError(model) {
		return
	}
	if funk.IsEmpty(model.ID) {
		err = errors.New(e.GetMsg(e.NotFoundId))
	}

	// todo 更新之前校验是是否有重复商品
	// 更新接口
	err = model.Update()

	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
