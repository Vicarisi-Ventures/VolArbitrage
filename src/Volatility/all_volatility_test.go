package Volatility

import (
	"fmt"
	"log"
	"testing"
	tradier "v2/src/Tradier"
)

func TestVolatilityMethods(t *testing.T) {

	var VMP VolatilityMethodsParameters

	api_key := ""
	account_id := ""

	// Initialize Tradier Client
	client := tradier.NewTradierClient(api_key, account_id)

	historical, err := client.GetHistoricalData("LCID")

	if err != nil {
		log.Println("Error Fetching Historical Data")
	}

	length := len(historical.History.Day)

	for j := 0; j < length; j++ {

		VMP.Open = append(VMP.Open, historical.History.Day[j].Open)
		VMP.High = append(VMP.High, historical.History.Day[j].High)
		VMP.Low = append(VMP.Low, historical.History.Day[j].Low)
		VMP.Close = append(VMP.Close, historical.History.Day[j].Close)

	}

	VMP.RollingPeriod = 30
	hv30 := GetVolatilityMethods(VMP).YangZhang[(length - VMP.RollingPeriod - 3)]

	VMP.RollingPeriod = 60
	hv60 := GetVolatilityMethods(VMP).YangZhang[(length - VMP.RollingPeriod - 3)]

	VMP.RollingPeriod = 90
	hv90 := GetVolatilityMethods(VMP).YangZhang[(length - VMP.RollingPeriod - 3)]

	VMP.RollingPeriod = 120
	hv120 := GetVolatilityMethods(VMP).YangZhang[(length - VMP.RollingPeriod - 3)]

	fmt.Println("Historical Volatility Metrics")
	fmt.Println("HV30: ", hv30)
	fmt.Println("HV60: ", hv60)
	fmt.Println("HV90: ", hv90)
	fmt.Println("HV120: ", hv120)
	fmt.Println("")

}
