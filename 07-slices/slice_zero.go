package main

import "fmt"

func main() {
	langs := []string{}
	fmt.Printf("langs: %#v\n", langs)

	if langs == nil{
		fmt.Println("Yes nil 'lang' is a nil slices")
	} else {
		fmt.Printf("lang is NOT nil , value %#v\n", langs)
	}
}
