package Screener

import (
	m "v2/src/Mongo"

	"github.com/montanaflynn/stats"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCorrelationMetrics(mongo *mongo.Client) {

	// Iterate Thru Stock Tickers
	ticker_symbols := m.GetStockTickers(true)

	for i := 0; i < len(ticker_symbols); i++ {

		stockA := m.FetchHistoricalMongoDB(mongo, ticker_symbols[i])

		for j := 0; j < len(ticker_symbols); j++ {

			if ticker_symbols[i] != ticker_symbols[j] {

				stockB := m.FetchHistoricalMongoDB(mongo, ticker_symbols[j])

				open_corr, _ := stats.Correlation(stockA.Open, stockB.Open)
				high_corr, _ := stats.Correlation(stockA.High, stockB.High)
				low_corr, _ := stats.Correlation(stockA.Low, stockB.Low)
				close_corr, _ := stats.Correlation(stockA.Close, stockB.Close)

				m.AppendCorrelationMongo(mongo, ticker_symbols[i], ticker_symbols[j], open_corr, high_corr, low_corr, close_corr)

			}

		}

	}

}
