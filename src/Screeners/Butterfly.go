package Screener

import (
	"fmt"
	"log"
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

	for i := 0; i < len(top_tickers); i++ {

		fmt.Println("Ticker Symbol: ", top_tickers[i])
		fmt.Println("")

		expiry, err := client.GetExpirations(top_tickers[i], "false", "false")

		if err != nil {
			log.Println("Error Fetching Expirations")
		}

		var OC []t.OptionsChain

		if len(expiry.Expirations.Dates) > 0 {

			for j := 0; j < len(expiry.Expirations.Dates); j++ {

				go func() {

					oc, err := client.GetOptions(top_tickers[i], expiry.Expirations.Dates[j], "false")

					if err != nil {
						log.Println("Error Fetching Option Chain Expiration")
					}

					OC = append(OC, oc)

				}()

			}

		}

		if len(OC) > 0 {

			go func() {

			}()

		}

	}

}
