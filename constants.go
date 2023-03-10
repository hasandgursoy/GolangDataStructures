package main

import (
	"fmt"
	"math"
)

func main()  {
	
const pi float64 = math.Pi
const log2E float64 = math.Log2E

// Golang doesn't have another using for constant
// Just this use

areaOfCircle := pi * (math.Pow(4,4))

fmt.Printf("%T , %v" ,areaOfCircle,areaOfCircle)

}