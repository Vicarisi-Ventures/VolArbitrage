package ImpliedVolatility

import (
	"math"
	g "v2/src/ImpliedVolatility/BlackScholes"
)

func GetNewtonMethod(class g.BlackScholesParameters, option_price float64, option_type string) float64 {

	class.Implied_volatility = getManasterKoehler(class)

	var vega float64

	count := 0

	if option_type == "call" || option_type == "Call" {

		guess_price := g.GetCall(class)

		for math.Abs(guess_price-option_price) > error_tol && count < max_iteration {

			vega = g.GetVega(class)
			guess_price = g.GetCall(class)
			class.Implied_volatility = class.Implied_volatility - ((guess_price - option_price) / vega)

			if math.IsNaN(class.Implied_volatility) || class.Implied_volatility < 0.0 || class.Implied_volatility > 5.0 {
				return 0
			}

			count++

		}

	} else {

		guess_price := g.GetPut(class)

		for math.Abs(guess_price-option_price) > error_tol && count < max_iteration {

			vega = g.GetVega(class)
			guess_price = g.GetPut(class)
			class.Implied_volatility = class.Implied_volatility - ((guess_price - option_price) / vega)

			if math.IsNaN(class.Implied_volatility) || class.Implied_volatility < 0.0 || class.Implied_volatility > 5.0 {
				return 0
			}

			count++

		}

	}

	return class.Implied_volatility

}

func GetHalleyMethod(class g.BlackScholesParameters, option_price float64, option_type string) float64 {

	class.Implied_volatility = getManasterKoehler(class)

	var vega float64
	var vomma float64

	count := 0

	if option_type == "call" || option_type == "Call" {

		guess_price := g.GetCall(class)

		for math.Abs(guess_price-option_price) > error_tol && count < max_iteration {

			vega = g.GetVega(class)
			vomma = g.GetVomma(class)
			guess_price = g.GetCall(class)
			class.Implied_volatility = class.Implied_volatility - (2*(guess_price-option_price)*(vega))/(2*math.Pow(vega, 2)-(guess_price-option_price)*(vomma))

			if math.IsNaN(class.Implied_volatility) || class.Implied_volatility < 0.0 || class.Implied_volatility > 5.0 {
				return 0
			}

			count++

		}

	} else {

		guess_price := g.GetPut(class)

		for math.Abs(guess_price-option_price) > error_tol && count < max_iteration {

			vega = g.GetVega(class)
			vomma = g.GetVomma(class)
			guess_price = g.GetPut(class)
			class.Implied_volatility = class.Implied_volatility - (2*(guess_price-option_price)*(vega))/(2*math.Pow(vega, 2)-(guess_price-option_price)*(vomma))

			if math.IsNaN(class.Implied_volatility) || class.Implied_volatility < 0.0 || class.Implied_volatility > 5.0 {
				return 0
			}

			count++

		}

	}

	return class.Implied_volatility

}

func GetHouseholderMethod(class g.BlackScholesParameters, option_price float64, option_type string) float64 {

	class.Implied_volatility = getManasterKoehler(class)

	var vega float64
	var vomma float64
	var ultima float64

	var h float64
	var top_part float64
	var bottom_part float64

	count := 0

	if option_type == "call" || option_type == "Call" {

		guess_price := g.GetCall(class)

		for math.Abs(guess_price-option_price) > error_tol && count < max_iteration {

			vega = g.GetVega(class)
			vomma = g.GetVomma(class)
			ultima = g.GetUltima(class)
			guess_price = g.GetCall(class)
			h = -(guess_price - option_price) / vega
			top_part = 1 + (0.5*(vomma/vega))*class.Implied_volatility*h
			bottom_part = 1 + (vomma/vega)*class.Implied_volatility*h + (1/6.0*(ultima/vega))*class.Implied_volatility*math.Pow(h, 2)
			class.Implied_volatility = class.Implied_volatility + h*(top_part/bottom_part)

			if math.IsNaN(class.Implied_volatility) || class.Implied_volatility < 0.0 || class.Implied_volatility > 5.0 {
				return 0
			}

			count++

		}

	} else {

		guess_price := g.GetPut(class)

		for math.Abs(guess_price-option_price) > error_tol && count < max_iteration {

			vega = g.GetVega(class)
			vomma = g.GetVomma(class)
			ultima = g.GetUltima(class)
			guess_price = g.GetPut(class)
			h = -(guess_price - option_price) / vega
			top_part = 1 + (0.5*(vomma/vega))*class.Implied_volatility*h
			bottom_part = 1 + (vomma/vega)*class.Implied_volatility*h + (1/6.0*(ultima/vega))*class.Implied_volatility*math.Pow(h, 2)
			class.Implied_volatility = class.Implied_volatility + h*(top_part/bottom_part)

			if math.IsNaN(class.Implied_volatility) || class.Implied_volatility < 0.0 || class.Implied_volatility > 5.0 {
				return 0
			}

			count++

		}

	}

	return class.Implied_volatility

}
