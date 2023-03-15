package main

import (
	"fmt"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	c.Radius = 10 // struct örneği üzerinde kalıcı değişiklik yapmaz
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) AreaPointer() float64 {
	c.Radius = 10 // struct örneği üzerinde kalıcı değişiklik yapar
	return 3.14 * c.Radius * c.Radius
}

func main() {
	c1 := Circle{Radius: 5}
	c2 := Circle{Radius: 5}

	fmt.Println("Before calling any method")
	fmt.Println("c1:", c1.Radius) // c1: 5
	fmt.Println("c2:", c2.Radius) // c2: 5

	fmt.Println("After calling Area method")
	c1.Area()                     // struct kopyası üzerinde işlem yapar, değişiklikleri kaydetmez
	fmt.Println("c1:", c1.Radius) // c1: 5

	fmt.Println("After calling AreaPointer method")
	c2.AreaPointer()              // struct üzerinde doğrudan işlem yapar, değişiklikleri kaydeder
	fmt.Println("c2:", c2.Radius) // c2: 10

	/*
			main() fonksiyonunda, önce c1 ve c2 isimli iki adet Circle struct'ı örneği oluşturulur. Daha sonra, her iki struct örneği üzerinde de önce Area() metodu çağrılır. Area() metodunun çağrıldığı c1 struct'ının örneği üzerinde bir değişiklik yapılmadığı görülür. Ancak, AreaPointer() metodunun çağrıldığı c2 struct'ının örneği üzerinde Radius özelliği 10'a ayarlanır ve bu değişiklik kalıcı olur.

		Sonuç olarak, fonksiyon alıcısında pointer kullanmak, struct örneği üzerinde kalıcı değişiklikler yapmamızı sağlar.
	*/
}
