package addressController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/common/models/addressModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
)

// 删除地址
// params:
// 	id	地址id	int
func Delete(c *gin.Context) {
	api := app.ApiFunction{C: c}

	query := addressModel.DeleteRequestParamsStruct{}
	if err := c.ShouldBindJSON(&query); err != nil {
		api.ResFail(e.FailBindJson)
		return
	}

	model := addressModel.Address{}
	model.ID = uint(query.Id)

	if err := model.DeleteAddress(); err != nil {
		api.ResFail(e.AddressDeleteFail)
		return
	}

	api.ResponseNoData()
}
