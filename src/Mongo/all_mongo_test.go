package Mongo

import (
	"fmt"
	"testing"
)

func TestMongoDB(t *testing.T) {

	client := GetMongoConnection()

	// Each Collection ==> Stock Ticker
	ticker_symbols := GetStockTickers(false)

	// Variance Risk Premium
	var vrp60 []float64

	// Implied Distribution Skewness
	var skew []float64

	for i := 0; i < len(ticker_symbols); i++ {

		VAD := FetchVolatilityMongoDB(client, ticker_symbols[i])

		vrp60 = append(vrp60, VAD.VRP.VRP60)

		if VAD.CallIV.IV60 < 0.10 || VAD.PutIV.IV60 < 0.10 {
			skew = append(skew, 0)
		} else {
			skew = append(skew, (VAD.CallIV.IV60 / VAD.PutIV.IV60))
		}

	}

	/*
		Variance Risk Premium Screen
	*/

	var top_tickers []string
	index, top_five_vrp60 := findTopFive(vrp60)

	for i := 0; i < 5; i++ {
		top_tickers = append(top_tickers, ticker_symbols[index[i]])
	}

	fmt.Println("*** Top Five Highest VRP ***")
	for i := 0; i < 5; i++ {

		fmt.Println("Trade: ", i)
		fmt.Println("Risk Premium: ", top_five_vrp60[i])
		fmt.Println("Ticker Symbol: ", top_tickers[i])
		fmt.Println("")

	}

	var bottom_tickers []string
	index, bottom_five_vrp60 := findBottomFive(vrp60)

	for i := 0; i < 5; i++ {
		bottom_tickers = append(bottom_tickers, ticker_symbols[index[i]])
	}

	fmt.Println("*** Top Five Lowest VRP ***")
	for i := 0; i < 5; i++ {

		fmt.Println("Trade: ", i)
		fmt.Println("Risk Premium: ", bottom_five_vrp60[i])
		fmt.Println("Ticker Symbol: ", bottom_tickers[i])
		fmt.Println("")

	}

	/*
		Implied Distribution Skewness Screen
	*/

	var top_call_skew_tickers []string
	index, top_five_call_skew := findTopFive(skew)

	for i := 0; i < 5; i++ {
		top_call_skew_tickers = append(top_call_skew_tickers, ticker_symbols[index[i]])
	}

	fmt.Println("*** Top Five Highest Call Skew ***")
	for i := 0; i < 5; i++ {

		fmt.Println("Trade: ", i)
		fmt.Println("Call Skew: ", top_five_call_skew[i])
		fmt.Println("Ticker Symbol: ", top_call_skew_tickers[i])
		fmt.Println("")

	}

	var top_put_skew_tickers []string
	index, top_five_put_skew := findBottomFive(skew)

	for i := 0; i < 5; i++ {
		top_put_skew_tickers = append(top_put_skew_tickers, ticker_symbols[index[i]])
	}

	fmt.Println("*** Top Five Highest Put Skew ***")
	for i := 0; i < 5; i++ {

		fmt.Println("Trade: ", i)
		fmt.Println("Put Skew: ", top_five_put_skew[i])
		fmt.Println("Ticker Symbol: ", top_put_skew_tickers[i])
		fmt.Println("")

	}

}

func findTopFive(arr []float64) ([5]int, [5]float64) {

	var index [5]int
	var values [5]float64

	for i := 0; i < len(arr); i++ {

		for j := 0; j < 5; j++ {

			var isBreak bool

			// Begin Swap & Shift
			if arr[i] > values[j] {

				// Store Old Values
				old_index := index[j]
				old_value := values[j]

				// Update New Values
				index[j] = i
				values[j] = arr[i]

				// Shift Values
				for k := (j + 1); k < 4; k++ {

					index[k] = old_index
					old_index = index[k+1]

					values[k] = old_value
					old_value = values[k+1]

				}

				isBreak = true

			}

			if isBreak {
				break
			}

		}

	}

	return index, values

}

func findBottomFive(arr []float64) ([5]int, [5]float64) {

	var index [5]int
	var values [5]float64

	for i := 0; i < len(arr); i++ {

		for j := 0; j < 5; j++ {

			var isBreak bool

			// Begin Swap & Shift
			if arr[i] < values[j] {

				// Store Old Values
				old_index := index[j]
				old_value := values[j]

				// Update New Values
				index[j] = i
				values[j] = arr[i]

				// Shift Values
				for k := (j + 1); k < 4; k++ {

					index[k] = old_index
					old_index = index[k+1]

					values[k] = old_value
					old_value = values[k+1]

				}

				isBreak = true

			}

			if isBreak {
				break
			}

		}

	}

	return index, values

}
