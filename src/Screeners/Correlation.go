package Screener

import (
	m "v2/src/Mongo"
	v "v2/src/Volatility"

	"github.com/montanaflynn/stats"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCorrelationMetrics(mongo *mongo.Client) {

	// Iterate Thru Stock Tickers
	ticker_symbols := m.GetStockTickers(true)

	for i := 0; i < (len(ticker_symbols) - 1); i++ {

		stockA := m.FetchHistoricalMongoDB(mongo, ticker_symbols[i])
		computeCorrelation(mongo, ticker_symbols, i, stockA)

	}

}

func computeCorrelation(mongo *mongo.Client, ticker_symbols []string, i int, stockA v.VolatilityMethodsParameters) {

	for j := 0; j < (len(ticker_symbols) - 1); j++ {

		if ticker_symbols[i] != ticker_symbols[j] {

			stockB := m.FetchHistoricalMongoDB(mongo, ticker_symbols[j])

			var open_corr float64
			var high_corr float64
			var low_corr float64
			var close_corr float64
			var err error

			if len(stockA.Close) == len(stockB.Close) {

				open_corr, err = stats.Correlation(stockA.Open, stockB.Open)

				if err != nil {
					open_corr = 0.0
				}

				high_corr, err = stats.Correlation(stockA.High, stockB.High)

				if err != nil {
					high_corr = 0.0
				}

				low_corr, err = stats.Correlation(stockA.Low, stockB.Low)

				if err != nil {
					low_corr = 0.0
				}

				close_corr, err = stats.Correlation(stockA.Close, stockB.Close)

				if err != nil {
					close_corr = 0.0
				}

			}

			m.AppendCorrelationMongo(mongo, ticker_symbols[i], ticker_symbols[j], open_corr, high_corr, low_corr, close_corr)

		} else {

			m.AppendCorrelationMongo(mongo, ticker_symbols[i], ticker_symbols[j], 0.0, 0.0, 0.0, 0.0)

		}

	}

}
