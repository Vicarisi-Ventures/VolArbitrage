package Mongo

import (
	"context"
	"log"
	"time"

	v "v2/src/Volatility"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchVolatilityMongoDB(client *mongo.Client, coll_name string) VolArbitrageData {

	var VAD VolArbitrageData

	collection := client.Database("VolArbitrage").Collection(coll_name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Println(err)
	}

	var iterations []bson.M
	err = cursor.All(ctx, &iterations)

	if err != nil {
		log.Println(err)
	}

	if coll_name != "FB" {

		for _, itr := range iterations {

			// Realized Volatility
			for i := 0; i < len(itr["historical"].(primitive.M)["hv30"].(primitive.A)); i++ {
				VAD.HV.HV30 = append(VAD.HV.HV30, itr["historical"].(primitive.M)["hv30"].(primitive.A)[i].(float64))
			}

			for i := 0; i < len(itr["historical"].(primitive.M)["hv60"].(primitive.A)); i++ {
				VAD.HV.HV60 = append(VAD.HV.HV60, itr["historical"].(primitive.M)["hv60"].(primitive.A)[i].(float64))
			}

			for i := 0; i < len(itr["historical"].(primitive.M)["hv90"].(primitive.A)); i++ {
				VAD.HV.HV90 = append(VAD.HV.HV90, itr["historical"].(primitive.M)["hv90"].(primitive.A)[i].(float64))
			}

			for i := 0; i < len(itr["historical"].(primitive.M)["hv120"].(primitive.A)); i++ {
				VAD.HV.HV120 = append(VAD.HV.HV120, itr["historical"].(primitive.M)["hv120"].(primitive.A)[i].(float64))
			}

			// Variance Risk Premium
			VAD.VRP.VRP30 = itr["riskPremia"].(primitive.M)["vrp30"].(float64)
			VAD.VRP.VRP60 = itr["riskPremia"].(primitive.M)["vrp60"].(float64)
			VAD.VRP.VRP90 = itr["riskPremia"].(primitive.M)["vrp90"].(float64)
			VAD.VRP.VRP120 = itr["riskPremia"].(primitive.M)["vrp120"].(float64)

			// Call Skew
			VAD.CallIV.IV30 = itr["callSkew"].(primitive.M)["cs30"].(float64)
			VAD.CallIV.IV60 = itr["callSkew"].(primitive.M)["cs60"].(float64)
			VAD.CallIV.IV90 = itr["callSkew"].(primitive.M)["cs90"].(float64)
			VAD.CallIV.IV120 = itr["callSkew"].(primitive.M)["cs120"].(float64)

			// Put Skew
			VAD.PutIV.IV30 = itr["putSkew"].(primitive.M)["ps30"].(float64)
			VAD.PutIV.IV60 = itr["putSkew"].(primitive.M)["ps60"].(float64)
			VAD.PutIV.IV90 = itr["putSkew"].(primitive.M)["ps90"].(float64)
			VAD.PutIV.IV120 = itr["putSkew"].(primitive.M)["ps120"].(float64)

		}

	}

	return VAD

}

func FetchHistoricalMongoDB(client *mongo.Client, coll_name string) v.VolatilityMethodsParameters {

	var VMP v.VolatilityMethodsParameters

	collection := client.Database("Historical").Collection(coll_name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Println(err)
	}

	var iterations []bson.M
	err = cursor.All(ctx, &iterations)

	if err != nil {
		log.Println(err)
	}

	for _, itr := range iterations {

		for i := 0; i < len(itr["open"].(primitive.A)); i++ {

			VMP.Open = append(VMP.Open, itr["open"].(primitive.A)[i].(float64))
			VMP.High = append(VMP.Open, itr["high"].(primitive.A)[i].(float64))
			VMP.Low = append(VMP.Open, itr["low"].(primitive.A)[i].(float64))
			VMP.Close = append(VMP.Open, itr["close"].(primitive.A)[i].(float64))

		}

	}

	return VMP

}

func FetchCorrelationMongoDB(client *mongo.Client, coll_name string, pair_name string, index int) [4]float64 {

	var arr [4]float64

	collection := client.Database("Correlation").Collection(coll_name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Println(err)
	}

	var iterations []bson.M
	err = cursor.All(ctx, &iterations)

	if err != nil {
		log.Println(err)
	}

	itr := iterations[index]

	// fmt.Println("Correlation: ", itr[pair_name].(primitive.M)["open"].(float64))
	// fmt.Println("Correlation: ", itr[pair_name].(primitive.M)["high"].(float64))
	// fmt.Println("Correlation: ", itr[pair_name].(primitive.M)["low"].(float64))
	// fmt.Println("Correlation: ", itr[pair_name].(primitive.M)["close"].(float64))

	arr[0] = itr[pair_name].(primitive.M)["open"].(float64)
	arr[1] = itr[pair_name].(primitive.M)["high"].(float64)
	arr[2] = itr[pair_name].(primitive.M)["low"].(float64)
	arr[3] = itr[pair_name].(primitive.M)["close"].(float64)

	return arr

}
