package addressController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/addressModel"
	"hd-mall-ed/packages/client/pkg/app"
	"hd-mall-ed/packages/client/pkg/e"
)

func Update(c *gin.Context) {
	api := app.ApiFunction{C: c}
	var err error
	model := addressModel.Address{}

	// 获取请求参数
	updateParams := addressModel.UpdateRequestParamsStruct{}
	err = c.ShouldBindJSON(&updateParams)

	if err != nil {
		api.ResFail(e.InvalidParams)
		return
	}

	// 校验
	if api.ValidateHasError(updateParams) {
		return
	}

	// 更新数据
	err = model.UpdateAddress(updateParams)
	if err != nil {
		api.ResFail(e.AddressUpdateFail)
		return
	}

	api.ResponseNoData()
}
