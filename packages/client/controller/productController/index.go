package productController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/productModel"
	"net/http"
)

func GetPrimaryCategoryProductList(c *gin.Context) {
	model := &productModel.Product{}
	list := model.GetPrimaryCategoryAndProductList()
	c.JSON(http.StatusOK, gin.H{
		"list": list,
	})
}
