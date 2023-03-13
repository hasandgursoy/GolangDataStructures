package main

import "fmt"

// Bu artık kendi veri tipimiz
type mile float64
type kilometer float64

func own_types()  {
	var m1 mile = 3.2
	var km1 kilometer = 90.3
	
	fmt.Println(m1)
	fmt.Println(km1)

	fmt.Printf("%T type , %v value",m1,m1)

	var f1 float64 = 4.4

	// m1 variable i float64 tabanında mile turunde kurulmus olsada f1 float64 tipindeki degisken ile 
	// toplattırmadi bu yuzden type conversion islemine basvurmak zorunda kaldim bu islemde
	// 2 side ayni underline type a sahip olsada islem yaptirmiyor.

	fmt.Println(m1 + mile(f1))


}