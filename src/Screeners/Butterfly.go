package Screener

import (
	"fmt"
	"log"
	"math"
	"time"
	roots "v2/src/ImpliedVolatility"
	b "v2/src/ImpliedVolatility/BlackScholes"
	t "v2/src/Tradier"
)

/*
	This Method Computes the Price of All Butterfly Spreads

	Input:

	1. Short List and Scan for Trade Opportunities

	Note:

	1. Each Trade Opportunity Will Arise From Cheap Spread, or Negative Spread

	2. Each Trade Opportunity Represents Possible Liquidity Holes

	3. This Method Places Trades !!!
*/
func GetButterflyMetrics(client *t.TradierClient, top_tickers []string) {

	for _, ticker := range top_tickers {

		fmt.Println("Ticker Symbol: ", ticker)
		fmt.Println("")

		expiry, err := client.GetExpirations(ticker, "false", "false")

		if err != nil {
			log.Println("Error Fetching Expirations")
		}

		var OC []t.OptionsChain

		if len(expiry.Expirations.Dates) > 0 {

			for j := 0; j < len(expiry.Expirations.Dates); j++ {

				oc, err := client.GetOptions(ticker, expiry.Expirations.Dates[j], "false")

				if err != nil {
					log.Println("Error Fetching Option Chain Expiration")
				}

				OC = append(OC, oc)

			}

		}

		quote_resp, _ := client.GetQuote(ticker)
		stock_price := quote_resp.Quotes.Quote.Last

		var P b.BlackScholesParameters
		P.Stock = stock_price

		var implied_vol [][]float64

		if len(OC) > 0 {

			for i := 0; i < len(OC); i++ {

				P.Expiry = b.GetDaysBetween(time.Now(), expiry.Expirations.Dates[i]) / 365.0
				var temp []float64

				for j := 0; j < len(OC[i].Option.Options); j++ {

					if OC[i].Option.Options[j].Strike < stock_price {

						if OC[i].Option.Options[j].OptionType == "put" {

							P.Strike = OC[i].Option.Options[j].Strike

							bid := OC[i].Option.Options[j].Bid
							ask := OC[i].Option.Options[j].Ask
							mid := (bid + ask) / 2.0

							ivol := roots.GetImpliedVolatility(P, mid, "put")
							temp = append(temp, ivol)

						}

					}

					if OC[i].Option.Options[j].Strike >= stock_price {

						if OC[i].Option.Options[j].OptionType == "call" {

							P.Strike = OC[i].Option.Options[j].Strike

							bid := OC[i].Option.Options[j].Bid
							ask := OC[i].Option.Options[j].Ask
							mid := (bid + ask) / 2.0

							ivol := roots.GetImpliedVolatility(P, mid, "call")
							temp = append(temp, ivol)

						}

					}

				}

				implied_vol = append(implied_vol, temp)

			}

		}

		var deltas [][]float64

		for i := 0; i < len(implied_vol); i++ {

			P.Expiry = b.GetDaysBetween(time.Now(), expiry.Expirations.Dates[i]) / 365.0
			var temp []float64

			for j := 0; j < len(OC[i].Option.Options); j++ {

				if OC[i].Option.Options[j].Strike < stock_price {

					if OC[i].Option.Options[j].OptionType == "put" {

						P.Strike = OC[i].Option.Options[j].Strike

						delta := b.GetCallDelta(P)
						temp = append(temp, delta)

					}

				}

				if OC[i].Option.Options[j].Strike >= stock_price {

					if OC[i].Option.Options[j].OptionType == "call" {

						P.Strike = OC[i].Option.Options[j].Strike

						P.Strike = OC[i].Option.Options[j].Strike

						delta := b.GetPutDelta(P)
						temp = append(temp, delta)

					}

				}

			}

			deltas = append(deltas, temp)

		}

		// Fill Order Ticker
		var ticket t.OrderTicket
		ticket.Class = "multileg"
		ticket.Symbol = ticker
		ticket.Type = "limit"
		ticket.Duration = "day"

		for i := 0; i < len(deltas); i++ {

			for j := 0; j < (len(OC[i].Option.Options) - 2); j++ {

				if OC[i].Option.Options[j].Strike < stock_price {

					if OC[i].Option.Options[j].OptionType == "put" && deltas[i][j] < -0.30 && deltas[i][j+2] > -0.50 {

						// Price Butterfly Spread
						lower := (OC[i].Option.Options[j].Bid + OC[i].Option.Options[j].Ask) / 2.0
						k := (OC[i].Option.Options[j+1].Bid + OC[i].Option.Options[j+1].Ask) / 2.0
						upper := (OC[i].Option.Options[j+2].Bid + OC[i].Option.Options[j+2].Ask) / 2.0

						fly := lower - (2 * k) + upper
						lower_wing := math.Abs(OC[i].Option.Options[j+1].Strike - OC[i].Option.Options[j].Strike)
						upper_wing := math.Abs(OC[i].Option.Options[j+2].Strike - OC[i].Option.Options[j+1].Strike)

						// Possible Arbitrage Opportunity
						if fly < 0 && lower_wing == upper_wing {

							// Place Bid at 0.01
							ticket.Price = "0.01"

						}

					}

				}

				if OC[i].Option.Options[j].Strike >= stock_price {

					if OC[i].Option.Options[j].OptionType == "call" && deltas[i][j] < 0.50 && deltas[i][j+2] > 0.30 {

						// Price Butterfly Spread
						lower := (OC[i].Option.Options[j].Bid + OC[i].Option.Options[j].Ask) / 2.0
						k := (OC[i].Option.Options[j+1].Bid + OC[i].Option.Options[j+1].Ask) / 2.0
						upper := (OC[i].Option.Options[j+2].Bid + OC[i].Option.Options[j+2].Ask) / 2.0

						fly := lower - (2 * k) + upper
						lower_wing := math.Abs(OC[i].Option.Options[j+1].Strike - OC[i].Option.Options[j].Strike)
						upper_wing := math.Abs(OC[i].Option.Options[j+2].Strike - OC[i].Option.Options[j+1].Strike)

						// Possible Arbitrage Opportunity
						if fly < 0 && lower_wing == upper_wing {

							// Place Bid at 0.01
							ticket.Price = "0.01"

						}

					}

				}

			}

		}

	}

}
