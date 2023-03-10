package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// IF - Switch

	x := 23 < 5

	b := rand.Int() % 2

	if 15 < 21 {
		fmt.Println("Hi guys")
	} else if x {
		fmt.Println(x)
	}

	switch b {
	case 0:
		b = 1
		fmt.Println(b)
	case 1:
		b = 1
		fmt.Println(b)
	default:
		fmt.Println("Go yu denedigin icin tesekurler")
	}

}
