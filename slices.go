package main

import (
	"fmt"
	"sort"
)

func slices() {

	exmpSlice := []int{}
	exmpSlice = append(exmpSlice, 10)

	fmt.Println(exmpSlice[len(exmpSlice)-1])

	fmt.Println(exmpSlice)

	// create slice with make
	// make ile slice oluşturursak gider zero verilen uzunluğa göre zero değer atar.
	var makeSlice []int = make([]int, 0)
	makeSlice = append(makeSlice, 20, 30, 40)
	fmt.Println(makeSlice)

	// Array ve Slice arasındaki temel fark

	// Array deger tiplidir. Pass by value
	// Slice referance tiplidir. Pass by reference
	// Bu durumda bir slice'i baska bir degere atarsak slice da degisiklik olursa onu kullanan butun yapilar etkilenir

	// example

	slc1 := []int{10, 20, 30}
	slc2 := slc1

	slc2[0] = 4000

	// Bu kosulda ikisinde 0. indexinin degerinin degismesini bekleriz.

	fmt.Println(slc1)
	fmt.Println(slc2)

	// ... işareti underlaying Array demektir.
	underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	mySlc := underArray[0:]
	fmt.Println(mySlc)

	mySlc2 := underArray[:8]
	fmt.Println(mySlc2)

	sort.Slice(mySlc, func(i, j int) bool {
		return mySlc[i] > mySlc[j]
	})

	// Variadic funcs

	mySlc = append(mySlc, mySlc2...)
	fmt.Println(mySlc)

	// copy methodu ve = asign islemi arasindaki fark
	// copy de yeni bir underlaying array oluşturur referance tutmaz degeri tutar
	// assign da direk gider reference'sıyla birlikte atar ona.

	aslc := mySlc

	fmt.Println(&aslc)

	deger := copy(mySlc, aslc)
	fmt.Println(deger)
	fmt.Println(&aslc)

}
