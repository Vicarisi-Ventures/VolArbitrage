package main

import (
	"encoding/csv"
	"log"
	"os"
)

func getStockTickers() []string {

	f, err := os.Open("StockTickers.csv")

	if err != nil {
		log.Fatal("Unable to read input file")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV")
	}

	var tickers []string
	for i := 0; i < len(records); i++ {

		tickers = append(tickers, records[i][0])

	}

	return tickers

}
