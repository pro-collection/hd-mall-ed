package main

import "fmt"

type userStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	list := []userStruct{}
	list = append(list, userStruct{Name: "yanle"})
	list = append(list, userStruct{Name: "lele"})

	fmt.Println(list)

	for index, _ := range list {
		list[index].Age = userStruct{Age: 15}.Age
	}

	fmt.Println(list)
}
