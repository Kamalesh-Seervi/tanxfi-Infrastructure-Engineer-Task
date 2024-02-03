package utils

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/kamalesh-seervi/tanxfi/models"
)

func StringtoInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	return i
}

func StringtoFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return 0
	}
	return f
}

func SortOrdersRevDesc(orders []models.Order) {
	sort.Slice(orders, func(i, j int) bool {
		revenue1 := orders[i].ProductPrice * float64(orders[i].Quantity)
		revenue2 := orders[j].ProductPrice * float64(orders[j].Quantity)
		return revenue1 > revenue2
	})

}

