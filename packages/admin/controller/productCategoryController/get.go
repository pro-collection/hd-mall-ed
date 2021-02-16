package productCategoryController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/admin/models/productCategoryModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"log"
	"strconv"
)

func GetList(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &productCategoryModel.ProductCategory{}
	idString := c.DefaultQuery("product_id", "")
	id, err := strconv.Atoi(idString)

	log.Println("funk.IsEmpty(idString) : ", funk.IsEmpty(idString))

	if funk.IsEmpty(idString) {
		err = errors.New(e.GetMsg(e.NotFoundId))
		api.ResFail(e.NotFoundId)
		return
	}

	model.ProductId = uint(id)

	// 查询逻辑
	list, err := model.GetList()

	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.Response(list)
}
