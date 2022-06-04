package main

import (
	"fmt"
	"log"
	"time"
	b "v2/src/BlackScholes"
	m "v2/src/Mongo"
	t "v2/src/Tradier"
	v "v2/src/Volatility"
)

func main() {

	// Enter Api Key
	var api_key string
	fmt.Println("Enter Api Key: ")
	fmt.Scanln(&api_key)

	// Enter Account ID
	var account_id string
	fmt.Println("Enter Account ID: ")
	fmt.Scanln(&account_id)

	// Initialize Client
	client := t.NewTradierClient(api_key, account_id)

	// Iterate Thru Stock Tickers
	ticker_symbols := getStockTickers()

	// Initialize Objects
	var VMP v.VolatilityMethodsParameters
	var VAD m.VolArbitrageData
	var BSP b.BlackScholesParameters

	for i := 0; i < len(ticker_symbols); i++ {

		fmt.Println("Ticker Symbol: ", ticker_symbols[i])
		fmt.Println("")

		/*
			- Fetch Historical Data
			- Compute Yang Zhang Volatility
		*/

		historical, err := client.GetHistoricalData(ticker_symbols[i])

		if err != nil {
			log.Println("Error Fetching Historical Data")
		}

		length := len(historical.History.Day)

		for j := 0; j < length; j++ {

			VMP.OHLC.Open = append(VMP.OHLC.Open, historical.History.Day[i].Open)
			VMP.OHLC.High = append(VMP.OHLC.High, historical.History.Day[i].High)
			VMP.OHLC.Low = append(VMP.OHLC.Low, historical.History.Day[i].Low)
			VMP.OHLC.Close = append(VMP.OHLC.Close, historical.History.Day[i].Close)

		}

		// Set Stock Price
		BSP.Stock = historical.History.Day[length-1].Close

		VMP.RollingPeriod = 30
		VAD.HV.HV30 = v.GetVolatilityMethods(VMP).YangZhang[(length - VMP.RollingPeriod - 3)]

		VMP.RollingPeriod = 60
		VAD.HV.HV60 = v.GetVolatilityMethods(VMP).YangZhang[(length - VMP.RollingPeriod - 3)]

		VMP.RollingPeriod = 90
		VAD.HV.HV90 = v.GetVolatilityMethods(VMP).YangZhang[(length - VMP.RollingPeriod - 3)]

		VMP.RollingPeriod = 120
		VAD.HV.HV120 = v.GetVolatilityMethods(VMP).YangZhang[(length - VMP.RollingPeriod - 3)]

		fmt.Println("Historical Volatility Metrics")
		fmt.Println("HV30: ", VAD.HV.HV30)
		fmt.Println("HV60: ", VAD.HV.HV60)
		fmt.Println("HV90: ", VAD.HV.HV90)
		fmt.Println("HV120: ", VAD.HV.HV120)
		fmt.Println("")

		/*
			- Fetch Expirations
			- Fetch Option Chain
			- Compute Implied Volatility via Linear Interpolation
		*/

		expiry, err := client.GetExpirations(ticker_symbols[i], "false", "false")

		if err != nil {
			log.Println("Error Fetching Expirations")
		}

		var dte []float64

		// Compute Days to Expiry, Annualized

		for j := 0; j < len(expiry.Expirations.Dates); j++ {

			tau := b.GetDaysBetween(time.Now(), expiry.Expirations.Dates[i])
			tau /= 365.0
			dte = append(dte, tau)

		}

		oc, err := client.GetOptions(ticker_symbols[i], "false", "false")

		if err != nil {
			log.Println("Error Fetching Option Chain")
		}

		var atm_iv []float64

		// Compute ATM Implied Volatility

		for j := 0; j < len(dte); j++ {

			BSP.Expiry = dte[j]
			var atm_price float64

			// Find ATM Price
			for k := 0; k < (len(oc.Option.Options) - 1); k++ {

				if (oc.Option.Options[j].Strike < BSP.Stock) && (oc.Option.Options[j+1].Strike >= BSP.Stock) {
					atm_price = (oc.Option.Options[j+1].Bid + oc.Option.Options[j+1].Ask) / 2.0
					BSP.Strike = oc.Option.Options[j+1].Strike
				}

			}

			iv := b.GetImpliedVolatility(BSP, atm_price, "Call")
			atm_iv = append(atm_iv, iv)

		}

		fmt.Println("ATM IV: ", atm_iv)

	}

}
