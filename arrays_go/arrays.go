package main

import "fmt"

// number sequence of specific length
func main() {
	var nums [4]int
     nums[0] = 1 
	fmt.Println(len(nums))
	fmt.Println(nums[0])
	fmt.Println(nums)
	// if we dont add any element in array , it adds the zeroth value for the whole array ,incase of integer its zero , 
	var vals [4]bool 
	fmt.Println(vals) // here the vals value will be false 
	var names [3]string
	names[0] = "ali"
	fmt.Println(names) // incase of empty is "" --> its a empty size 


	nums1 := [3]int{2,3,5}
	fmt.Println(nums1)

	// How can we make the 2d arrays 
	 
	nums2 := [2][2]int{{2,3},{4,5}}
	fmt.Println(nums2)
	// -fixed size 
	// memory optimizatio 
	// constant time access 
	
}
