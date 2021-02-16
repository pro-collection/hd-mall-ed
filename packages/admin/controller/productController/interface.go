package productController


type deleteParamsStruct struct {
	Id int `json:"id" valid:"require"`
}
