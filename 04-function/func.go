package main

import "fmt"

func info(name, msg string, age int) {
	fmt.Printf("My name is %v \n", name)
	fmt.Printf("My message is %v \n", msg)
	fmt.Printf("I'm %v year old \n", age)
}

func today() string {
	return "วันนี้"
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	info("Pond", "Nickname", 23)
	info("Worrameth", "Name", 23)
	fmt.Println("")

	day := today()
	fmt.Println(day)

	a, b := swap("a", "b")
	fmt.Println(a, b)
}
