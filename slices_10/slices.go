package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used onstruct in go
// + useful methods it provides
func main() {
	// uninitilized slice is nil
	// var nums []int
	// fmt.Println(nums)
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	var nums = make([]int, 2, 5)
	// This is the initial size we have given
	// we normally keep the lenghth is zero , so that it doesn't print the first zero , zero
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
	nums = append(nums, 1)
	nums = append(nums, 6)
	nums = append(nums, 10)
	nums = append(nums, 34)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
	fmt.Println(nums)

	// nums2 := []int{}
	// fmt.Println(nums2)
	// nums2 = append(nums2, 10)
	// fmt.Println(cap(nums2))
    // // Here we are using the copy method 
	// var nums01 = make([]int, 0, 5)
	// nums01 = append(nums01, 2)
	// var nums02 = make([]int, len(nums01))
    
	//  copy(nums02,nums01)
	//  fmt.Println(nums01)
	//  fmt.Println(nums02)

	 // slice operator 
	//  var arr = []int{1,2,3,4}
	//  fmt.Println(arr[0:4]) //1,2,3,4 
	//  fmt.Println(arr[1:4]) // 2,3,4
	//  fmt.Println(arr[:2]) // it will show upto 1,2
	//  fmt.Println(arr[1:]) // here the first element is for skipping the number of element , and the last one is upto 

	// comparing the slices 
	 var nums1 = []int{1,2,3}
	 var nums2 = []int{1,2,4}
	 fmt.Println(slices.Equal(nums1,nums2)) // here the slices have different methods 
	 // slices package has alot of methods
      // 2D slices 
	 var nums0 = [][]int{{1,2,3},{4,5,6}}
	 fmt.Println(nums0);





}
