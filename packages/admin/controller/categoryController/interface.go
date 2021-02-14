package categoryController

type deleteParamsStruct struct {
	Id int `json:"id" valid:"required~ID缺失"`
}
