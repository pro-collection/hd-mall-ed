package staticController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/staticModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

// @title           单个创建
// @description
// @auth            晴小篆  331393627@qq.com
// @param           staticModel.Static
// @return
func Create(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &staticModel.Static{}
	err := c.ShouldBindJSON(model)

	err = model.Create()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}

// @title           批量创建
// @description
// @auth            晴小篆  331393627@qq.com
// @param           createStaticsParams
// @return
func CreateStatics(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &staticModel.Static{}

	query := &createStaticsParams{}
	err := c.ShouldBindJSON(query)
	if api.ValidateHasError(query) {
		return
	}

	err = model.CreateStatics(&query.List)
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
