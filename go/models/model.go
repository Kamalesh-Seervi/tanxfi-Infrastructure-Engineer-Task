package models

import "time"

type Order struct {
	OrderID      int
	CustomerID   int
	OrderDate    time.Time
	ProductID    int
	ProductName  string
	ProductPrice float64
	Quantity     int
}

type CustomerRevenue struct {
	CustomerID   int
	TotalRevenue float64
}
