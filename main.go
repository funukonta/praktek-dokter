package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}

	username := os.Getenv("USERMONGO")
	pass := os.Getenv("PASS")
	cluster := os.Getenv("CLUSTER")

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", url.QueryEscape(username), url.QueryEscape(pass), cluster)
	fmt.Println(uri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("klinik-dokter")
	coll := db.Collection("dokter")
	// find code goes here
	filter := bson.D{{Key: `nama`, Value: `evan`}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	// iterate code goes here
	for cursor.Next(context.TODO()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			panic(err)
		}
		fmt.Println(result)
	}
	if err := cursor.Err(); err != nil {
		panic(err)
	}

}
