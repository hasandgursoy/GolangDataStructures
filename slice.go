package main

import(
"fmt"
"time"
)

func add(y int , x int ) int {
	return x + y
}
func swap(x,y string) (string,string) {
	return x,y
}

func deneme()  {
	fmt.Println("Welcome to the playgorund !")
	fmt.Println("The time is ",time.Now().UTC())
	fmt.Println(add(10,20))
	fmt.Println(swap("Hasan","Gursoy"))
}