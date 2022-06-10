package Mongo

import (
	"encoding/csv"
	"log"
	"os"
)

func GetStockTickers(isMain bool) []string {

	var f *os.File
	var err error

	if isMain {
		f, err = os.Open("Mongo/StockTickers.csv")

	} else {
		f, err = os.Open("StockTickers.csv")

	}

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
