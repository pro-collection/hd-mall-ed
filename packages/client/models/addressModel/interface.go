package addressModel

type UpdateRequestParamsStruct struct {
	Province      string `json:"province" valid:"required~缺少省市区信息"`
	AddressDetail string `json:"address_detail" valid:"required~缺少地址详情"`
	AddresseeName string `json:"addressee_name" valid:"required~缺少收件人姓名"`
	Mobile        string `json:"mobile" valid:"required~缺少收件人电话号码"`
}

type DeleteRequestParamsStruct struct {
	Id int `json:"id" valid:"required~缺少 id 参数"`
}

type UpdateDefaultRequestParamsStruct = DeleteRequestParamsStruct
