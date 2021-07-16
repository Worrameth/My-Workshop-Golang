package main

import "fmt"

func main() {
	//var name string = "Pond KuyDum"
	var name string
	var i int = 5
	var f float64 = 3.7
	var b = true
	r := 'P'
	//name := "Pond KuyDum" // := สามารถใช้ได้ใน func เท่านั้น

	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Type: %T\n", name)

	fmt.Printf("Int : %v\n", i)
	fmt.Printf("Float : %v\n", f)
	fmt.Printf("Bool : %v\n", b)
	fmt.Printf("Rune : %v\n", r)

}
