package main

import (
	"fmt"
)

// multiple return value

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}

	return a / b, nil
}

// passing funcition as argument
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	// Functions Are First-Class Citizens

	// function as variables
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(100, 99))

	fmt.Println(operate(10, 4, add))

	// annonymous function
	func() {
		fmt.Println("Hello")
	}()

	sum := counter()
	fmt.Println("Sum ", sum())
	fmt.Println("Sum ", sum())
	fmt.Println("Sum ", sum())
	fmt.Println("Sum ", sum())

}
