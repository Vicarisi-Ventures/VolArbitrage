package Mongo

import (
	"context"
	"fmt"
	"log"
	v "v2/src/Volatility"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AppendVolatilityMongo(client *mongo.Client, class VolArbitrageData, coll_name string) {

	collection := client.Database("VolArbitrage").Collection(coll_name)

	result, err := collection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println(result)

	insert, err := collection.InsertOne(
		context.Background(),
		bson.D{

			{Key: "historical", Value: bson.D{
				{Key: "hv30", Value: class.HV.HV30},
				{Key: "hv60", Value: class.HV.HV60},
				{Key: "hv90", Value: class.HV.HV90},
				{Key: "hv120", Value: class.HV.HV120},
			}},

			{Key: "implied", Value: bson.D{
				{Key: "iv30", Value: class.IV.IV30},
				{Key: "iv60", Value: class.IV.IV60},
				{Key: "iv90", Value: class.IV.IV90},
				{Key: "iv120", Value: class.IV.IV120}},
			},

			{Key: "riskPremia", Value: bson.D{
				{Key: "vrp30", Value: class.VRP.VRP30},
				{Key: "vrp60", Value: class.VRP.VRP60},
				{Key: "vrp90", Value: class.VRP.VRP90},
				{Key: "vrp120", Value: class.VRP.VRP120}},
			},

			{Key: "callSkew", Value: bson.D{
				{Key: "cs30", Value: class.CallIV.IV30},
				{Key: "cs60", Value: class.CallIV.IV60},
				{Key: "cs90", Value: class.CallIV.IV90},
				{Key: "cs120", Value: class.CallIV.IV120},
				{Key: "callStrikes", Value: class.CallIV.TailStrikes}},
			},

			{Key: "putSkew", Value: bson.D{
				{Key: "ps30", Value: class.PutIV.IV30},
				{Key: "ps60", Value: class.PutIV.IV60},
				{Key: "ps90", Value: class.PutIV.IV90},
				{Key: "ps120", Value: class.PutIV.IV120},
				{Key: "putStrikes", Value: class.PutIV.TailStrikes}},
			},

			{Key: "expectedMove", Value: bson.D{
				{Key: "em30", Value: class.EM.EM30},
				{Key: "em60", Value: class.EM.EM60},
				{Key: "em90", Value: class.EM.EM90},
				{Key: "em120", Value: class.EM.EM120}},
			},
		})

	if err != nil {
		log.Println("Err with database")
	}

	fmt.Println(insert)

}

func AppendHistoricalMongo(client *mongo.Client, class v.VolatilityMethodsParameters, coll_name string) {

	collection := client.Database("Historical").Collection(coll_name)

	result, err := collection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println(result)

	insert, err := collection.InsertOne(
		context.Background(),
		bson.D{
			{Key: "open", Value: class.Open},
			{Key: "high", Value: class.High},
			{Key: "low", Value: class.Low},
			{Key: "close", Value: class.Close},
		})

	if err != nil {
		log.Println("Err with database")
	}

	fmt.Println(insert)

}

func AppendCorrelationMongo(client *mongo.Client, coll_name string, pair_name string, open_corr float64, high_corr float64, low_corr float64, close_corr float64) {

	collection := client.Database("Correlation").Collection(coll_name)

	pair := coll_name + "-" + pair_name
	result, err := collection.DeleteOne(context.Background(), bson.M{"title": pair})

	if err != nil {
		log.Println(err)
	}

	fmt.Println(result)

	insert, err := collection.InsertOne(
		context.Background(),
		bson.D{

			{Key: pair, Value: bson.D{
				{Key: "open", Value: open_corr},
				{Key: "high", Value: high_corr},
				{Key: "low", Value: low_corr},
				{Key: "close", Value: close_corr},
			}},
		})

	if err != nil {
		log.Println("Err with database")
	}

	fmt.Println(insert)

}
