package main

import (
	"fmt"
	"portfolio-rebalancer/internal/application"
	"portfolio-rebalancer/internal/domain"
)


func main() {

	portfolio := domain.Portfolio{
		Holdings: map[string]domain.Stock {
			"META": {Symbol: "META", Quantity: 10, CurrentPrice: 60},
			"AAPL": {Symbol: "AAPL", Quantity: 5, CurrentPrice: 80},
		},
		TargetAllocation: map[string]float64 {
			"META": 0.4,
			"AAPL": 0.6,
		},
	}

	service := application.RebalanceService{}	
	actions := service.Execute(portfolio)

	for _, a := range actions {
		fmt.Println(a)
	}
}


