package ImpliedVolatility

import (
	"log"
	"testing"
	"time"

	b "v2/src/ImpliedVolatility/BlackScholes"
	p "v2/src/Plot"
	t "v2/src/Tradier"
)

func TestImpliedVolatility(test *testing.T) {

	ticker_symbol := "NVDA"

	api_key := ""
	account_id := ""

	// Initialize Tradier Client
	client := t.NewTradierClient(api_key, account_id)

	expiry, err := client.GetExpirations(ticker_symbol, "false", "false")

	if err != nil {
		log.Println("Error Fetching Expirations")
	}

	var OC []t.OptionsChain

	for j := 0; j < len(expiry.Expirations.Dates); j++ {

		oc, err := client.GetOptions(ticker_symbol, expiry.Expirations.Dates[j], "false")

		if err != nil {
			log.Println("Error Fetching Option Chain Expiration")
		}

		OC = append(OC, oc)

	}

	quote_resp, _ := client.GetQuote(ticker_symbol)
	stock_price := quote_resp.Quotes.Quote.Last

	var P b.BlackScholesParameters
	P.Stock = stock_price

	var approximation [][]float64
	var estimation [][]float64
	var implied_vol [][]float64

	// Compute Implied Volatility
	for i := 0; i < len(OC); i++ {

		P.Expiry = b.GetDaysBetween(time.Now(), expiry.Expirations.Dates[i]) / 365.0
		var temp []float64
		var a []float64
		var g []float64

		for j := 0; j < len(OC[i].Option.Options); j++ {

			if OC[i].Option.Options[j].Strike < stock_price {

				if OC[i].Option.Options[j].OptionType == "put" {

					g = append(g, getManasterKoehler(P))

					P.Strike = OC[i].Option.Options[j].Strike

					bid := OC[i].Option.Options[j].Bid
					ask := OC[i].Option.Options[j].Ask
					mid := (bid + ask) / 2.0

					ivol := GetImpliedVolatility(P, mid, "put")
					temp = append(temp, ivol)

					// a = append(a, getDanStefanica(P, mid))

				}

			}

			if OC[i].Option.Options[j].Strike >= stock_price {

				if OC[i].Option.Options[j].OptionType == "call" {

					g = append(g, getManasterKoehler(P))

					P.Strike = OC[i].Option.Options[j].Strike

					bid := OC[i].Option.Options[j].Bid
					ask := OC[i].Option.Options[j].Ask
					mid := (bid + ask) / 2.0

					ivol := GetImpliedVolatility(P, mid, "call")
					temp = append(temp, ivol)

					a = append(a, getDanStefanica(P, mid))

				}

			}

		}

		implied_vol = append(implied_vol, temp)
		estimation = append(estimation, g)
		approximation = append(approximation, a)

	}

	// Visualize Volatility Curve

	p.CreateTimeSeriesPlot(approximation, "Strike", "Implied Volatility", "Implied Volatility: "+ticker_symbol, "vol_approx")
	p.CreateTimeSeriesPlot(estimation, "Strike", "Implied Volatility", "Implied Volatility: "+ticker_symbol, "vol_est")
	p.CreateTimeSeriesPlot(implied_vol, "Strike", "Implied Volatility", "Implied Volatility: "+ticker_symbol, "vol_curve")

}
