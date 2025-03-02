package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5
	// simple switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}
	// multiple condition switch
      switch time.Now().Weekday(){
	  case time.Saturday, time.Sunday :
		fmt.Println("its a weekend")
	  default:
		fmt.Println("workday")
	  }


	  // type switch 
	  whoAmI := func(i interface{}){
		switch t := i.(type){
		case int:
			fmt.Println("its an Integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other",t)
		}

	  }

	  whoAmI(1)
	  whoAmI("hello my friend")
}
