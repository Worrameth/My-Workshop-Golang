package main

import "fmt"

var name = "Pond"

//var plus = func(a, b int) int { return a + b }
func say(n string) {
	fmt.Printf("My name is %s\n", n)
}

func cal(op func(int, int) int) {
	r := op(4, 6)
	fmt.Printf("Result : %v\n", r)
}

func main() {
	fmt.Printf("Value: %v\n", name)
	fmt.Printf("Type: %T\n", name)
	say(name)
	plus := func(a, b int) int { return a + b }
	p := plus(2, 3)
	fmt.Printf("Value: %v\n", p)
	fmt.Printf("Type: %T\n", plus)

	cal(plus)

	minus := func(a, b int) int { return a - b }
	cal(minus)
}
