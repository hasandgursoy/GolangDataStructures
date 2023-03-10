package main

import (
	"fmt"
	"strconv"
)

func main()  {
	
	var (
		age int32 = 46
		height float64 = 182.30
		weight string = "90"
	)

	// Int To Float

	fmt.Printf("Age : %v , Type : %T\n " , float32(age),float32(age))

	// Float To Int
	fmt.Printf("Height : %v , Type : %T\n " , int16(height),int16(height))

	// String To Int

	intWeight,_ := strconv.Atoi(weight)

	fmt.Printf("intWeight : %v , Type : %T\n " ,intWeight,intWeight)

	// Int To String

	stringWeight := strconv.Itoa(intWeight)

	fmt.Printf("stringWeight : %v , Type : %T \n",stringWeight,stringWeight)
	
	// Float To String

	// way 1

	stringHeight1 := strconv.FormatFloat(height,'f',2,64)
	fmt.Println(stringHeight1)

	fmt.Printf("stringHeight : %v , Type : %T \n", stringHeight1,stringHeight1)
	
	// way 2

	stringHeight2 := fmt.Sprintf("%v",height)

	fmt.Println(stringHeight2)
	

	// String To Float

	if floatHeight,err := strconv.ParseFloat(stringHeight2,64); err == nil {
		fmt.Printf("floatHeight : %v , Type : %T ",floatHeight,floatHeight)
	}
	

}