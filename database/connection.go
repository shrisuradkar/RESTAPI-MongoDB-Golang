package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}
	mongoDB := os.Getenv("MONGODB_URI")
	if mongoDB == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDB))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MonogoDB")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading env file")
	}
	dbname := os.Getenv("DBNAME")
	if dbname == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}
	var collection *mongo.Collection = client.Database(dbname).Collection(collectionName)
	return collection
}
