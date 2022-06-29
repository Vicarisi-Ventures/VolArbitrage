package VIX

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
	This Method Returns VIX Futures Data

	Please Note:

	1. This Data Should be Compared with a Risk Asset, and VIX

	2. 	This Object Contains Price Dynamics for the Front Month
*/
func FetchVixFuturesHistorical() PriceEOD {

	api_token := "61da06e72c1842.41500629"

	current_time := time.Now()
	current := current_time.Format("2006-01-02")
	past_time := current_time.AddDate(-1, 0, 0)
	past := past_time.Format("2006-01-02")

	resp, err := http.Get("https://eodhistoricaldata.com/api/eod/VIX.COMM?from=" + past + "&to=" + current + "&period=d&fmt=json&api_token=" + api_token)

	if err != nil {
		log.Println(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	var response PriceEOD
	json.NewDecoder(bytes.NewReader(responseData)).Decode(&response)

	return response

}
