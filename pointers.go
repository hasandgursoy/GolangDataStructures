package main

import (
	"fmt"
)

func pointers() {
	a, b := 56, 97

	fmt.Println(b, a)

	i := 0
	for i < 20 {
		i++
	}

	var mehmetali float64 = 1

	fmt.Println(mehmetali)

	// 0x14000014070 = 56
	// 0x14000014070 = 97

}
