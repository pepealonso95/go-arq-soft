package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-rabbit/config"
)

var client *mongo.Client = nil

// GetConnection returns a mongo connection
func GetConnection() (*mongo.Client, error) {
	if client != nil {
		return client, nil
	}
	connectionString := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		config.Get.MongoDB.User,
		config.Get.MongoDB.Password,
		config.Get.MongoDB.Address,
		config.Get.MongoDB.Name,
	)
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return client, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	return client, err
}
