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

// collection
var collection *mongo.Collection

// Open opens mongodb/establishes connection
func Open() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	collection = Client.Database("economy").Collection("tokens")
	defer func() {
		if err = Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

// CreateToken creates an individual token (currency of the virtual-economy)
func CreateToken() error {
	identifier := uuid.New()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, bson.M{"tokenid": identifier})
	if err != nil {
		return err
	}
	return nil
}

// AllTokens Returns all the tokens in the document
func AllTokens() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return err
		}
	}
	if err := cur.Err(); err != nil {
		return err
	}
}
