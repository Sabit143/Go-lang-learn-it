package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task",id)
}
func main() {
   for i :=0 ; i <= 10; i++{
	go task(i)

	go func(i int){
		fmt.Println(i)
	   }(i)   
    // This is a anoymous function 
	
   }
   
   time.Sleep(time.Second*1)  // here iam making main thread to sleep for a moment like 1second 
   // these will run concurrently , its working on multithreaded 
}
