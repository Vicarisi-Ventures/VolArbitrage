package VIX

import (
	"fmt"
	"log"
	"testing"

	p "v2/src/Plot"
	t "v2/src/Tradier"
)

func TestVIX(test *testing.T) {

	api_key := ""
	account_id := ""

	// Initialize Tradier Client
	client := t.NewTradierClient(api_key, account_id)

	vix_hist, err := client.GetHistoricalData("VIX")

	if err != nil {
		log.Println("Error Fetching Historical Data")
	}

	vix_front_month_hist := FetchVixFuturesHistorical()

	var contango []float64

	// Relative Dynamics Between Spot VIX and Front Month Futures Contract
	for i := 0; i < len(vix_front_month_hist); i++ {

		spot := vix_hist.History.Day[i].Close
		front := vix_front_month_hist[i].Close

		risk_premia := (front - spot) / spot * 100
		contango = append(contango, risk_premia)

	}

	// Visualize Data
	p.CreateTimeSeriesPlot([][]float64{contango}, "Time", "Contango", "VIX Futures", "contango")

	// Backtest TQQQ Strategy
	tqqq_hist, err := client.GetHistoricalData("TQQQ")

	if err != nil {
		log.Println("Error Fetching Historical Data")
	}

	position := false
	var entry float64
	var exit float64

	var win_count int
	var loss_count int
	var trade_count int

	var avg_win float64
	var avg_loss float64
	var avg_trade float64

	var cumulative_pnl []float64
	cumulative_pnl = append(cumulative_pnl, 10000)

	for i := 1; i < len(vix_front_month_hist); i++ {

		// Enter Long
		if contango[i] > 0 && !position {

			entry = tqqq_hist.History.Day[i].Close
			position = true

		}

		// Exit Long
		if contango[i] < 0 && position {

			exit = tqqq_hist.History.Day[i].Close
			position = false

			trade := (exit - entry) * 100

			if trade > 0 {
				win_count++
				avg_win += trade
			}

			if trade < 0 {
				loss_count++
				avg_loss += trade
			}

			trade_count++
			avg_trade += trade

		}

		// Enter Short

		// Exit Short

		// Update PnL
		if position {

			diff := (tqqq_hist.History.Day[i].Close - tqqq_hist.History.Day[i-1].Close) * 100
			cumulative_pnl = append(cumulative_pnl, (cumulative_pnl[i-1] + diff))

		} else {

			cumulative_pnl = append(cumulative_pnl, cumulative_pnl[i-1])

		}

	}

	fmt.Println("Win Count: ", win_count)
	fmt.Println("Loss Count: ", loss_count)
	fmt.Println("Trade Count: ", trade_count)
	fmt.Println("")

	fmt.Println("Average Win: ", avg_win/float64(win_count))
	fmt.Println("Average Loss: ", avg_loss/float64(loss_count))
	fmt.Println("Average Trade: ", avg_trade/float64(trade_count))
	fmt.Println("")

	// Visualize Data
	p.CreateTimeSeriesPlot([][]float64{cumulative_pnl}, "Time", "Profit / Loss", "TQQQ Apeing", "tqqq_test")

}
