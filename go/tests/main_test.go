package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/kamalesh-seervi/tanxfi/models"
	"github.com/kamalesh-seervi/tanxfi/utils"
)

func TestCSVParser(t *testing.T) {

	orders, err := utils.ParseCSV("orders.csv")
	if err != nil {
		t.Errorf("Error parsing CSV: %v", err)
	}
	if len(orders) != 25 {
		t.Fatalf("Expected 25 orders, got %d", len(orders))
	}

	// the first order detials check
	firstOrder := orders[0]
	expectedDate := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	expectedOrder := models.Order{
		OrderID:      1,
		CustomerID:   101,
		ProductID:    1,
		ProductName:  "PRODUCT A",
		Quantity:     2,
		ProductPrice: 250,
		OrderDate:    expectedDate,
	}

	if !reflect.DeepEqual(firstOrder, expectedOrder) {
		t.Fatalf("Incorrect details for the first order")
	}
}

func TestRevenueCalculation(t *testing.T) {
	// Mock orders for testing revenue calculation
		orders := []models.Order{
			{
				OrderID:      1,
				CustomerID:   101,
				OrderDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
				ProductID:    1,
				ProductName:  "Product A",
				Quantity:     2,
				ProductPrice: 250.0,
			},
	
	}

	monthlyRevenue, productRevenue, customerRevenue := utils.CalculateRevenue(orders)


	if monthlyRevenue["2024-01"] != 500.0 { // 2 * 250 = 500
		t.Errorf("Incorrect monthly revenue")
	}

	if productRevenue["Product A"] != 500.0 { // 2 * 250 = 500
		t.Errorf("Incorrect product revenue for Product A")
	}

	if customerRevenue[101] != 500.0 { // 2 * 250 = 500
		t.Errorf("Incorrect customer revenue for Customer 101")
	}
}
