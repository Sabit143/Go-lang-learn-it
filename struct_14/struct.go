package main

import (
	"fmt"
	"time"
)

// suppose we want to make the classes , thats the equivalent

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nonosecond precision
}

// in other programming language we have constructor , but in go we create function using the new /New
func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// By using the above struct we can make the multiple copies of instance
func (o *order) changeStatus(status string) {
	o.status = status
}
// Here we can see that we are passing the reference by attaching the function
func (o *order) getAmount() float32 {
	return o.amount
}

// in go there is a convention , we shall use the * when we want to modify the value , otherwise , we can use normal
func main() {

	myOrder := newOrder("1", 30.50, "recived")

	fmt.Println(myOrder)
    // This is kind of struct we can make 
	language := struct{
		name string
		isGood bool
	}{"golang",true}

	fmt.Println(language)
	// like we create the instance of class , we can create the instances of struct too

	//  myOrder.createdAt = time.Now()

	//  fmt.Println(myOrder.status)
	//  myOrder.changeStatus("confirmed")

	//   amount :=myOrder.getAmount()
	//   fmt.Println(amount)
	//  fmt.Println(myOrder)

	// if we dont set the field of any struct , it sets the zero value , for string its "" and for boolean flase
}
