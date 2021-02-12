package addressController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/addressModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
	"log"
)

// @title     更新地址
// @description
// @auth      晴小篆  331393627@qq.com
// @param     addressModel.UpdateRequestParamsStruct
// @return
func Update(c *gin.Context) {
	api := app.ApiFunction{C: c}
	var err error
	model := addressModel.Address{}

	// 获取请求参数
	updateParams := addressModel.UpdateRequestParamsStruct{}
	err = c.ShouldBindJSON(&updateParams)

	if err != nil {
		log.Println(err)
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

// @title     更细默认地址
// @description
// @auth      晴小篆  331393627@qq.com
// @param     addressModel.UpdateDefaultRequestParamsStruct
// @return
func UpdateDefault(c *gin.Context) {
	api := app.ApiFunction{C: c}
	var err error
	model := addressModel.Address{}

	// 获取数据
	params := addressModel.UpdateDefaultRequestParamsStruct{}
	if err = c.ShouldBindJSON(&params); err != nil {
		log.Println(err)
		api.ResFail(e.InvalidParams)
		return
	}

	// 校验参数
	if api.ValidateHasError(params) {
		return
	}

	user, _ := api.GetUser()
	model.ID = uint(params.Id)

	// 更新默认地址
	if err = model.UpdateDefaultAddressId(int(user.ID)); err != nil {
		api.ResFail(e.AddressSetDefaultFail)
		return
	}
	api.ResponseNoData()
}
