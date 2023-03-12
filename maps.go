package main

import "fmt"

func maps()  {

	myMap := map[string]int{
		"Ahmet":40,
		"Ayşe":17,
		"Selim":14,
		"Mustafa":70,
	}
	
	fmt.Println(myMap)

	myDefaultMap := map[int]string{}

	fmt.Println(myDefaultMap)

	// olmayan bir key aradigimiz zaman bize default deger doner

	fmt.Println(myMap["asli"])

	// make methodu ile map oluşturabiliyoruz.

	myMakeMap := make(map[string]int)

	myMakeMap["selam"] = 15
	fmt.Println(myMakeMap)

	studentGrades := map[string]int{
		"Arin":80,
		"Ahmet":72,
		"Selim":90,
		"Cinar":0,
	}

	// Map in icinde ilgili deger varmı nasıl bakariz diyorsak degiskene atama yapiyoruz ikinci deger 

	value, ok := studentGrades["ciar"]
	fmt.Println(value,ok)

	// maps -> pass by reference

	delete(studentGrades,"Cinar")
	fmt.Println(studentGrades)

	for k,v := range studentGrades{
		fmt.Println("Key : ",k," Value : ",v)
	}



}
