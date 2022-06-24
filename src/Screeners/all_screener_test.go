package Screener

import (
	"fmt"
	"log"
	"testing"
	m "v2/src/Mongo"

	"github.com/montanaflynn/stats"
)

func TestCorrelation(t *testing.T) {

	mongo := m.GetMongoConnection()

	stockA := m.FetchHistoricalMongoDB(mongo, "ABNB")
	stockB := m.FetchHistoricalMongoDB(mongo, "FB")

	var corr float64
	var err error

	if len(stockA.Close) == len(stockB.Close) {

		corr, err = stats.Correlation(stockA.Open, stockB.Open)

	}

	if err != nil {
		log.Println(err)
	}

	fmt.Println(corr)

}
