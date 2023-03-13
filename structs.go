package main

import "fmt"

func main() {

	type employee struct { // underlying type
		name      string
		age       int
		isMarried bool
		kids      []string
	}

	// Kalıtım gibi birşey yapıyor burda garip bir kullanim sekli var acikcasi
	// Has A Relation
	// Diğer diller deki gibi Is A Relation degil

	type manager struct {
		employee
		hasDegree bool
	}

	manager1 := manager{
		employee: employee{
			name:      "Marcus",
			age:       32,
			isMarried: false,
			kids:      []string{},
		},
		hasDegree: true,
	}

	fmt.Println(manager1)

	emp1 := employee{
		name:      "hasan",
		age:       30,
		isMarried: false,
		kids: []string{
			"marcus",
			"micheal",
		},
	}

	fmt.Println(emp1)

	// Anonim Struct

	theBoss := struct {
		name  string
		money bool
	}{name: "SP", money: true}

	fmt.Println(theBoss)

}
