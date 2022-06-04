package Tradier

type HistoricalData struct {
	History struct {
		Day []struct {
			Date   string  `json:"date"`
			Open   float64 `json:"open"`
			High   float64 `json:"high"`
			Low    float64 `json:"low"`
			Close  float64 `json:"close"`
			Volume int     `json:"volume"`
		} `json:"day"`
	} `json:"history"`
}

type OptionExpirations struct {
	Expirations struct {
		Dates []string `json:"date"`
	} `json:"expirations"`
}

type OptionsChain struct {
	Option struct {
		Options []OptionInformation `json:"option"`
	} `json:"options"`
}

type OptionInformation struct {
	Symbol           string      `json:"symbol"`
	Description      string      `json:"description"`
	Exch             string      `json:"exch"`
	Type             string      `json:"type"`
	Last             interface{} `json:"last"`
	Change           interface{} `json:"change"`
	Volume           int         `json:"volume"`
	Open             interface{} `json:"open"`
	High             interface{} `json:"high"`
	Low              interface{} `json:"low"`
	Close            interface{} `json:"close"`
	Bid              float64     `json:"bid"`
	Ask              float64     `json:"ask"`
	Underlying       string      `json:"underlying"`
	Strike           float64     `json:"strike"`
	ChangePercentage interface{} `json:"change_percentage"`
	AverageVolume    int         `json:"average_volume"`
	LastVolume       int         `json:"last_volume"`
	TradeDate        int         `json:"trade_date"`
	Prevclose        interface{} `json:"prevclose"`
	Week52High       float64     `json:"week_52_high"`
	Week52Low        float64     `json:"week_52_low"`
	Bidsize          int         `json:"bidsize"`
	Bidexch          string      `json:"bidexch"`
	BidDate          int64       `json:"bid_date"`
	Asksize          int         `json:"asksize"`
	Askexch          string      `json:"askexch"`
	AskDate          int64       `json:"ask_date"`
	OpenInterest     int         `json:"open_interest"`
	ContractSize     int         `json:"contract_size"`
	ExpirationDate   string      `json:"expiration_date"`
	ExpirationType   string      `json:"expiration_type"`
	OptionType       string      `json:"option_type"`
	RootSymbol       string      `json:"root_symbol"`
	Greeks           struct {
		Delta     float64 `json:"delta"`
		Gamma     float64 `json:"gamma"`
		Theta     float64 `json:"theta"`
		Vega      float64 `json:"vega"`
		Rho       float64 `json:"rho"`
		Phi       float64 `json:"phi"`
		BidIv     float64 `json:"bid_iv"`
		MidIv     float64 `json:"mid_iv"`
		AskIv     float64 `json:"ask_iv"`
		SmvVol    float64 `json:"smv_vol"`
		UpdatedAt string  `json:"updated_at"`
	} `json:"option"`
}
