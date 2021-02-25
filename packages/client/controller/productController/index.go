package productController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/productModel"
	"hd-mall-ed/packages/common/pkg/app"
)

func GetPrimaryCategoryProductList(c *gin.Context) {
	model := &productModel.Product{}
	api := app.ApiFunction{C: c}
	list := model.GetPrimaryCategoryAndProductList()
	api.Response(list)
}
