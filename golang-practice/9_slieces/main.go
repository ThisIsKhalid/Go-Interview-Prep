package main

import (
	"fmt"
)


func main () {
	// slice -> dynamic array
	// most used construct in go
	// + use full methods



	// uninitialized slice in nill
	// var nums []int
	// fmt.Println(nums == nil) // output: true

	// fmt.Println(len(nums)) // output: 0

	// var nums = make([]int, 0, 5)
	 // third element is basically capacity => maximum numbers of elemetns can fit

	//  fmt.Println(cap(nums)) // output: 5
	// fmt.Println(nums == nil)

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)

	// fmt.Println(nums) // output: [0 0 1 2 3 4 5 6]
	// fmt.Println(cap(nums)) // output: 10

	// copy slice
	

	// nums = append(nums, 2)

	// var nums2 = make([]int, len(nums) )

	// copy(nums2, nums)

	// fmt.Println(nums, nums2)


	// slice oparator

	// var nums = []int {1,2,3,4,5,6,7,8,9}

	// fmt.Println(nums[2:5]) // output: [3 4 5]

	// fmt.Println(nums[:3]) // output: [1 2 3]

	// fmt.Println(nums[5:]) // output: [6 7 8 9]


	// Slices package comparison

	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int {1, 2, 3}

	// fmt.Println(slices.Equal(nums1, nums2))


	// 2 D array
	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)


}


// this is important topic