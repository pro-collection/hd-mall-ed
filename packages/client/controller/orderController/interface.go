package orderController

import "hd-mall-ed/packages/client/models/orderModel"

type createRes struct {
	OrderId uint `json:"order_id"`
	Id uint `json:"id"`
}

type productReduceCountStruct struct {
	ProductId uint `json:"product_id"`
	Count int `json:"count"`
}

type createReq struct {
	orderModel.Order
	List []productReduceCountStruct `json:"list"`
}
