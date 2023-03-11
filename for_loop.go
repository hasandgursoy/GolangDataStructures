package main

import "fmt"

func for_loop() {

	// Conventional For

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Weird looks like while loop
	a := 0
	for a < 5 {
		fmt.Println(a)
		a++
		if a == 5 {
			break
		}
		continue
	}

	// Infinitive loop

	// for {
	// 	fmt.Println("selamlar")
	// }

}
