package main

import "fmt"

var packVar = "Package Scpoe"

func scopes()  {

	// Scope yapisi diger dillerindekinden farkli degil akil yuruterek bulunabilinir.

	if true {
		var blockvar string = "Block var"
		fmt.Println(blockvar)
	
	}

	var funcvar = "funcScope"

	fmt.Println(packVar)
	fmt.Println(funcvar)

	ScopeFunc()
}

func ScopeFunc()  {
	var scopefunc int = 15
	fmt.Printf("Scope Func %v", scopefunc)
}