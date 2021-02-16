package staticController

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/admin/models/staticModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"hd-mall-ed/packages/common/pkg/utils"
)

// @title           获取所有静态文件
// @description
// @auth            晴小篆  331393627@qq.com
// @param
// @return
func GetAllList(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &staticModel.Static{}

	list, err := model.GetAllList()

	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.Response(list)
}

// @title           通过入参条件获取静态文件列表
// @description     staticModel.Static
// @auth            晴小篆  331393627@qq.com
// @param
// @return
func GetListByQuery(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &staticModel.Static{}
	query := make(map[string]interface{})
	list := &[]staticModel.Static{}

	err := c.ShouldBindQuery(model)
	if api.ValidateHasError(model) {
		return
	}

	utils.Struct2Map(query)
	if funk.IsEmpty(query) {
		list, err = model.GetAllList()
	} else {
		list, err = model.GetListByQuery(query)
	}

	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.Response(list)
}
