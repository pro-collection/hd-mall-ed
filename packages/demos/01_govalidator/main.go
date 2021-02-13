package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

type addressStruct struct {
	AddressName string
	AddressCode string
}

type userStruct struct {
	Name    string        `valid:"required~缺少姓名"`
	Age     int           `valid:"required~缺少年龄"`
	Address addressStruct `valid:"required~缺少地址信息"`
}

func main() {
	user := userStruct{
		Name: "yanle",
		Age: 30,
		Address: addressStruct{"12", ""},
	}

	result, err := govalidator.ValidateStruct(&user)
	fmt.Println("result: ", result)
	if err != nil {
		fmt.Println(strings.Split(err.Error(), ";")[0])
	}
}
