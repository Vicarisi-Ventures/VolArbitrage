package Mongo

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoConnection() *mongo.Client {

	password := "APuXI7kPYRKNhaYA"

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://vicarisiventures:" + password + "@optioncluster.ing0x.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Println("Error Initializing Mongo Client")
	}

	return client

}

func AppendMongo(client *mongo.Client, class VolArbitrageData, coll_name string) {

	collection := client.Database("VolArbitrage").Collection(coll_name)

	result, err := collection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println(result)

	currentTime := time.Now()

	// Insert newest document
	insert, err := collection.InsertOne(
		context.Background(),
		bson.D{

			{Key: "date", Value: currentTime.Format("01-02-2006")},

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

func FetchMongoDB(client *mongo.Client, coll_name string) VolArbitrageData {

	var VAD VolArbitrageData

	collection := client.Database("VolArbitrage").Collection(coll_name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var iterations []bson.M
	err = cursor.All(ctx, &iterations)

	if err != nil {
		log.Fatal(err)
	}

	for _, itr := range iterations {

		// Variance Risk Premium
		VAD.VRP.VRP30 = itr["riskPremia"].(primitive.M)["vrp30"].(float64)
		VAD.VRP.VRP60 = itr["riskPremia"].(primitive.M)["vrp60"].(float64)
		VAD.VRP.VRP90 = itr["riskPremia"].(primitive.M)["vrp90"].(float64)
		VAD.VRP.VRP120 = itr["riskPremia"].(primitive.M)["vrp120"].(float64)

	}

	return VAD

}

func GetStockTickers(isMain bool) []string {

	var f *os.File
	var err error

	if isMain {
		f, err = os.Open("Mongo/StockTickers.csv")

	} else {
		f, err = os.Open("StockTickers.csv")

	}

	if err != nil {
		log.Fatal("Unable to read input file")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV")
	}

	var tickers []string
	for i := 0; i < len(records); i++ {

		tickers = append(tickers, records[i][0])

	}

	return tickers

}
