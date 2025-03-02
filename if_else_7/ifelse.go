package main

import "fmt"

func main() {
	// age := 18

	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// }else if age<=18{
	// 	fmt.Println("person is not an adult")
	// }else{
	// 	fmt.Println("invalid age")
	// }

	var role = "admin"
	var hasPermission = false
	if role == "admin" || hasPermission {
		fmt.Println("yes")
	}
    // we can declare a variable inside the if , so we 
	if age := 15 ; age <=18 {
          fmt.Println("the age is",age)
	}
}
