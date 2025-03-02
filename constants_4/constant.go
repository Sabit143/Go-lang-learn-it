package main

import "fmt"

func main() {
	const name string = "ali"
	// name = "sabit" we cannot do this , as we assigned this as constant 
	fmt.Println(name)

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)
}

