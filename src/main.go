package main

import (
	"fmt"

	m "v2/src/Mongo"
	s "v2/src/Screeners"
	t "v2/src/Tradier"
)

func main() {

	// Enter Api Key
	var api_key string
	fmt.Println("Enter Api Key: ")
	fmt.Scanln(&api_key)
	fmt.Println("")

	// Enter Account ID
	var account_id string
	fmt.Println("Enter Account ID: ")
	fmt.Scanln(&account_id)
	fmt.Println("")

	// Initialize Tradier Client
	client := t.NewTradierClient(api_key, account_id)

	// Initialize Mongo Client
	mongo := m.GetMongoConnection()

	// Compute Volatility Metrics
	s.GetVolatilityMetrics(client, mongo)

	// Compute Correlation Metrics
	s.GetCorrelationMetrics(mongo)

	// Fetch Top Five Variance Risk Premium
	var vrp_tickers []string

	// Fetch Top Five Call Skew

	// Fetch Top Five Put Skew

	// Run Tickers Through Butterfly Screener
	s.GetButterflyMetrics(client, vrp_tickers)

}
