package shoppingCartController

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/admin/models/productModel"
	"hd-mall-ed/packages/admin/models/shoppingCartModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"strconv"
)

// 参数 type
func GetList(c *gin.Context) {
	api := adminApp.ApiInit(c)

	// 必须要 type 参数
	model := &shoppingCartModel.ShoppingCart{}

	typeString := c.Query("type")
	typeNumber, _ := strconv.Atoi(typeString)
	model.Type = typeNumber

	model.UserId = uint(api.GetUserId())
	list, err := model.GetList()
	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	api.Response(list)
}

// 通过临时订单获取所有信息
func GetDetailByID(c *gin.Context) {
	api := adminApp.ApiInit(c)

	model := &shoppingCartModel.ShoppingCart{}
	productMapper := &productModel.Product{}
	response := &getDetailByIDResStruct{}

	idString := c.DefaultQuery("temp_id", "")
	if idString == "" {
		api.ResFail(e.Fail)
		return
	}

	id, _ := strconv.Atoi(idString)
	cart, err := model.GetDetailById(uint(id))

	err = deepcopier.Copy(cart).To(response)

	// 获取 product 信息
	err = productMapper.GetById(int(cart.ProductId))
	response.ProductInfo = *productMapper

	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	api.Response(response)
}
