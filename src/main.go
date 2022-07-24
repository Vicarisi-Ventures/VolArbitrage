package main

import (
	"fmt"
	m "v2/src/Mongo"
	s "v2/src/Screeners"
	t "v2/src/Tradier"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	FE "v2/src/FrontEnd"
)

var arr []string

func main() {

	// runScreeners()
	// os.Exit(0)

	a := app.New()
	w := a.NewWindow("Vicarisi Ventures: Volatility Arbitrage")
	w.CenterOnScreen()
	w.Padded()

	r, _ := FE.LoadResourceFromPath("C:/Users/18627/Documents/GitHub/VolArbitrage/Icon.png")
	w.SetIcon(r)

	c1 := make(chan *fyne.Container, 1)
	c2 := make(chan []string, 1)

	go FE.LoginPanel(c1, c2)

	initial_content := <-c1
	w.SetContent(initial_content)

	arr = <-c2

	fmt.Println("Starting")
	arbitrage_panel := FE.ScreenerPanel()
	contango_panel := FE.ContangoPanel()
	simulation_panel := FE.SimulationPanel()
	volatility_panel := FE.ImpliedVolatilityPanel()

	// App Tabs
	tabs := container.NewAppTabs(

		container.NewTabItem("Arbitrage Screener", container.NewVBox(

			arbitrage_panel,
		)),

		container.NewTabItem("Contango Backtest Engine", container.NewVBox(

			contango_panel,
		)),

		container.NewTabItem("Stochastic Simulation Engine", container.NewVBox(

			simulation_panel,
		)),

		container.NewTabItem("Volatility Curve", container.NewVBox(

			volatility_panel,
		)),
	)
	fmt.Println("Done")

	w.SetContent(tabs)
	w.ShowAndRun()

}

func runScreeners() {

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
	// client := t.NewTradierClient(arr[0], arr[1])
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
