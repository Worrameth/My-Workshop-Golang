package user

import "fmt"

func Info(name, mgs string, age int) {
	fmt.Printf("My name is %v\n", name)
	fmt.Printf("My message is %v\n", mgs)
	fmt.Printf("I'm %v year old\n", age)
}
