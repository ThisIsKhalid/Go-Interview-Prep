package main

import "fmt"

func main () {

	// while loop ( in go there is no while)
	// i := 1

	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// classic for loop
	// for i :=  0; i <= 3; i++ {

	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// range keyword
	for i :=  range 3 {
		fmt.Println(i)
	}
}
