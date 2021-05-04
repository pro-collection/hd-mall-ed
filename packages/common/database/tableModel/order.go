package tableModel

import "time"

/*
订单状态字段 status
	1 - 确认订单
	2 - 发货
	3 - 确认收货
	4 - 订单完成
*/
type OrderBase struct {
	ID uint `json:"id" gorm:"primarykey"`
	OrderId uint `json:"order_id" gorm:"index;comment:订单id,关联cart表 temp_order_id 字段" valid:"required"`
	Status int `json:"status" gorm:"comment:订单的状态"`
	SendTime time.Time `json:"send_time" gorm:"发货时间"`
	ConfirmTime time.Time `json:"confirm_time" gorm:"确认收货时间"`
	CompleteTime time.Time `json:"complete_time" gorm:"订单完成时间"`
}

type Order struct {
	OrderBase
	ModelStruct
}
