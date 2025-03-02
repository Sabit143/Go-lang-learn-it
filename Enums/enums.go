package main

import "fmt"

// enumerated types

type OrderStatus int

// Here we have created a type called OrderStatus , and gave the type as int 
// here iota is only for string

type UserRoles string 

const(
	ADMIN UserRoles = "admin"
	USER            = "normal user"
	SUPERADMIN      = "super admin"
)

const (
	Recived OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)
func findRoles(roles UserRoles){
	fmt.Println("user role is ",roles)
}

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to",status)
}
func main() {
    changeOrderStatus(Prepared)
	findRoles(SUPERADMIN)

}
