package main

import "fmt"

// for -> only construct in go for looping

func main(){
	//while loop kind of function 
	// i := 1
	// for i <=3 {
	// 	fmt.Println(i)
	// 	i = i+1
	// }

	// //infinite loop 
	// for{
	// 	fmt.Println("1")
	// }

	//classic for loop 
	for j := 0 ; j < 3 ; j++ {
		if j==2 {
			continue
		}
		fmt.Println(j)

		// a new feature came range 

		for i := range 6 {
			fmt.Println(i)
		}
		// Here it will print 1, 2, 3, 4, 5

		for a := range 7{
			if a==6{
				continue
			}
			fmt.Println("hi my name is sabit-",a)
		}

		for b :=7 ; b>0 ; b--{
			fmt.Println("print the number")
			fmt.Println(b);
			
		}
}





}