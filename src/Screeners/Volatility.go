package Screener

import (
	"fmt"
	"log"
	"math"
	"time"
	iv "v2/src/ImpliedVolatility"
	b "v2/src/ImpliedVolatility/BlackScholes"
	m "v2/src/Mongo"
	t "v2/src/Tradier"
	v "v2/src/Volatility"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetVolatilityMetrics(client *t.TradierClient, mongo *mongo.Client) {

	// Iterate Thru Stock Tickers
	ticker_symbols := m.GetStockTickers(true)

	count := 0
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {

		// Initialize Objects
		var VMP v.VolatilityMethodsParameters
		var VAD m.VolArbitrageData
		var BSP b.BlackScholesParameters

		fmt.Println("Ticker Symbol: ", ticker_symbols[count])
		fmt.Println("")

		/*
			- Fetch Historical Data
			- Compute Yang Zhang Volatility
		*/

		historical, err := client.GetHistoricalData(ticker_symbols[count])

		if err != nil {
			log.Println("Error Fetching Historical Data")
		}

		length := len(historical.History.Day)

		if length > 0 {

			for j := 0; j < length; j++ {

				VMP.Open = append(VMP.Open, historical.History.Day[j].Open)
				VMP.High = append(VMP.High, historical.History.Day[j].High)
				VMP.Low = append(VMP.Low, historical.History.Day[j].Low)
				VMP.Close = append(VMP.Close, historical.History.Day[j].Close)

			}

			// Store Historical Data For Correlation Metrics
			m.AppendHistoricalMongo(mongo, VMP, ticker_symbols[count])

			// Set Stock Price
			BSP.Stock = historical.History.Day[length-1].Close

			if length > 120 {

				VMP.RollingPeriod = 30
				VAD.HV.HV30 = v.GetVolatilityMethods(VMP).YangZhang

				VMP.RollingPeriod = 60
				VAD.HV.HV60 = v.GetVolatilityMethods(VMP).YangZhang

				VMP.RollingPeriod = 90
				VAD.HV.HV90 = v.GetVolatilityMethods(VMP).YangZhang

				VMP.RollingPeriod = 120
				VAD.HV.HV120 = v.GetVolatilityMethods(VMP).YangZhang

			}

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

			expiry, err := client.GetExpirations(ticker_symbols[count], "false", "false")

			if err != nil {
				log.Println("Error Fetching Expirations")
			}

			/*
				Not Every Stock has Options
				Check For Length
			*/

			if len(expiry.Expirations.Dates) > 0 {

				var dte []float64

				// Compute Days to Expiry, Annualized

				for j := 0; j < len(expiry.Expirations.Dates); j++ {

					tau := b.GetDaysBetween(time.Now(), expiry.Expirations.Dates[j]) / 365.0
					dte = append(dte, tau)

				}

				fmt.Println("DTE: ", dte)
				fmt.Println("")

				var OC []t.OptionsChain

				for j := 0; j < len(dte); j++ {

					oc, err := client.GetOptions(ticker_symbols[count], expiry.Expirations.Dates[j], "false")

					if err != nil {
						log.Println("Error Fetching Option Chain Expiration")
					}

					OC = append(OC, oc)
				}

				var atm_iv []float64

				// Compute ATM Implied Volatility

				for j := 0; j < len(OC); j++ {

					BSP.Expiry = dte[j]
					var atm_price float64

					// Find ATM Price
					for k := 0; k < (len(OC[j].Option.Options) - 1); k++ {

						if OC[j].Option.Options[k].OptionType == "call" {

							if (OC[j].Option.Options[k].Strike < BSP.Stock) && (OC[j].Option.Options[k+2].Strike >= BSP.Stock) {

								atm_price = (OC[j].Option.Options[k+2].Bid + OC[j].Option.Options[k+2].Ask) / 2.0
								BSP.Strike = OC[j].Option.Options[k+2].Strike

							}

						}

					}

					ivol := iv.GetImpliedVolatility(BSP, atm_price, "Call")
					atm_iv = append(atm_iv, ivol)

				}

				fmt.Println("ATM IV: ", atm_iv)
				fmt.Println("")

				// Interpolate Implied Volatility

				var interpolated30 float64
				var interpolated60 float64
				var interpolated90 float64
				var interpolated120 float64

				var expiry_slice [4]int
				var dte_slice [4]float64

				for j := 0; j < (len(atm_iv) - 1); j++ {

					if dte[j] < (30/365.0) && dte[j+1] >= (30/365.0) {

						x := 30.0 / 365.0

						x0 := dte[j]
						x1 := dte[j+1]

						y0 := atm_iv[j]
						y1 := atm_iv[j+1]

						interpolated30 = y0 + ((x - x0) * ((y1 - y0) / (x1 - x0)))
						expiry_slice[0] = j
						dte_slice[0] = dte[j]

					}

					if dte[j] < (60/365.0) && dte[j+1] >= (60/365.0) {

						x := 60.0 / 365.0

						x0 := dte[j]
						x1 := dte[j+1]

						y0 := atm_iv[j]
						y1 := atm_iv[j+1]

						interpolated60 = y0 + ((x - x0) * ((y1 - y0) / (x1 - x0)))
						expiry_slice[1] = j
						dte_slice[1] = dte[j]

					}

					if dte[j] < (90/365.0) && dte[j+1] >= (90/365.0) {

						x := 90.0 / 365.0

						x0 := dte[j]
						x1 := dte[j+1]

						y0 := atm_iv[j]
						y1 := atm_iv[j+1]

						interpolated90 = y0 + ((x - x0) * ((y1 - y0) / (x1 - x0)))
						expiry_slice[2] = j
						dte_slice[2] = dte[j]

					}

					if dte[j] < (120/365.0) && dte[j+1] >= (120/365.0) {

						x := 120.0 / 365.0

						x0 := dte[j]
						x1 := dte[j+1]

						y0 := atm_iv[j]
						y1 := atm_iv[j+1]

						interpolated120 = y0 + ((x - x0) * ((y1 - y0) / (x1 - x0)))
						expiry_slice[3] = j
						dte_slice[3] = dte[j]

					}

				}

				VAD.IV.IV30 = interpolated30
				VAD.IV.IV60 = interpolated60
				VAD.IV.IV90 = interpolated90
				VAD.IV.IV120 = interpolated120

				fmt.Println("Implied Volatility Metrics")
				fmt.Println("IV30: ", interpolated30)
				fmt.Println("IV60: ", interpolated60)
				fmt.Println("IV90: ", interpolated90)
				fmt.Println("IV120: ", interpolated120)
				fmt.Println("")

				VAD.VRP.VRP30 = interpolated30 - VAD.HV.HV30[len(VAD.HV.HV30)-1]
				VAD.VRP.VRP60 = interpolated60 - VAD.HV.HV60[len(VAD.HV.HV60)-1]
				VAD.VRP.VRP90 = interpolated90 - VAD.HV.HV90[len(VAD.HV.HV90)-1]
				VAD.VRP.VRP120 = interpolated120 - VAD.HV.HV120[len(VAD.HV.HV120)-1]

				fmt.Println("Variance Risk Premium Metrics")
				fmt.Println("VRP30: ", VAD.VRP.VRP30)
				fmt.Println("VRP60: ", VAD.VRP.VRP60)
				fmt.Println("VRP90: ", VAD.VRP.VRP90)
				fmt.Println("VRP120: ", VAD.VRP.VRP120)
				fmt.Println("")

				/*
					- Compute Expected Move Metrics
					- Fetch Skew Data
				*/

				var expected_move [4]float64
				expected_move[0] = BSP.Stock * interpolated30 * math.Sqrt(30/365.0)
				VAD.EM.EM30 = expected_move[0]
				expected_move[1] = BSP.Stock * interpolated60 * math.Sqrt(60/365.0)
				VAD.EM.EM60 = expected_move[1]
				expected_move[2] = BSP.Stock * interpolated90 * math.Sqrt(90/365.0)
				VAD.EM.EM90 = expected_move[2]
				expected_move[3] = BSP.Stock * interpolated120 * math.Sqrt(120/365.0)
				VAD.EM.EM120 = expected_move[3]

				fmt.Println("Expected Move Metrics")
				fmt.Println("EM30: ", expected_move[0])
				fmt.Println("EM60: ", expected_move[1])
				fmt.Println("EM90: ", expected_move[2])
				fmt.Println("EM120: ", expected_move[3])
				fmt.Println("")

				/*
					- Find Strike With Least Expected Move Residual
				*/

				var strikes [4][2]float64
				var tail_prices [4][2]float64

				for j := 0; j < len(expiry_slice); j++ {

					index := expiry_slice[j]

					var put_tail bool
					var call_tail bool
					var temp_strikes [2]float64
					var temp_tails [2]float64

					for k := 0; k < (len(OC[index].Option.Options) - 2); k++ {

						lower_expectation := BSP.Stock - expected_move[0]

						if OC[index].Option.Options[k].OptionType == "put" && !put_tail {

							lower := OC[index].Option.Options[k].Strike
							upper := OC[index].Option.Options[k+2].Strike

							if lower < lower_expectation && upper >= lower_expectation {

								put_tail = true
								temp_strikes[0] = OC[index].Option.Options[k].Strike
								temp_tails[0] = (OC[index].Option.Options[k].Bid + OC[index].Option.Options[k].Ask) / 2.0

							}

						}

						upper_expectation := BSP.Stock + expected_move[0]

						if OC[index].Option.Options[k].OptionType == "call" && !call_tail {

							lower := OC[index].Option.Options[k].Strike
							upper := OC[index].Option.Options[k+2].Strike

							if lower < upper_expectation && upper >= upper_expectation {

								call_tail = true
								temp_strikes[1] = OC[index].Option.Options[k].Strike
								temp_tails[1] = (OC[index].Option.Options[k].Bid + OC[index].Option.Options[k].Ask) / 2.0
								break

							}

						}

					}

					strikes[j] = temp_strikes
					tail_prices[j] = temp_tails

				}

				fmt.Println("Expiry Slice: ", expiry_slice)
				fmt.Println("")

				fmt.Println("DTE Slice: ", dte_slice)
				fmt.Println("")

				fmt.Println("Tail Strikes: ", strikes)
				fmt.Println("")

				fmt.Println("Tail Prices", tail_prices)
				fmt.Println("")

				/*
					Compute Implied Volatility of Tails
				*/

				// 30 Day Tails
				BSP.Expiry = dte_slice[0]
				BSP.Strike = strikes[0][0]
				VAD.PutIV.IV30 = iv.GetImpliedVolatility(BSP, tail_prices[0][0], "Put")
				BSP.Strike = strikes[0][1]
				VAD.CallIV.IV30 = iv.GetImpliedVolatility(BSP, tail_prices[0][1], "Call")

				// 60 Day Tails
				BSP.Expiry = dte_slice[1]
				BSP.Strike = strikes[1][0]
				VAD.PutIV.IV60 = iv.GetImpliedVolatility(BSP, tail_prices[1][0], "Put")
				BSP.Strike = strikes[1][1]
				VAD.CallIV.IV60 = iv.GetImpliedVolatility(BSP, tail_prices[1][1], "Call")

				// 90 Day Tails
				BSP.Expiry = dte_slice[2]
				BSP.Strike = strikes[2][0]
				VAD.PutIV.IV90 = iv.GetImpliedVolatility(BSP, tail_prices[2][0], "Put")
				BSP.Strike = strikes[2][1]
				VAD.CallIV.IV90 = iv.GetImpliedVolatility(BSP, tail_prices[2][1], "Call")

				// 120 Day Tails
				BSP.Expiry = dte_slice[3]
				BSP.Strike = strikes[3][0]
				VAD.PutIV.IV120 = iv.GetImpliedVolatility(BSP, tail_prices[3][0], "Put")
				BSP.Strike = strikes[3][1]
				VAD.CallIV.IV120 = iv.GetImpliedVolatility(BSP, tail_prices[3][1], "Call")

				// Additional Information
				VAD.PutIV.TailStrikes = append(VAD.PutIV.TailStrikes, strikes[0][0])
				VAD.PutIV.TailStrikes = append(VAD.PutIV.TailStrikes, strikes[1][0])
				VAD.PutIV.TailStrikes = append(VAD.PutIV.TailStrikes, strikes[2][0])
				VAD.PutIV.TailStrikes = append(VAD.PutIV.TailStrikes, strikes[3][0])

				VAD.CallIV.TailStrikes = append(VAD.CallIV.TailStrikes, strikes[0][1])
				VAD.CallIV.TailStrikes = append(VAD.CallIV.TailStrikes, strikes[1][1])
				VAD.CallIV.TailStrikes = append(VAD.CallIV.TailStrikes, strikes[2][1])
				VAD.CallIV.TailStrikes = append(VAD.CallIV.TailStrikes, strikes[3][1])

				fmt.Println("Tail Metrics")
				fmt.Println("Put Tail 30: ", VAD.PutIV.IV30)
				fmt.Println("Call Tail 30: ", VAD.CallIV.IV30)
				fmt.Println("Put Tail 60: ", VAD.PutIV.IV60)
				fmt.Println("Call Tail 60: ", VAD.CallIV.IV60)
				fmt.Println("Put Tail 90: ", VAD.PutIV.IV90)
				fmt.Println("Call Tail 90: ", VAD.CallIV.IV90)
				fmt.Println("Put Tail 120: ", VAD.PutIV.IV120)
				fmt.Println("Call Tail 120: ", VAD.CallIV.IV120)
				fmt.Println("")

				m.AppendVolatilityMongo(mongo, VAD, ticker_symbols[count])
				fmt.Println("Appending to MongoDB")
				fmt.Println("")

			}

		}

		count++

		if count == len(ticker_symbols) {
			return
		}

	}

}
