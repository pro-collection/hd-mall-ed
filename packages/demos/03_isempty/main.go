package main

import (
	"fmt"
	"github.com/thoas/go-funk"
)

type userStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var count = 0

	user := &userStruct{}

	fmt.Println("funk.IsEmpty(user)", funk.IsEmpty(user))

	fmt.Println("funk.IsEmpty(count)", funk.IsEmpty(count))

}
