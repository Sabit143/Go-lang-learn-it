package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time   //nanosecond
	customer
}

func main() {
	newCustomer := customer{
		name:"ali",
		phone:"89867867868768",
	}
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "recived",
		customer: newCustomer,   // we can add the inline too cutsomer{name:"ali", phone:"89867867868768",}
	}

	fmt.Println(newOrder)
	fmt.Println(newOrder.customer)
}
