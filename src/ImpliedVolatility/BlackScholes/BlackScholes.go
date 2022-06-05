package BlackScholes

import (
	"log"
	"math"
	"time"

	n "v2/src/Numerics"
)

/*
	Object Declarations
*/

type RandomWalk struct {
	D1 float64
	D2 float64
}

type BlackScholes struct {
	CallPrice float64
	PutPrice  float64
}

/*
	Individual Function Calls
*/

func GetRandomWalk(stock float64, strike float64, sigma float64, tau float64) RandomWalk {

	var walk RandomWalk

	walk.D1 = (math.Log(stock/strike) + ((math.Pow(sigma, 2) / 2) * tau)) / (sigma * math.Sqrt(tau))
	walk.D2 = (math.Log(stock/strike) + ((-math.Pow(sigma, 2) / 2) * tau)) / (sigma * math.Sqrt(tau))

	return walk

}

func GetCall(class BlackScholesParameters) float64 {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return (math.Exp((-class.Dividends * class.Expiry)) * class.Stock * n.AdaptiveBoole(-5, d.D1)) - (class.Strike * n.AdaptiveBoole(-5, d.D2))

}

func GetPut(class BlackScholesParameters) float64 {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return (class.Strike * n.AdaptiveBoole(-5, -d.D2)) - (math.Exp((-class.Dividends * class.Expiry)) * class.Stock * n.AdaptiveBoole(-5, -d.D1))

}

/*
	Object Function Calls
*/

func GetBlackScholes(class BlackScholesParameters) BlackScholes {

	d := GetRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	var optionPrice BlackScholes
	optionPrice.CallPrice = (math.Exp((-class.Dividends * class.Expiry)) * class.Stock * n.AdaptiveBoole(-5, d.D1)) - (class.Strike * n.AdaptiveBoole(-5, d.D2))
	optionPrice.PutPrice = (class.Strike * n.AdaptiveBoole(-5, -d.D2)) - (math.Exp((-class.Dividends * class.Expiry)) * class.Stock * n.AdaptiveBoole(-5, -d.D1))

	return optionPrice

}

func GetDaysBetween(start time.Time, end string) float64 {

	parsedEnd, err := time.Parse("2006-01-02", end[:10])

	if err != nil {
		log.Println(err)
	}

	// contracts expire at 4:00 pm EST...
	// convert UTC to EST
	parsedEnd = parsedEnd.Add(time.Hour * 21)

	daysBetween := parsedEnd.Sub(start).Hours() / 24.0

	return daysBetween

}
