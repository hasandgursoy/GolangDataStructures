package main

import "fmt"

func arrays() {

	arrays := []int{10, 20, 30, 40, 50}

	for a, b := range arrays {

		fmt.Println(a, ". index degeri : ", b)
	}

	for _, b := range arrays {

		fmt.Println(b)
	}

	for a := range arrays {

		fmt.Println(a)
	}

	powArray := getPowArray(arrays)
	fmt.Println(powArray)

	// Range komutu

}

func getPowArray(array []int) []int {

	for i := 0; i < len(array); i++ {

		array[i] = array[i] * array[i]
	}

	return array

}
