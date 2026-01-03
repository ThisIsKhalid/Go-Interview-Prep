package main

import "fmt"

func main() {

	// var x int = 10
	// x = "hello" // compile time error

	// x := 10

	// if true {
	// 	x := 20

	// 	fmt.Println(x)
	// }

	// fmt.Println(x)

	var x int
	y := 5

	if y > 3 {
		x := 10
		fmt.Println(x)
	}

	fmt.Println(x)

}
