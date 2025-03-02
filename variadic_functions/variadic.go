package main

import "fmt"

// variadic function are the function which accepts n number of arguments

func main() {
	// suppose we want to pass as slice 

	nums := []int{3,4,5,7,8}
    result1 := sum(nums...)
	// This is how we normally pass the parameters 
	result := sum(3, 4, 5, 6)

	fmt.Println(result)
	fmt.Println(result1)

}

// here it can accept any number of argument
//interface{} - denotes that we can pass any type 
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}
	return total

}
