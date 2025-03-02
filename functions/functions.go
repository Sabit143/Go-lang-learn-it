package main



// func add(a int, b int) int {
// 	return a + b
// }

//suppose we have a function add which takes the parameter same data type , we can write like this also
func add(a, b int) int {
	return a + b
}

func getLanguage() (string, string, string) {
	return "golang ", "javascript", "c"
}

// for supressing any variable can be done by _ like this

func main() {
	// result := add(3, 5)
	// fmt.Println(result)
	// a,b,c := getLanguage()
	// fmt.Println(a,b,c)

	// Here we made a function and stored in the variable
	fn := func(a int) int {
		return 2
	}
	processIt()
     
	process(fn)

}

// This is a function which recives another function
func process(fn func(a int) int) {
	fn(1)
}

func processIt() func(a int) int{
	return func(a int) int{
            return 4
	}
}
