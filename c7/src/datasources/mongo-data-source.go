package datasources

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDataSource(connectionURI string) (*mongo.Client, error) {
	// uri := os.Getenv("MONGODB_URI")
	// if uri == "" {
	// 	log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	// }
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionURI))
	if err != nil {
		return nil, err
	}
	return client, nil
}
