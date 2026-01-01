package main

import "fmt"

const age = 30; // can be assigned outside of funtion

const (
	port = 5000
	host = "localhost"
)

var (
	salary = 5000
	location = "Dhaka"
)

func main() {

	// var  name string = "go lang"

	// var name = "go lang"

	// name := "go lang"

	// var name string
	// name = "go lang"

	const name = "go lang" // can not be assigned later


	fmt.Println(name)


}
