package db

import (
	"context"
	"time"

	"github.com/pborman/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client of the db
var Client *mongo.Client

// Open opens mongodb/establishes connection
func Open() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

// CreateToken creates an individual token (currency of the virtual-economy)
func CreateToken() error {
	identifier := uuid.New()
	collection := Client.Database("economy").Collection("tokens")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, bson.M{"tokenid": identifier})
	if err != nil {
		return err
	}
	return nil
}
