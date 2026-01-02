package main

import (
	"fmt"
	"maps"
)

func main () {

	// map -> hash,  object, dict

	// creating a map
	// m := make(map[string]string)

	// setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"
	// m["price"] = "100"

	// get an element
	// fmt.Println(m["name"], m["area"], m["price"])

	// if key does not exist in the map then it returns zero/blank/false value base on int, string, bool
	// fmt.Println(m["age"])

	// get length
	// fmt.Println(len(m))

	// delete element from map
	// delete(m, "price")

	// all clear
	// clear(m)

	// fmt.Println(m["name"], m["area"], m["price"])

	// --------------------------

	// m := map[string]int{"price": 40, "phones": 3}

	// value, ok := m["prices"] // first is value, and second is bool: true or false

	// fmt.Println(value)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("Not ok")
	// }



	// -------------------------------
	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 4}

	fmt.Println(maps.Equal(m1, m2))
}