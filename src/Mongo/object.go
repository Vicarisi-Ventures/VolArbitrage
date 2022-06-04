package Mongo

type VolArbitrageData struct {
	HV struct {
		HV30  float64
		HV60  float64
		HV90  float64
		HV120 float64
	}

	IV struct {
		IV30  float64
		IV60  float64
		IV90  float64
		IV120 float64
	}

	VRP struct {
		VRP30  float64
		VRP60  float64
		VRP90  float64
		VRP120 float64
	}

	CallIV struct {
		IV30  float64
		IV60  float64
		IV90  float64
		IV120 float64
	}

	PutIV struct {
		IV30  float64
		IV60  float64
		IV90  float64
		IV120 float64
	}
}
