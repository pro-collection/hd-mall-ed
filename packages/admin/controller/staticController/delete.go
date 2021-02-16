package staticController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/admin/models/staticModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"hd-mall-ed/packages/common/pkg/utils"
)

// @title           删除单个文件
// @description
// @auth            晴小篆  331393627@qq.com
// @param           id
// @return
func Delete(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &staticModel.Static{}

	err := c.ShouldBindJSON(model)
	if api.ValidateHasError(model) {
		return
	}

	if funk.IsEmpty(model.ID) {
		err = errors.New(e.GetMsg(e.NotFoundId))
		api.ResFail(e.NotFoundId)
		return
	}

	err = model.Delete()

	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}

// @title           批量删除文件
// @description
// @auth            晴小篆  331393627@qq.com
// @param     		staticModel.Static
// @return
func Deletes(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &staticModel.Static{}

	err := c.ShouldBindJSON(model)
	if api.ValidateHasError(model) {
		return
	}

	query := utils.Struct2Map(model)
	err = model.Deletes(query)
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
