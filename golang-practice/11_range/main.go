package main

import "fmt"

// iterating over data structure
func main () {

	// nums := []int {5, 6, 7}
	// fmt.Println(len(nums))

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// better version using range
	// for _, num := range nums {  // _ this here is basically if you want index number also
	// 	fmt.Println(num)
	// }


	// sum := 0

	// for i, num := range nums { 
	// 	sum = sum + num
	// 	fmt.Println(i , num)
	// }
	// fmt.Println(sum)



	// ------------ map/object --------------
	// m := map[string]string{"fname": "john", "lname": "doe"}

	// for key, value := range m {
	// 	fmt.Println(key, value)
	// }

	// for key := range m {
	// 	fmt.Println(key)
	// }

	// for value := range m {
	// 	fmt.Println(value)
	// }




	// string
	for i, c := range "go lang" {
		fmt.Println(i, string(c))
	}

}
