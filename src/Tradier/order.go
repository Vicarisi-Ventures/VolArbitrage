package Tradier

type OrderTicket struct {
	Class         string
	Symbol        string
	Type          string
	Duration      string
	Price         string
	OptionSymbols []string
	Side          []string
	Quantity      []string
}

func (client *TradierClient) PlaceMultiLegOptionOrder(account_id string, class OrderTicket) (Orders, error) {

	resp, err := client._post("/accounts/"+account_id+"/orders", []byte(""))

	var response Orders

	err = client.handleReponse(resp, &response)

	return response, err

}
