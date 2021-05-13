package order

import (
	"time"
)

type Address struct {
	Country string
	City 	string
	Street  string
	NumberHome int
}

type Order struct {
	OrderId			string
	CreatedTime 	time.Time
	OrderNumber		uint64
	CustomerName 	string
	PhoneNumber 	string
	Address 		Address
	NodeForOrder 	string
	ItemsOrders 	string
	Link 			string
}


