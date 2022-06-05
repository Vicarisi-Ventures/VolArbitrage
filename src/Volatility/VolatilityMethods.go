package Volatility

import (
	"math"
)

func GetVolatilityMethods(class VolatilityMethodsParameters) VolatilityMethodsObj {

	var v VolatilityMethodsObj

	v.CloseToClose = GetCloseToClose(class)
	v.Overnight = getOvernightVol(class)
	v.Intraday = getIntraDayVol(class)
	v.RogersSatchell = getRogersSatchell(class)
	v.YangZhang = getYangZhang(class)

	// fmt.Println(len(v.CloseToClose))
	// fmt.Println(len(v.Overnight))
	// fmt.Println(len(v.Intraday))
	// fmt.Println(len(v.RogersSatchell))
	// fmt.Println(len(v.YangZhang))

	return v

}

// COMPUTE STANDARD VOLATILITY

func GetCloseToClose(class VolatilityMethodsParameters) []float64 {

	var sd []float64
	arr := getLogReturn(class.OHLC.Close)

	for i := class.RollingPeriod; i < len(arr); i++ {

		var mu float64

		for j := i; j > (i - class.RollingPeriod); j-- {
			mu += math.Pow(arr[j], 2)
		}

		mu = mu / float64(class.RollingPeriod)
		var summation float64

		for j := i; j > (i - class.RollingPeriod); j-- {
			summation += math.Pow(arr[j]-mu, 2)
		}

		variance := summation / float64(class.RollingPeriod)
		sigma := math.Sqrt(variance)
		sd = append(sd, (sigma * math.Sqrt(252)))

	}

	return sd

}

// COMPUTE OVERNIGHT VOLATILITY

func getOvernightVol(class VolatilityMethodsParameters) []float64 {

	var sd []float64
	var arr []float64

	for i := 0; i < len(class.OHLC.Open); i++ {

		if i > 0 {
			arr = append(arr, (class.OHLC.Open[i] / class.OHLC.Close[i-1]))
		}

	}

	log_returns := getLogReturn(arr)

	for i := class.RollingPeriod; i < len(log_returns); i++ {

		var mu float64

		for j := i; j > (i - class.RollingPeriod); j-- {
			mu += math.Pow(arr[j], 2)
		}

		mu = mu / float64(class.RollingPeriod)
		var summation float64

		for j := i; j > (i - class.RollingPeriod); j-- {
			summation += math.Pow(arr[j]-mu, 2)
		}

		variance := summation / float64(class.RollingPeriod)
		sigma := math.Sqrt(variance)
		sd = append(sd, (sigma * math.Sqrt(252)))

	}

	return sd

}

// COMPUTE INTRADAY VOLATILITY

func getIntraDayVol(class VolatilityMethodsParameters) []float64 {

	var sd []float64
	var arr []float64

	for i := 0; i < len(class.OHLC.Open); i++ {

		if i > 0 {
			arr = append(arr, (class.OHLC.Close[i] / class.OHLC.Open[i]))
		}

	}

	log_returns := getLogReturn(arr)

	for i := class.RollingPeriod; i < len(log_returns); i++ {

		var mu float64

		for j := i; j > (i - class.RollingPeriod); j-- {
			mu += math.Pow(arr[j], 2)
		}

		mu = mu / float64(class.RollingPeriod)
		var summation float64

		for j := i; j > (i - class.RollingPeriod); j-- {
			summation += math.Pow(arr[j]-mu, 2)
		}

		variance := summation / float64(class.RollingPeriod)
		sigma := math.Sqrt(variance)
		sd = append(sd, (sigma * math.Sqrt(252)))

	}

	return sd

}

// COMPUTE ROGERS SATCHELL VOLATILITY

func getRogersSatchell(class VolatilityMethodsParameters) []float64 {

	var sd []float64
	var hc []float64
	var ho []float64
	var lc []float64
	var lo []float64

	for i := 0; i < len(class.OHLC.Close); i++ {

		hc = append(hc, class.OHLC.High[i]/class.OHLC.Close[i])
		ho = append(ho, class.OHLC.High[i]/class.OHLC.Open[i])
		lc = append(lc, class.OHLC.Low[i]/class.OHLC.Close[i])
		lo = append(lo, class.OHLC.Low[i]/class.OHLC.Open[i])

	}

	var arr_hc []float64
	var arr_ho []float64
	var arr_lc []float64
	var arr_lo []float64

	for j := 0; j < len(hc); j++ {

		arr_hc = append(arr_hc, math.Log(hc[j]))
		arr_ho = append(arr_ho, math.Log(ho[j]))
		arr_lc = append(arr_lc, math.Log(lc[j]))
		arr_lo = append(arr_lo, math.Log(lo[j]))

	}

	for i := class.RollingPeriod; i < len(arr_hc); i++ {

		var summation float64

		for j := i; j > (i - class.RollingPeriod); j-- {
			summation += (arr_hc[j] * arr_ho[j]) + (arr_lc[j] * arr_lo[j])
		}

		variance := summation / float64(class.RollingPeriod)
		sigma := math.Sqrt(variance)
		sd = append(sd, (sigma * math.Sqrt(252)))

	}

	return sd
}

// COMPUTE YANG ZHANG VOLATILITY

func getYangZhang(class VolatilityMethodsParameters) []float64 {

	var sd []float64
	alpha := 1.34

	k := (alpha - 1) / (alpha + (float64((class.RollingPeriod + 1) / (class.RollingPeriod - 1))))

	overnight_vol := getOvernightVol(class)
	open_to_close_vol := getIntraDayVol(class)
	rogers_satchell_vol := getRogersSatchell(class)

	for i := 0; i < len(overnight_vol); i++ {

		volatility := math.Sqrt(math.Pow(overnight_vol[i], 2) + (k * math.Pow(open_to_close_vol[i], 2)) + ((1 - k) * math.Pow(rogers_satchell_vol[i], 2)))
		sd = append(sd, volatility)

	}

	return sd

}

func getLogReturn(arr []float64) []float64 {

	var prc []float64

	for i := 0; i < len(arr)-1; i++ {
		prc = append(prc, (math.Log(arr[i+1] / arr[i])))
	}

	return prc

}
