package main

import "fmt"

func main() {
	langs := []string{"golang", "node", "vue"}
	fmt.Printf("langs: %#v\n", langs)

	n := langs[0]
	fmt.Printf("langs[0]: %#v\n", n)

	langs[1] = "Nothing"
	fmt.Printf("%v\n", langs)

	l := len(langs)
	fmt.Println("lenght of langs: ",l)

	langs = append(langs,"F#","Em","Am")
	fmt.Printf("langs: %#v\n", langs)
	fmt.Println("lenght of langs: ",len(langs))
}
