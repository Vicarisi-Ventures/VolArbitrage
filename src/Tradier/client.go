package Tradier

import "net/http"

type TradierClient struct {
	Client      *http.Client
	AccessToken string
	AccountId   string
}

func NewTradierClient(accessToken string, accountId string) *TradierClient {

	return &TradierClient{
		Client:      &http.Client{},
		AccessToken: accessToken,
		AccountId:   accountId,
	}

}
