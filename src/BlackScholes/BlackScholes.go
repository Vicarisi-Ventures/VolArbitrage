package BlackScholes

import (
	"log"
	"math"
	"time"

	n "v2/src/Numerics"
)

// COMPUTE BLACK SCHOLES RANDOM WALK PARAMETERS

func getRandomWalk(stock float64, strike float64, sigma float64, tau float64) RandomWalk {

	var walk RandomWalk
	walk.D1 = (math.Log(stock/strike) + ((math.Pow(sigma, 2) / 2) * tau)) / (sigma * math.Sqrt(tau))
	walk.D2 = (math.Log(stock/strike) + ((-math.Pow(sigma, 2) / 2) * tau)) / (sigma * math.Sqrt(tau))

	return walk

}

func GetCall(class BlackScholesParameters) float64 {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return (math.Exp((-class.Dividends * class.Expiry)) * class.Stock * n.AdaptiveBoole(-5, d.D1)) - (class.Strike * n.AdaptiveBoole(-5, d.D2))

}

func GetPut(class BlackScholesParameters) float64 {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return (class.Strike * n.AdaptiveBoole(-5, -d.D2)) - (math.Exp((-class.Dividends * class.Expiry)) * class.Stock * n.AdaptiveBoole(-5, -d.D1))

}

func GetVega(class BlackScholesParameters) float64 {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return math.Exp(-class.Dividends*class.Expiry) * class.Stock * n.GaussianDistribution(d.D1, 0, 1) * math.Sqrt(class.Expiry)

}

// COMPUTE BLACK SCHOLES OPTION PRICE

func GetBlackScholes(class BlackScholesParameters) BlackScholes {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	var optionPrice BlackScholes
	optionPrice.CallPrice = (math.Exp((-class.Dividends * class.Expiry)) * class.Stock * n.AdaptiveBoole(-5, d.D1)) - (class.Strike * n.AdaptiveBoole(-5, d.D2))
	optionPrice.PutPrice = (class.Strike * n.AdaptiveBoole(-5, -d.D2)) - (math.Exp((-class.Dividends * class.Expiry)) * class.Stock * n.AdaptiveBoole(-5, -d.D1))

	return optionPrice

}

func GetCallDelta(class BlackScholesParameters) float64 {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return math.Exp(-(class.Dividends * class.Expiry)) * n.AdaptiveBoole(-5, d.D1)

}

func GetPutDelta(class BlackScholesParameters) float64 {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return -math.Exp(-(class.Dividends * class.Expiry)) * n.AdaptiveBoole(-5, -d.D1)

}

// COMPUTE BLACK SCHOLES FIRST ORDER GREEKS

func GetFirstOrderGreeks(class BlackScholesParameters) FirstOrderGreeks {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

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

// INDIVIDUAL FUNCTIONS TO SPEED UP SECOND ORDER EXPOSURES

func GetGamma(class BlackScholesParameters) float64 {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return (math.Exp(-class.Dividends*class.Expiry) * n.GaussianDistribution(d.D1, 0, 1)) / (class.Stock * class.Implied_volatility * math.Sqrt(class.Expiry))

}

func GetVanna(class BlackScholesParameters) float64 {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return -math.Exp(-(class.Dividends * class.Expiry)) * n.GaussianDistribution(d.D1, 0, 1) * d.D2 / class.Implied_volatility

}

func GetCallCharm(class BlackScholesParameters) float64 {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return class.Dividends*math.Exp(-(class.Dividends*class.Expiry))*n.AdaptiveBoole(-5, d.D1) - math.Exp(-(class.Dividends*class.Expiry))*n.GaussianDistribution(d.D1, 0, 1)*(2*-class.Dividends*class.Expiry-d.D2*class.Implied_volatility*math.Sqrt(class.Expiry)/2*class.Expiry*class.Implied_volatility*math.Sqrt(class.Expiry))

}

func GetPutCharm(class BlackScholesParameters) float64 {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return -class.Dividends*math.Exp(-(class.Dividends*class.Expiry))*n.AdaptiveBoole(-5, -d.D1) - math.Exp(-(class.Dividends*class.Expiry))*n.GaussianDistribution(d.D1, 0, 1)*(2*-class.Dividends*class.Expiry-d.D2*class.Implied_volatility*math.Sqrt(class.Expiry)/2*class.Expiry*class.Implied_volatility*math.Sqrt(class.Expiry))

}

func GetVomma(class BlackScholesParameters) float64 {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return class.Stock * n.GaussianDistribution(d.D1, 0, 1) * math.Sqrt(class.Expiry) * d.D1 * d.D2 / class.Implied_volatility

}

// COMPUTE BLACK SCHOLES SECOND ORDER GREEKS

func GetSecondOrderGreeks(class BlackScholesParameters) SecondOrderGreeks {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	var greeks SecondOrderGreeks
	greeks.Gamma = (math.Exp(-class.Dividends*class.Expiry) * n.GaussianDistribution(d.D1, 0, 1)) / (class.Stock * class.Implied_volatility * math.Sqrt(class.Expiry))
	greeks.Vanna = -math.Exp(-(class.Dividends * class.Expiry)) * n.GaussianDistribution(d.D1, 0, 1) * d.D2 / class.Implied_volatility
	greeks.CallCharm = class.Dividends*math.Exp(-(class.Dividends*class.Expiry))*n.AdaptiveBoole(-5, d.D1) - math.Exp(-(class.Dividends*class.Expiry))*n.GaussianDistribution(d.D1, 0, 1)*(2*-class.Dividends*class.Expiry-d.D2*class.Implied_volatility*math.Sqrt(class.Expiry)/2*class.Expiry*class.Implied_volatility*math.Sqrt(class.Expiry))
	greeks.PutCharm = -class.Dividends*math.Exp(-(class.Dividends*class.Expiry))*n.AdaptiveBoole(-5, -d.D1) - math.Exp(-(class.Dividends*class.Expiry))*n.GaussianDistribution(d.D1, 0, 1)*(2*-class.Dividends*class.Expiry-d.D2*class.Implied_volatility*math.Sqrt(class.Expiry)/2*class.Expiry*class.Implied_volatility*math.Sqrt(class.Expiry))
	greeks.Vomma = class.Stock * math.Exp(-(class.Dividends * class.Expiry)) * n.GaussianDistribution(d.D1, 0, 1) * math.Sqrt(class.Expiry) * d.D1 * d.D2 / class.Implied_volatility
	greeks.Veta = -class.Stock * math.Exp(-(class.Dividends * class.Expiry)) * n.GaussianDistribution(d.D1, 0, 1) * math.Sqrt(class.Expiry) * (class.Dividends + (-class.Dividends * d.D1 / class.Implied_volatility * math.Sqrt(class.Expiry)) - ((1 + d.D1*d.D2) / 2 * class.Expiry))

	return greeks

}

func GetUltima(class BlackScholesParameters) float64 {

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	return (((class.Stock * n.GaussianDistribution(d.D1, 0, 1)) * math.Sqrt(class.Expiry)) / math.Pow(class.Implied_volatility, 2)) * (d.D1*d.D2*(1-d.D1*d.D2) + math.Pow(d.D1, 2) + math.Pow(d.D2, 2))

}

func GetThirdOrderGreeks(class BlackScholesParameters) ThirdOrderGreeks {

	var greeks ThirdOrderGreeks

	d := getRandomWalk(class.Stock, class.Strike, class.Implied_volatility, class.Expiry)

	greeks.Speed = (n.GaussianDistribution(d.D1, 0, 1) / math.Pow(class.Stock, 2) * class.Implied_volatility * math.Sqrt(class.Expiry)) * ((d.D1 / class.Implied_volatility * math.Sqrt(class.Expiry)) + 1)
	greeks.Zomma = n.GaussianDistribution(d.D1, 0, 1) * (d.D1*d.D2 - 1) / (class.Stock * math.Pow(class.Implied_volatility, 2) * math.Sqrt(class.Expiry))
	greeks.Color = ((n.GaussianDistribution(d.D1, 0, 1)) / (2 * class.Stock * class.Expiry * class.Implied_volatility * math.Sqrt(class.Expiry))) * (1 + (-d.D2 * class.Implied_volatility * math.Sqrt(class.Expiry) * d.D1 / (class.Implied_volatility * math.Sqrt(class.Expiry))))
	greeks.Ultima = (((class.Stock * n.GaussianDistribution(d.D1, 0, 1)) * math.Sqrt(class.Expiry)) / math.Pow(class.Implied_volatility, 2)) * (d.D1*d.D2*(1-d.D1*d.D2) + math.Pow(d.D1, 2) + math.Pow(d.D2, 2))

	return greeks

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
