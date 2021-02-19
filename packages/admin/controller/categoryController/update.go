package categoryController

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/admin/models/categoryModel"
	"hd-mall-ed/packages/common/database/tableModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Update(c *gin.Context) {
	api := adminApp.ApiInit(c)
	var err error
	model := &categoryModel.Category{}

	// 获取请求参数
	params := &tableModel.CategoryBase{}
	err = c.ShouldBindJSON(params)
	if funk.IsEmpty(params.ID) {
		api.ResFail(e.InvalidParams)
		return
	}

	// 校验参数
	if api.ValidateHasError(params) {
		return
	}

	// 复制参数
	err = deepcopier.Copy(params).To(model)

	// 校验是否是存在
	if handleUpdateCheckUserNameIsExistHelper(model, api) {
		return
	}

	// 如果是创建主分类的场景，需要检验限制
	if !funk.IsEmpty(model.ParentId) && handleListCheckHelper(model, api) {
		return
	}

	// 更新
	err = model.Update()

	// 数据返回
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}
	api.ResponseNoData()
}
