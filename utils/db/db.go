package db

import (
	"context"
	"time"

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
func CreateToken() {
	
}
