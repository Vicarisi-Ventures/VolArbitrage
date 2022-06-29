package BlackScholes

import (
	"math"
	n "v2/src/Numerics"
)

/*
	Object Declarations
*/

type SecondOrderGreeks struct {
	Gamma     float64
	Vanna     float64
	CallCharm float64
	PutCharm  float64
	Vomma     float64
	Veta      float64
}

/*
	Individual Function Calls
*/

func GetGamma(class BlackScholesParameters) float64 {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return (math.Exp(-class.Dividends*class.Expiry) * n.GaussianDistribution(d.D1, 0, 1)) / (class.Stock * class.Implied_volatility * math.Sqrt(class.Expiry))

}

func GetVanna(class BlackScholesParameters) float64 {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return -math.Exp(-(class.Dividends * class.Expiry)) * n.GaussianDistribution(d.D1, 0, 1) * d.D2 / class.Implied_volatility

}

func GetCallCharm(class BlackScholesParameters) float64 {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return class.Dividends*math.Exp(-(class.Dividends*class.Expiry))*n.AdaptiveBoole(-5, d.D1) - math.Exp(-(class.Dividends*class.Expiry))*n.GaussianDistribution(d.D1, 0, 1)*(2*-class.Dividends*class.Expiry-d.D2*class.Implied_volatility*math.Sqrt(class.Expiry)/2*class.Expiry*class.Implied_volatility*math.Sqrt(class.Expiry))

}

func GetPutCharm(class BlackScholesParameters) float64 {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return -class.Dividends*math.Exp(-(class.Dividends*class.Expiry))*n.AdaptiveBoole(-5, -d.D1) - math.Exp(-(class.Dividends*class.Expiry))*n.GaussianDistribution(d.D1, 0, 1)*(2*-class.Dividends*class.Expiry-d.D2*class.Implied_volatility*math.Sqrt(class.Expiry)/2*class.Expiry*class.Implied_volatility*math.Sqrt(class.Expiry))

}

func GetVomma(class BlackScholesParameters) float64 {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return class.Stock * n.GaussianDistribution(d.D1, 0, 1) * math.Sqrt(class.Expiry) * d.D1 * d.D2 / class.Implied_volatility

}

/*
	Object Function Calls
*/

func GetSecondOrderGreeks(class BlackScholesParameters) SecondOrderGreeks {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	var greeks SecondOrderGreeks
	greeks.Gamma = (math.Exp(-class.Dividends*class.Expiry) * n.GaussianDistribution(d.D1, 0, 1)) / (class.Stock * class.Implied_volatility * math.Sqrt(class.Expiry))
	greeks.Vanna = -math.Exp(-(class.Dividends * class.Expiry)) * n.GaussianDistribution(d.D1, 0, 1) * d.D2 / class.Implied_volatility
	greeks.CallCharm = class.Dividends*math.Exp(-(class.Dividends*class.Expiry))*n.AdaptiveBoole(-5, d.D1) - math.Exp(-(class.Dividends*class.Expiry))*n.GaussianDistribution(d.D1, 0, 1)*(2*-class.Dividends*class.Expiry-d.D2*class.Implied_volatility*math.Sqrt(class.Expiry)/2*class.Expiry*class.Implied_volatility*math.Sqrt(class.Expiry))
	greeks.PutCharm = -class.Dividends*math.Exp(-(class.Dividends*class.Expiry))*n.AdaptiveBoole(-5, -d.D1) - math.Exp(-(class.Dividends*class.Expiry))*n.GaussianDistribution(d.D1, 0, 1)*(2*-class.Dividends*class.Expiry-d.D2*class.Implied_volatility*math.Sqrt(class.Expiry)/2*class.Expiry*class.Implied_volatility*math.Sqrt(class.Expiry))
	greeks.Vomma = class.Stock * math.Exp(-(class.Dividends * class.Expiry)) * n.GaussianDistribution(d.D1, 0, 1) * math.Sqrt(class.Expiry) * (d.D1 * d.D2 / class.Implied_volatility)
	greeks.Veta = -class.Stock * math.Exp(-(class.Dividends * class.Expiry)) * n.GaussianDistribution(d.D1, 0, 1) * math.Sqrt(class.Expiry) * (class.Dividends + (-class.Dividends * d.D1 / class.Implied_volatility * math.Sqrt(class.Expiry)) - ((1 + d.D1*d.D2) / 2 * class.Expiry))

	return greeks

}
