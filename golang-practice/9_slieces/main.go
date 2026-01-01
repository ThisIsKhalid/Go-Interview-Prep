package main

import "fmt"


func main () {
	// slice -> dynamic array
	// most used construct in go
	// + use full methods



	// uninitialized slice in nill
	// var nums []int
	// fmt.Println(nums == nil) // output: true

	// fmt.Println(len(nums)) // output: 0

	var nums = make([]int, 2, 5)
	 // third element is basically capacity => maximum numbers of elemetns can fit

	fmt.Println(nums == nil)

}
