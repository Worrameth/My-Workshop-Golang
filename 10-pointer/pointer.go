package main

import "fmt"

func main() {
	me := "Golang"
	fmt.Printf("You are %v\n", me)

	//addr := &me
	var addr *string = &me
	fmt.Println(addr)
	fmt.Printf("%T\n", addr)

	*addr = "Worrameth"
	fmt.Printf("You are %v\n", me)
	fmt.Println(addr)
}
