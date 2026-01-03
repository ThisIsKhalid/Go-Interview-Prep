package main

import "fmt"

func test() {
	defer fmt.Println("A")
	defer fmt.Println("B")

	for i := 0; i < 2; i++ {
		defer fmt.Println(i)
	}
}

func main() {

	test()
}
