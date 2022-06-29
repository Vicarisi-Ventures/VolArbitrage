package Tradier

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (client *TradierClient) GetQuote(ticker string) (QuoteChain, error) {

	resp, err := client._get("/markets/quotes?symbols="+ticker+"&greeks=false", []byte(""))

	var response QuoteChain

	err = client.handleReponse(resp, &response)

	return response, err

}

func (client *TradierClient) GetHistoricalData(ticker string) (HistoricalData, error) {

	current_time := time.Now()
	current := current_time.Format("2006-01-02")

	past_time := current_time.AddDate(-1, 0, 0)
	past := past_time.Format("2006-01-02")

	resp, err := client._get("/markets/history?symbol="+ticker+"&interval=daily&start="+past+"&end="+current, []byte(""))

	var response HistoricalData

	err = client.handleReponse(resp, &response)

	return response, err

}

func (client *TradierClient) GetExpirations(symbol string, includeAllRoots string, strikes string) (OptionExpirations, error) {

	resp, err := client._get("/markets/options/expirations?symbol="+symbol+"&includeAllRoots="+includeAllRoots+"&strikes="+strikes, []byte(""))

	var response OptionExpirations

	err = client.handleReponse(resp, &response)

	return response, err

}

func (client *TradierClient) GetOptions(symbol string, expiration string, greeks string) (OptionsChain, error) {

	resp, err := client._get("/markets/options/chains?symbol="+symbol+"&expiration="+expiration+"&greeks="+greeks, []byte(""))

	var response OptionsChain

	err = client.handleReponse(resp, &response)

	return response, err

}

func (client *TradierClient) handleReponse(resp *http.Response, result interface{}) error {

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error processing response 1:", err)
		return err
	}

	err = json.Unmarshal(body, result)

	if err != nil {
		log.Println(string(body))
		log.Println("Error processing response 2:", err)
		return err
	}

	return nil

}
