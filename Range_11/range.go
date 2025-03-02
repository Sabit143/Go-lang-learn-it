package main

import "fmt"

// iterating over the slices
func main() {
	// nums := []int{6, 7, 8}
	   

	// sum := 0;

	// for i, num := range nums {
	// 	fmt.Println(i)
	// 	sum = sum + num

	// }
	// fmt.Println(sum)

	m := map[string]string{"fname":"jhon","lname":"ali"}


    // suppose we want to iterate both the key and map 
	for k, v := range m {
		fmt.Println(k,v)
	}

	// suppose we want to print only the key 

	for k := range m{
		fmt.Println(k)
	}
      // Here it will print only the value here 
	for _,v := range m{
		fmt.Println(v)
	}
        // unicode code point - rune 
		 for k , l := range "golang"{
			fmt.Println(k,l)
		 }

}
