package main

import "fmt"

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("in changeNum",num)
// }

// here we are changing the copy , not the original number , Here we are passing the reference 

func changeNum(num *int){
	*num = 5 
	fmt.Println("in changeNum", *num)
}
func main() {
    num := 1
	// changeNum(num)
	// fmt.Println("after changeNum in main",num)
     changeNum(&num)
	 fmt.Println("after change num",num)
	// suppose i want to print the address the of memory location 
	fmt.Println("Memeory address",&num) // Memeory address 0xc00000a0e8



}

// There could be the use cases , where we can change the number directly , for that reason , we need to change the refrence 
