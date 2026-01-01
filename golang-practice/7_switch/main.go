package main

import (
	"fmt"
	"time"
)


func main () {

	// i := 5

	// simple switch case
	// switch i {
	// 	case 1:
	// 		fmt.Println("one")
	// 	case 2:
	// 		fmt.Println("Two")
	// }

	// -> break is not needed in go

	// multiple switch case
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday: 
			fmt.Println("It's weekend")

		default:
			fmt.Println("it's workday")
	}

	// type switch
	whoami := func(i interface{} ) {
		switch t := i.(type) {
			case int:
				fmt.Println("It's an integer")
			case string:
				fmt.Println("It's a string")
			case bool:
				fmt.Println("It's a boolean", t)
		}
	}

	whoami("golang")

}
