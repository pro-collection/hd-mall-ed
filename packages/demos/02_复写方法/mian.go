package main

import "fmt"

type userStruct struct {
}

type meStruct struct {
	userStruct
}

func (user *userStruct) println() {
	fmt.Println("userStruct")
}

func (*userStruct) getName() {
	fmt.Println("get name")
}

func (*meStruct) println() {
	fmt.Println("meStruct")
}

func main() {
	user := &userStruct{}
	user.println()
	user.getName()

	me := &meStruct{}
	me.println()
	me.getName()
}
