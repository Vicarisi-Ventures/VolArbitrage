package BlackScholes

type RandomWalk struct {
	D1 float64
	D2 float64
}

type BlackScholes struct {
	CallPrice float64
	PutPrice  float64
}

type FirstOrderGreeks struct {
	CallDelta float64
	PutDelta  float64
	Vega      float64
	CallTheta float64
	PutTheta  float64
	CallRho   float64
	PutRho    float64
}

type SecondOrderGreeks struct {
	Gamma     float64
	Vanna     float64
	CallCharm float64
	PutCharm  float64
	Vomma     float64
	Veta      float64
}

type ThirdOrderGreeks struct {
	Speed  float64
	Zomma  float64
	Color  float64
	Ultima float64
}
