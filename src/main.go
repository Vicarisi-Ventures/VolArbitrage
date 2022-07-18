package main

import (
	"log"

	m "v2/src/Mongo"
	s "v2/src/Screeners"
	t "v2/src/Tradier"

	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Vicarisi Ventures: Volatility Arbitrage")
	w.CenterOnScreen()
	w.Padded()

	// Enter Api Key
	var api_key string
	api_label := widget.NewLabel("Enter Api Key: ")
	api_input := widget.NewEntry()
	api_input.SetPlaceHolder("Enter text...")

	// Enter Account ID
	var account_id string
	account_label := widget.NewLabel("Enter Account ID: ")
	account_input := widget.NewEntry()
	account_input.SetPlaceHolder("Enter text...")

	initial_content := container.NewVBox(

		api_label,

		api_input, widget.NewButton("Enter",
			func() {

				api_key = api_input.Text
				log.Println("Api Key:", api_input.Text)

			}),

		account_label,

		account_input, widget.NewButton("Enter",
			func() {

				account_id = account_input.Text
				log.Println("Account ID:", account_input.Text)

			}),
	)

	w.SetContent(initial_content)

	secondary_content := container.NewVBox()

	w.SetContent(secondary_content)

	w.ShowAndRun()

	// Initialize Tradier Client
	client := t.NewTradierClient(api_key, account_id)

	// Initialize Mongo Client
	mongo := m.GetMongoConnection()

	// Compute Volatility Metrics
	s.GetVolatilityMetrics(client, mongo)

	// Compute Correlation Metrics
	s.GetCorrelationMetrics(mongo)

	// Fetch Top Five Variance Risk Premium
	// var vrp_tickers []string

	// Fetch Top Five Call Skew

	// Fetch Top Five Put Skew

	// Run Tickers Through Butterfly Screener
	// s.GetButterflyMetrics(client, vrp_tickers)

}
