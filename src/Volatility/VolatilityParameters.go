package Volatility

type VolatilityMethodsParameters struct {
	OHLC struct {
		Open  []float64
		High  []float64
		Low   []float64
		Close []float64
	}

	RollingPeriod int
}
