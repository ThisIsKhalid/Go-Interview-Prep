package main

import "fmt"


func main () {
	age := 14

	if age >= 18 {
		fmt.Println("Persone is Adult")
	} else if age >= 12 {
		fmt.Println("Person is teenager")
	} else {
		fmt.Println("Person is under age")
	}

	if salary := 20000; salary >= 1000 { // we can declare variable directly here also
		fmt.Println("Low salary")
	}

	// go does not have ternary, you will have to use normal if else
}
