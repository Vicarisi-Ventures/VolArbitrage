package ImpliedVolatility

import (
	"math"
	b "v2/src/ImpliedVolatility/BlackScholes"
)

var max_iteration int = 10
var error_tol float64 = 0.001

func GetImpliedVolatility(class b.BlackScholesParameters, option_price float64, option_type string) float64 {

	iVol := GetHouseholderMethod(class, option_price, option_type)

	return iVol

}

func getManasterKoehler(class b.BlackScholesParameters) float64 {

	/*
	 * Formula Derived From: Manaster and Koehler (1982)
	 *
	 * The seed value will guarantee convergence if the implied volatility exists.
	 * Additionally, it will require fewer iterations to converge to the root.
	 */

	abs := math.Abs(math.Log(class.Stock/class.Strike)) + (class.Stochastic_drift * class.Expiry)

	return math.Sqrt(abs * (2.0 * class.Expiry))

}
