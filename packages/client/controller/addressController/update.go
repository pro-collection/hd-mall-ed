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
		api.ResFail(e.FailBindJson)
		return
	}

	// 校验
	if api.ValidateHasError(updateParams) {
		return
	}

	// 更新的业务逻辑
	if handleUpdateAddressHelper(handleUpdateAddressHelperOptions{&api, &model, &updateParams}) {
		return
	}

	api.ResponseNoData()
}

func UpdateDefault(c *gin.Context) {
	api := app.ApiFunction{C: c}
	var err error
	model := addressModel.Address{}

	// 获取数据
	params := addressModel.UpdateDefaultRequestParamsStruct{}
	if err = c.ShouldBindJSON(params); err != nil {
		api.ResFail(e.InvalidParams)
		return
	}

	// 校验参数
	if api.ValidateHasError(params) {
		return
	}

	user, _ := api.GetUser()

	// 更新默认地址
	if err = model.UpdateDefaultAddressId(int(user.ID)); err != nil {
		api.ResFail(e.AddressSetDefaultFail)
		return
	}
	api.ResponseNoData()
}
