package BlackScholes

type BlackScholesParameters struct {
	Stock              float64
	Strike             float64
	Expiry             float64
	Implied_volatility float64
	Stochastic_drift   float64
	Dividends          float64
}
