package Volatility

type VolatilityMethodsObj struct {
	CloseToClose   []float64 `json:"CloseToClose"`
	Overnight      []float64 `json:"Overnight"`
	Intraday       []float64 `json:"Intraday"`
	RogersSatchell []float64 `json:"RogersSatchell"`
	YangZhang      []float64 `json:"YangZhang"`
}
