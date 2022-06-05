package BlackScholes

import (
	"math"
	n "v2/src/Numerics"
)

/*
	Object Declarations
*/

type ThirdOrderGreeks struct {
	Speed  float64
	Zomma  float64
	Color  float64
	Ultima float64
}

/*
	Individual Function Calls
*/

func GetUltima(class BlackScholesParameters) float64 {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return (((class.Stock * n.GaussianDistribution(d.D1, 0, 1)) * math.Sqrt(class.Expiry)) / math.Pow(class.Implied_volatility, 2)) * (d.D1*d.D2*(1-d.D1*d.D2) + math.Pow(d.D1, 2) + math.Pow(d.D2, 2))

}

/*
	Object Function Calls
*/

func GetThirdOrderGreeks(class BlackScholesParameters) ThirdOrderGreeks {

	var greeks ThirdOrderGreeks

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	greeks.Speed = (n.GaussianDistribution(d.D1, 0, 1) / math.Pow(class.Stock, 2) * class.Implied_volatility * math.Sqrt(class.Expiry)) * ((d.D1 / class.Implied_volatility * math.Sqrt(class.Expiry)) + 1)
	greeks.Zomma = n.GaussianDistribution(d.D1, 0, 1) * (d.D1*d.D2 - 1) / (class.Stock * math.Pow(class.Implied_volatility, 2) * math.Sqrt(class.Expiry))
	greeks.Color = ((n.GaussianDistribution(d.D1, 0, 1)) / (2 * class.Stock * class.Expiry * class.Implied_volatility * math.Sqrt(class.Expiry))) * (1 + (-d.D2 * class.Implied_volatility * math.Sqrt(class.Expiry) * d.D1 / (class.Implied_volatility * math.Sqrt(class.Expiry))))
	greeks.Ultima = (((class.Stock * n.GaussianDistribution(d.D1, 0, 1)) * math.Sqrt(class.Expiry)) / math.Pow(class.Implied_volatility, 2)) * (d.D1*d.D2*(1-d.D1*d.D2) + math.Pow(d.D1, 2) + math.Pow(d.D2, 2))

	return greeks

}
