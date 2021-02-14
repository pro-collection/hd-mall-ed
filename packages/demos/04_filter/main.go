package main

import (
	"fmt"
	"github.com/thoas/go-funk"
	"reflect"
)

func main() {
	//r := funk.Filter([]int{1, 2, 3, 4}, func(x int) bool {
	//	return x%2 == 0
	//}) // []int{2, 4}

	r2 := funk.Filter([]string{"yanle", "lele", "yanlele", "heihei"}, func(value string) bool {
		fmt.Println(value)
		return len(value) <=5
	}) // []int{2, 4}

	//fmt.Println(r)
	fmt.Println(reflect.TypeOf(r2))
}
