package Numerics

import "math"

func GaussianDistribution(x float64, mu float64, sigma float64) float64 {

	return (1 / (sigma * math.Sqrt(2*math.Pi))) * math.Exp((-0.5 * math.Pow((x-mu)/sigma, 2)))

}

func AdaptiveBoole(lower float64, upper float64) float64 {

	/*
		Fast Numerical Integration Routine for Smooth Functions
	*/

	err_tol := 0.01
	m := (lower + upper) / 2.0

	lm := (m + lower) / 2.0
	um := (m + upper) / 2.0
	h := upper - lower

	arr := []float64{GaussianDistribution(lower, 0, 1), GaussianDistribution(lm, 0, 1), GaussianDistribution(m, 0, 1), GaussianDistribution(um, 0, 1), GaussianDistribution(upper, 0, 1)}
	boole := (h / 90.0) * (7*arr[0] + 32*arr[1] + 12*arr[2] + 32*arr[3] + 7*arr[4])

	m1 := (lower + lm) / 2.0
	m2 := (lm + m) / 2.0
	m3 := (m + um) / 2.0
	m4 := (um + upper) / 2.0

	marr := []float64{GaussianDistribution(m1, 0, 1), GaussianDistribution(m2, 0, 1), GaussianDistribution(m3, 0, 1), GaussianDistribution(m4, 0, 1)}
	area := (h / 180.0) * (7*arr[0] + 32*marr[0] + 12*arr[1] + 32*marr[1] + 14*arr[2] + 32*marr[2] + 12*arr[3] + 32*marr[3] + 7*arr[4])

	err := math.Abs(boole - area)

	if err > err_tol {

		lower_area := AdaptiveBoole(lower, m)
		upper_area := AdaptiveBoole(m, upper)
		area = lower_area + upper_area

	}

	return area

}

func AdaptiveSimpson(lower float64, upper float64) float64 {

	/*
		Fast Numerical Integration Routine for Smooth Functions
	*/

	err_tol := 0.0001
	h := upper - lower
	m := (lower + upper) / 2.0
	lm := (lower + m) / 2.0
	um := (m + upper) / 2.0
	arr := []float64{GaussianDistribution(lower, 0, 1), GaussianDistribution(lm, 0, 1), GaussianDistribution(m, 0, 1), GaussianDistribution(um, 0, 1), GaussianDistribution(upper, 0, 1)}
	simpson := (h / 6) * (arr[0] + (4 * arr[2]) + arr[4])
	area := (h / 12) * (arr[0] + 4*arr[1] + 2*arr[2] + 4*arr[3] + arr[4])

	diff := math.Abs(simpson - area)
	if diff > err_tol {

		lower := AdaptiveSimpson(lower, m)
		upper := AdaptiveSimpson(m, upper)
		area = lower + upper

	}

	return area

}

func AdaptiveTrapezoid(lower float64, upper float64) float64 {

	/*
		Fast Numerical Integration Routine for Rigid Functions
	*/

	err_tol := 0.0001
	h := upper - lower
	m := (lower + upper) / 2.0
	arr := []float64{GaussianDistribution(lower, 0, 1), GaussianDistribution(m, 0, 1), GaussianDistribution(upper, 0, 1)}
	trapezoid := (h / 2) * (arr[0] + arr[1])
	area := (h / 4) * (arr[0] + (2 * arr[1]) + arr[2])

	diff := math.Abs(trapezoid - area)
	if diff > err_tol {

		lower := AdaptiveTrapezoid(lower, m)
		upper := AdaptiveTrapezoid(m, upper)
		area = lower + upper

	}

	return area

}
