package ImpliedVolatility

import (
	"math"
	b "v2/src/ImpliedVolatility/BlackScholes"
)

var max_iteration int = 10
var error_tol float64 = 0.001

func GetImpliedVolatility(class b.BlackScholesParameters, option_price float64, option_type string) float64 {

	iVol := GetNewtonMethod(class, option_price, option_type)

	return iVol

}

/*
	Formula Derived From: Manaster and Koehler (1982)

	The seed value will guarantee convergence if the implied volatility exists.

	Additionally, it will require fewer iterations to converge to the root.

*/
func getManasterKoehler(class b.BlackScholesParameters) float64 {

	abs := math.Abs(math.Log(class.Stock/class.Strike)) + (class.Stochastic_drift * class.Expiry)

	return math.Sqrt(abs * (2.0 * class.Expiry))

}

/*
	Formula Derived From: An Explicit Implied Volatility Formula
*/
func getDanStefanica(class b.BlackScholesParameters, option_price float64) float64 {

	y := math.Log(class.Stock / class.Strike)
	alpha := getAlpha(class, option_price)
	R := getR(alpha, y)

	A := getA(y)
	B := getB(y, R)
	C := getC(y, R)
	beta := getBeta(A, B, C)
	gamma := getGamma(beta)

	var C_naught float64

	if y >= 0 {

		C_naught = class.Strike * math.Exp(-class.Stochastic_drift*class.Expiry) * (math.Exp(y)*A*math.Sqrt(2*y) - 0.5)

		if option_price <= C_naught {

			return (1.0 / math.Sqrt(class.Expiry)) * (math.Sqrt(gamma+y) - math.Sqrt(gamma-y))

		} else {

			return (1.0 / math.Sqrt(class.Expiry)) * (math.Sqrt(gamma+y) + math.Sqrt(gamma-y))

		}

	} else {

		C_naught = class.Strike * math.Exp(-class.Stochastic_drift*class.Expiry) * ((math.Exp(y) / 2.0) - A*(-math.Sqrt(-2*y)))

		if option_price <= C_naught {

			return (1.0 / math.Sqrt(class.Expiry)) * (-math.Sqrt(gamma+y) + math.Sqrt(gamma-y))

		} else {

			return (1.0 / math.Sqrt(class.Expiry)) * (math.Sqrt(gamma+y) + math.Sqrt(gamma-y))

		}

	}

}

func getAlpha(class b.BlackScholesParameters, option_price float64) float64 {

	return option_price / (class.Strike * math.Exp(-class.Stochastic_drift*class.Expiry))

}

func getR(alpha float64, y float64) float64 {

	return (2 * alpha) - math.Exp(y) + 1

}

func getBeta(A float64, B float64, C float64) float64 {

	return (2 * C) / (B + math.Sqrt((math.Pow(B, 2) + (4 * A * C))))

}

func getGamma(beta float64) float64 {

	return (-math.Pi / 2.0) * math.Log(beta)

}

func getA(y float64) float64 {

	exp1 := math.Exp(1 - ((2 / math.Pi) * y))
	exp2 := math.Exp(-(1 - (2 / math.Pi)) * y)

	return math.Pow(exp1-exp2, 2)

}

func getB(y float64, R float64) float64 {

	exp1 := math.Exp(2 / math.Pi * y)
	exp2 := math.Exp(-2 / math.Pi * y)

	exp3 := math.Exp(1 - ((2 / math.Pi) * y))
	exp4 := math.Exp(-(1 - (2 / math.Pi)) * y)

	return (4 * (exp1 + exp2)) - (2*math.Exp(-y))*(exp3+exp4)*(math.Exp(2*y)+1-math.Pow(R, 2))

}

func getC(y float64, R float64) float64 {

	exp1 := math.Exp(y)
	R_squared := math.Pow(R, 2)

	return math.Exp(-2*y) * (R_squared - math.Pow(exp1-1, 2)) * (math.Pow(exp1+1, 2) - R_squared)

}
