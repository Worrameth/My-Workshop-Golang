package main

import "fmt"

func showAll(ns [4]string){
	fmt.Printf("show all : %#v\n",ns)
}

func main() {
	var langs = [4]string{"Golang", "Node", "Vue", "React"}
	fmt.Printf("langs: %#v\n", langs)

	n := langs[0]
	fmt.Printf("%v\n", n)
	
	langs[1] = "Nothing"
	fmt.Printf("%v\n", langs)
	
	showAll(langs)

	cords := [4]string{"Am","Gm","F7-F11"}
	showAll(cords)
}
