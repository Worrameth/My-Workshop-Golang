package main

import "fmt"

func main() {
	langs := []string{"golang", "node", "vue"}
	fmt.Printf("langs: %#v\n", langs)
	
	a := langs[0:2] // [0, 2] เอา 0 ถึง 2 แต่ไม่เอา 2
	fmt.Printf("a: %#v\n", a)
	fmt.Printf("lang: %#v\n",langs)

	b := langs[1:3]
	fmt.Printf("b: %#v\n", b)
	fmt.Printf("lang: %#v\n",langs)

	c := langs[0:len(langs)]
	fmt.Printf("c: %#v\n", c)
	fmt.Printf("lang: %#v\n",langs)

	r := langs[0:3]
	s := langs[:3]
	t := langs[0:]
	u := langs[:]

	fmt.Printf("r: %#v\n", r)
	fmt.Printf("s: %#v\n", s)
	fmt.Printf("t: %#v\n", t)
	fmt.Printf("u: %#v\n", u)

	printSlice(langs)
	cords := [4]string{"Am","B","G","F#"}
	printSlice(cords[:])
}

func printSlice(ns []string){
	fmt.Printf("printSlice: ns: %#v\n",ns)
}