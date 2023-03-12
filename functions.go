package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("My Age : %v ",(sayMyAge(1997)))

	bolum,kalan := divide(33,6)
	fmt.Printf("Bolum : %v , Kalan : %v ",bolum,kalan)


}

func sayMyAge(birthYear int) (age int) {

	age = time.Now().Year() - birthYear

	return age
}


func divide(bolunen, bolen int) (bolum,kalan int)  {
	
	bolum = bolunen / bolen
	kalan = bolunen % bolen

	return bolum,kalan

}