package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/kamalesh-seervi/tanxfi/models"
	"github.com/kamalesh-seervi/tanxfi/utils"
)

func main() {
	var orders []models.Order

	// Opening CSV file using OS library
	file, err := os.Open("orders.csv")
	if err != nil {
		fmt.Println("error in opening csv:", err)
		return
	}
	defer file.Close()

	// Parse the CSV file for line by line reading
	reader := csv.NewReader(file)
	datas, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return
	}

	monthlyRevenue := make(map[string]float64)
	productRevenue := make(map[string]float64)
	customerRevenue := make(map[int]float64)

	// Parse each record and compute revenue and also conversion of string to int or float64
	for i, record := range datas {
		if i == 0 {
			continue
		}
		order := models.Order{
			OrderID:      utils.StringtoInt(record[0]),
			CustomerID:   utils.StringtoInt(record[1]),
			ProductID:    utils.StringtoInt(record[3]),
			ProductName:  record[4],
			Quantity:     utils.StringtoInt(record[6]),
			ProductPrice: utils.StringtoFloat(record[5]),
		}

		orderDate, err := time.Parse("2006-01-02", record[2])
		if err != nil {
			fmt.Println("Error parsing order date:", err)
			return
		}
		order.OrderDate = orderDate
		//  Add monthly revenue
		month := orderDate.Format("2006-01")
		monthlyRevenue[month] += order.ProductPrice * float64(order.Quantity)

		// Add products to map
		productRevenue[order.ProductName] += order.ProductPrice * float64(order.Quantity)

		//  Add entry into customerRevenue map if it doesn't exist otherwise add to existing value
		orderRevenue := order.ProductPrice * float64(order.Quantity)
		customerRevenue[order.CustomerID] += orderRevenue

		// Store the order
		orders = append(orders, order)
	}

	// Total Revenue by Month
	fmt.Println("Total Revenue by Month:")
	for month, revenue := range monthlyRevenue {
		fmt.Printf("%s: $%.2f\n", month, revenue)
	}

	// Total Revenue by Product:
	fmt.Println("\nTotal Revenue by Product:")
	for product, revenue := range productRevenue {
		fmt.Printf("%s: $%.2f\n", product, revenue)
	}

	// Total Revenue by Customer
	fmt.Println("\nTotal Revenue by Customer:")
	for customer, revenue := range customerRevenue {
		fmt.Printf("%d: $%.2f\n", customer, revenue)
	}

	// Top 10 Customers by Revenue
	utils.SortOrdersRevDesc(orders)
	fmt.Println("\nTop 10 Customers by Revenue:")
	for i, order := range orders[:10] {
		fmt.Printf("%d. Customer %d: $%.2f\n", i+1, order.CustomerID, customerRevenue[order.CustomerID])
	}
}
