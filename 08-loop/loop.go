package main

import "fmt"

func main() {
	// for
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("i: %#v\n", i)
	//}
	
	// while
	// i := 0
	// for i < 10 {
	// 	fmt.Printf("i: %#v\n", i)
	// 	i++
	// }

	langs := []string{"golang", "node", "vue","F#"}
	fmt.Println("\nfor basic")

	for i := 0; i < len(langs); i++ {
		value := langs[i]
		fmt.Printf("%v: %v\n",i ,value)
	}

	fmt.Println("\nfor range slice")
	for _, value := range langs {
		fmt.Printf("value : %v\n", value)
	}

	fmt.Println("\nfor range index")
	for index := range langs {
		fmt.Printf("%v : slice\n", index)
	}
}
