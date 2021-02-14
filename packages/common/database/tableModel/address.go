package tableModel

type AddressBase struct {
	ID            uint   `json:"id" gorm:"primarykey"`
	UserId        uint   `json:"user_id" gorm:"not null;index;comment:关联的userId"`
	Province      string `json:"province" gorm:"not null;comment:省市区" valid:"required~缺少省市区信息"`
	AddressDetail string `json:"address_detail" gorm:"not null;type:longtext;comment:地址详情" valid:"required~缺少地址详情"`
	AddresseeName string `json:"addressee_name" gorm:"not null;comment:收件人姓名" valid:"required~缺少收件人姓名"`
	Mobile        string `json:"mobile" gorm:"not null;comment:收件人邮箱" valid:"required~缺少收件人电话号码"`
}

type Address struct {
	ModelStruct
	AddressBase
}
