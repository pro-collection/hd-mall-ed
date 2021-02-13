package userController

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/admin/models/adminUserModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

// @title           更新
// @description
// @auth            晴小篆  331393627@qq.com
// @param           userController.updateParamsStruct
// @return
func Update(c *gin.Context) {
	api := adminApp.ApiInit(c)
	params := &updateParamsStruct{}
	model := &adminUserModel.AdminUser{}
	var err error

	err = c.ShouldBindJSON(params)
	if api.ValidateHasError(params) {
		return
	}

	err = deepcopier.Copy(params).To(model)
	err = model.Update()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
