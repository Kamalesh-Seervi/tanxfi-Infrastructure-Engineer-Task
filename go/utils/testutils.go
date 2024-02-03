package utils

import (
	"encoding/csv"
	"os"
	"time"

	"github.com/kamalesh-seervi/tanxfi/models"
)

func ParseCSV(filePath string) ([]models.Order, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var orders []models.Order
	if len(records) > 0 {
		records = records[1:]
	}
	// Parse each record and create Order structs
	for _, record := range records {
		order := models.Order{
			OrderID:      StringtoInt(record[0]),
			CustomerID:   StringtoInt(record[1]),
			ProductID:    StringtoInt(record[3]),
			ProductName:  record[4],
			Quantity:     StringtoInt(record[6]),
			ProductPrice: StringtoFloat(record[5]),
		}
		// Parse order date
		orderDate, err := time.Parse("2006-01-02", record[2])
		if err != nil {
			return nil, err
		}
		order.OrderDate = orderDate

		orders = append(orders, order)
	}

	return orders, nil
}
func CalculateRevenue(orders []models.Order) (map[string]float64, map[string]float64, map[int]float64) {
	// Create maps to store monthly, product, and customer revenue
	monthlyRevenue := make(map[string]float64)
	productRevenue := make(map[string]float64)
	customerRevenue := make(map[int]float64)

	// Iterate over orders and calculate revenue
	for _, order := range orders {
		// Compute monthly revenue
		month := order.OrderDate.Format("2006-01")
		monthlyRevenue[month] += order.ProductPrice * float64(order.Quantity)

		// Compute product revenue
		productRevenue[order.ProductName] += order.ProductPrice * float64(order.Quantity)

		// Compute customer revenue
		customerRevenue[order.CustomerID] += order.ProductPrice * float64(order.Quantity)
	}

	return monthlyRevenue, productRevenue, customerRevenue
}
