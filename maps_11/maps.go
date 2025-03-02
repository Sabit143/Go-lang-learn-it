package main

import (
	"fmt"
	"maps"
)

func main() {
	// creaing map
	// m := make(map[string]string)

	// // setting an element
	// m["name"] = "golang"
	// m["area"] = "bacckend"

	// // get an element
	// fmt.Println(m["name"])
	// fmt.Println(m["ali"])  // it will give the empty value "", zero value 

	// // suppsoe i wantt to get the length
	// fmt.Println(len(m))

	// delete(m,"name")
	// clear(m)
	// fmt.Print(m)

	// how to create the map normally 
	 m := map[string]int{"price":49,"phone":3}
	// fmt.Println(m)
    v, ok := m["price"]
       fmt.Println(v)
	if ok {
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok") 
	}
     

	m1 := map[string]int{"price":49, "phones":4}
	m2 := map[string]int{"price":49, "phones":4}
       // for maps we have map package
	fmt.Println(maps.Equal(m1,m2))
}
