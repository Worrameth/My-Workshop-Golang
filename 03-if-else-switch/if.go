package main

import "fmt"

func main() {
	num := 33
	if num%2 == 0 {
		fmt.Println("เลขคู่")
	} else if num == 31 {
		fmt.Println("ว๊าย ๆ ")
	} else {
		fmt.Println("เลขคี่")
	}
}
