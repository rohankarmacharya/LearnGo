package main

import (
	"fmt"
	"time"
)

// there's no class in Go, instead we have structs
// structs are used to group related data together

//order structs

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

func (o *order) getStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

func main() {
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 99.99,
	// 	status: "received",
	// }

	// myOrder2 := order{
	// 	id:     "2",
	// 	amount: 150.50,
	// 	status: "pending",
	// }

	// myOrder.status = "paid"

	// fmt.Println("My order:", myOrder)
	// fmt.Println("Order ID:", myOrder.id)

	// fmt.Println("My order 2:", myOrder2)

	// myOrder.getStatus("shipped")
	// fmt.Println("Updated order status:", myOrder.getAmount())

	myOrder := newOrder("1", 99.99, "received")
	fmt.Println(myOrder)

}
