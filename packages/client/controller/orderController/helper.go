package orderController

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"github.com/ulule/deepcopier"
	productModel2 "hd-mall-ed/packages/admin/models/productModel"
	"hd-mall-ed/packages/client/models/orderModel"
	"hd-mall-ed/packages/client/models/productModel"
	"hd-mall-ed/packages/client/models/shoppingCartModel"
	"hd-mall-ed/packages/common/pkg/app"
)

// 减少库存
func updateReduceInventoryService(c *gin.Context, body *createReq) {
	api := app.ApiFunction{C: c}
	adminProductMapper := &productModel2.Product{}
	orderMapper := &orderModel.Order{}
	shoppingCartMapper := &shoppingCartModel.ShoppingCart{}
	productMapper := &productModel.Product{}

	// 复制信息
	_ = deepcopier.Copy(body).To(orderMapper)

	shoppingCartUpdateMap := make(map[string]interface{})
	shoppingCartUpdateMap["type"] = 3

	shoppingCartMapper.TempOrderId = orderMapper.OrderId
	shoppingCartMapper.UserId = uint(api.GetUserId())

	// 通过 product id 查询所有列表
	var productIdList []uint

	// 如果没有值就直接返回即可
	if funk.IsEmpty(body.List) {
		return
	}
	for _, item := range body.List {
		productIdList = append(productIdList, item.ProductId)
	}
	list := productMapper.GetDetailByIdList(&productIdList)
	if funk.IsEmpty(list) {
		return
	}

	for index, item := range *list {
		for _, item2 := range body.List {
			if item2.ProductId == item.ID {
				(*list)[index].Inventory = item.Inventory - item2.Count
				(*list)[index].Sales = item.Sales + item2.Count
			}
		}
	}

	for _, item := range *list {
		if item.Inventory <= 0 {
			return
		}
		// 更新库存
		go adminProductMapper.StandUpdate("id = ?", item.ID, &map[string]interface{}{
			"inventory": item.Inventory,
			"sales": item.Sales,
		})
	}
}
