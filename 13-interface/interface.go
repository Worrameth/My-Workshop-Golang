package main

import "fmt"

type Phone interface {
	Call(number string)
}

type Samsung struct {
	Name string
}

func (s Samsung) Call(number string){
	fmt.Println(s.Name, "calling",number)
}

func (s Samsung) Answer(){
	fmt.Println(s.Name,"Hello There!")
}

func Dial (p Phone){
	p.Call("+6666666666")
}
func main() {
	s := Samsung{
		Name: "S10",
	}

	Dial(s)
	s.Answer()
}