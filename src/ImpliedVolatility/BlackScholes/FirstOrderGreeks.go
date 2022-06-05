package BlackScholes

import (
	"math"
	n "v2/src/Numerics"
)

/*
	Object Declarations
*/

type FirstOrderGreeks struct {
	CallDelta float64
	PutDelta  float64
	Vega      float64
	CallTheta float64
	PutTheta  float64
	CallRho   float64
	PutRho    float64
}

/*
	Individual Function Calls
*/

func GetCallDelta(class BlackScholesParameters) float64 {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return math.Exp(-(class.Dividends * class.Expiry)) * n.AdaptiveBoole(-5, d.D1)

}

func GetPutDelta(class BlackScholesParameters) float64 {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return -math.Exp(-(class.Dividends * class.Expiry)) * n.AdaptiveBoole(-5, -d.D1)

}

func GetVega(class BlackScholesParameters) float64 {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return math.Exp(-class.Dividends*class.Expiry) * class.Stock * n.GaussianDistribution(d.D1, 0, 1) * math.Sqrt(class.Expiry)

}

/*
	Object Function Calls
*/

func GetFirstOrderGreeks(class BlackScholesParameters) FirstOrderGreeks {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	var greeks FirstOrderGreeks
	greeks.CallDelta = math.Exp(-(class.Dividends * class.Expiry)) * n.AdaptiveBoole(-5, d.D1)
	greeks.PutDelta = -math.Exp(-(class.Dividends * class.Expiry)) * n.AdaptiveBoole(-5, -d.D1)
	greeks.Vega = math.Exp(-(class.Dividends * class.Expiry)) * class.Stock * n.GaussianDistribution(d.D1, 0, 1) * math.Sqrt(class.Expiry)
	greeks.CallTheta = -math.Exp(-(class.Dividends*class.Expiry)*(class.Stock*n.GaussianDistribution(d.D1, 0, 1)*class.Implied_volatility/2*math.Sqrt(class.Expiry))) + (class.Dividends * class.Stock * math.Exp(-(class.Dividends * class.Expiry)) * n.AdaptiveBoole(-5, d.D1))
	greeks.PutTheta = -math.Exp(-(class.Dividends*class.Expiry))*(class.Stock*n.GaussianDistribution(d.D1, 0, 1)*class.Implied_volatility/2*math.Sqrt(class.Expiry)) - (class.Dividends * class.Stock * math.Exp(-(class.Dividends * class.Expiry)) * n.AdaptiveBoole(-5, -d.D1))
	greeks.CallRho = class.Strike * class.Expiry * n.AdaptiveBoole(-5, d.D2)
	greeks.PutRho = -class.Strike * class.Expiry * n.AdaptiveBoole(-5, -d.D2)

	return greeks

}
