package staticController

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/admin/models/staticModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
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
	queryParams := handleQueryParamsHelper(c)
	list := &[]staticModel.Static{}
	var err error

	_ = deepcopier.Copy(queryParams).To(model)
	if funk.IsEmpty(queryParams) {
		list, err = model.GetAllList()
	} else {
		list, err = model.GetListByQuery(queryParams)
	}

	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.Response(list)
}
