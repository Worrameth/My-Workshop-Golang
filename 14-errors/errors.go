package main

import (
	"errors"
	"fmt"
	"log"
)

func ReadFile(name string) (string, error) {
	// read file........
	return "data........", errors.New("File Not Found")
}

func main() {
	data, err := ReadFile("profile.txt")
	if err != nil {
		log.Println(err)
		return 
	}
	fmt.Println("read file success:",data)
}
