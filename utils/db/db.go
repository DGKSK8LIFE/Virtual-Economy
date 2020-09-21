package db

import (
	"context"
	"log"
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
		log.Fatal(err)
	}
	collection = Client.Database("economy").Collection("tokens")
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
func AllTokens() (bson.M, error) {
	var result bson.M
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

// TokenCount returns the number of tokens in circulation
func TokenCount() (bson.M, error) {
	var result bson.M
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return result, nil

}
