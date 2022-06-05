package Mongo

import (
	"context"
	"fmt"
	"testing"
)

func TestMongoDB(t *testing.T) {

	client := GetMongoConnection()

	fmt.Println("Connection: ", client.Connect(context.Background()))

}
