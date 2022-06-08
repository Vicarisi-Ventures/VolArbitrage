package Mongo

import (
	"fmt"
	"testing"
)

func TestMongoDB(t *testing.T) {

	client := GetMongoConnection()

	// Each Collection ==> Stock Ticker
	ticker_symbols := GetStockTickers(false)

	var vrp60 []float64

	for i := 0; i < len(ticker_symbols); i++ {

		VAD := FetchMongoDB(client, ticker_symbols[i])
		vrp60 = append(vrp60, VAD.VRP.VRP60)

	}

	var top_tickers []string
	index, top_ten_vrp60 := findTopTen(vrp60)

	for i := 0; i < 10; i++ {
		top_tickers = append(top_tickers, ticker_symbols[index[i]])
	}

	fmt.Println("*** Top Ten Highest VRP ***")
	for i := 0; i < 10; i++ {

		fmt.Println("Trade: ", i)
		fmt.Println("Risk Premium: ", top_ten_vrp60[i])
		fmt.Println("Ticker Symbol: ", top_tickers[i])
		fmt.Println("")

	}

	var bottom_tickers []string
	index, bottom_ten_vrp60 := findBottomTen(vrp60)

	for i := 0; i < 10; i++ {
		bottom_tickers = append(bottom_tickers, ticker_symbols[index[i]])
	}

	fmt.Println("*** Bottom Ten Lowest VRP ***")
	for i := 0; i < 10; i++ {

		fmt.Println("Trade: ", i)
		fmt.Println("Risk Premium: ", bottom_ten_vrp60[i])
		fmt.Println("Ticker Symbol: ", bottom_tickers[i])
		fmt.Println("")

	}

}

func findTopTen(arr []float64) ([10]int, [10]float64) {

	var index [10]int
	var values [10]float64

	for i := 0; i < len(arr); i++ {

		for j := 0; j < 10; j++ {

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
				for k := (j + 1); k < 9; k++ {

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

func findBottomTen(arr []float64) ([10]int, [10]float64) {

	var index [10]int
	var values [10]float64

	for i := 0; i < len(arr); i++ {

		for j := 0; j < 10; j++ {

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
				for k := (j + 1); k < 9; k++ {

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
