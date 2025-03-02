package main

import "fmt"

// by adding the T any  here , we can pass any kind of slice here  / or we can write like [T interface{} ]

func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// we can do incase of struct also 
// type stack struct{
// 	elements [] int
// }

type stack[T any] struct {
	elements[]T
}

// suppose we want only some data types it shall recive , so what we can do is 
func printSliceD[T int | string | bool](items[]T){
	for _,item := range items{
		fmt.Println(item)
	}
}
/**
we can use the comparable instead of writing the whole thing 
 Here comparable denotes all the data types here 
 func printSliceD[T comparable]{
 for _,item := range items{
	 	fmt.Println(item)
	}
 }
*/
func main() {
   printSlice([]int{1,2,3,5})
   names := []string{"ali", "golang"}
   printSlice(names)
   names2 := []bool{false,true,false}
   printSliceD(names2)
    
   myStack :=stack[int]{
	elements: []int{1,2,3,4,5},
   }
   fmt.Println(myStack) //{[1 2 3 4 5]}

   myStack1 := stack[string]{
	elements: []string{"ali","sabit","sheik"},
   }

   fmt.Println(myStack1) //{[ali sabit sheik]}

}
