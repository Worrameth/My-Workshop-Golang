package main

import (
	"fmt"

	"github.com/Worrameth/testpack/time"
	"github.com/Worrameth/testpack/user"
)

func main() {
	user.Info("Worrameth", "ไง", 23)

	t := time.Today()
	fmt.Println("today is", t)
}
