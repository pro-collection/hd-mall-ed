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

	var resList []getDetailByIDResStruct

	//ch := make(chan productModel.Product, len(*list))

	for _, cart := range *list {
		productMapper := &productModel.Product{}
		var res getDetailByIDResStruct
		err = productMapper.GetById(int(cart.ProductId))
		_ = deepcopier.Copy(&cart).To(&res)
		res.ProductInfo = *productMapper
		resList = append(resList, res)
	}

	api.Response(resList)
}

// 通过临时订单获取所有信息
func GetDetailByID(c *gin.Context) {
	api := adminApp.ApiInit(c)

	model := &shoppingCartModel.ShoppingCart{}

	idString := c.DefaultQuery("id", "")
	if idString == "" {
		api.ResFail(e.Fail)
		return
	}

	id, _ := strconv.Atoi(idString)
	model.ID = uint(id)
	err := model.GetDetailById()

	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	api.Response(model)
}

func GetDetailByTempOrderId(c *gin.Context) {
	api := adminApp.ApiInit(c)

	model := &shoppingCartModel.ShoppingCart{}

	tempOrderIdString := c.DefaultQuery("temp_order_id", "")
	if tempOrderIdString == "" {
		api.ResFail(e.Fail)
		return
	}

	id, _ := strconv.Atoi(tempOrderIdString)
	model.TempOrderId = uint(id)

	list, err := model.GetListByTempOrderId()
	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	api.Response(list)
}
