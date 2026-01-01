package main

import "fmt"

func main () {

	// maps -> hash,  object, dict

	// creating a map
	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"

	// get an element
	fmt.Println(m["name"])
}