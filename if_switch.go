package main

import (
	"fmt"
	"math/rand"
)

func if_switch() {

	// IF - Switch

	x := 23 < 5

	b := rand.Intn(10)

	if 15 < 21 {
		fmt.Println("Hi guys")
	} else if x {
		fmt.Println(x)
	}

	// falltrough ifadesi statement dogru olsa bile git diğerlerine de bak demek.

	switch {
	case b < 5:
		fmt.Println(b)
		fallthrough
	case b > 5:
		fmt.Println(b)
		fallthrough
	default:
		fmt.Println("Go yu denedigin icin tesekurler")
	}

}
