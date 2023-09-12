package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found")
	// }

	// uri := os.Getenv("MONGODB_URI")
	// if uri == "" {
	// 	log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	// }
	uri := "http://localhost:27017"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	Collection := client.Database("SampleDB").Collection("MyData")
	
	//CREATE
	document := bson.M{
		"name":  "hogehoge",
		"age":   30,
		"email": "hogehoge@example.com",
	}
	res, err := Collection.InsertOne(context.Background(), document)
	if err != nil { 
		log.Fatal(err)
	}
	log.Println(res.InsertedID)

	//READ
	readFilter := bson.M{"name": "hogehoge"}
	cur, err := Collection.Find(context.Background(), readFilter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	
	for cur.Next(context.Background()) {
		var result bson.M
		if err := cur.Decode(&result); err != nil {
			log.Fatal(err)
		}
		jsonData, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", jsonData)
	}

	//UPDATE
	updateFilter := bson.M{"name": "hogehoge"}
	update := bson.M{"$set": bson.M{"age": 31}}

	_, err = Collection.UpdateOne(context.Background(), updateFilter, update)
	if err != nil {
		log.Fatal(err)
	}

	
	//DELETE
	deleteFilter := bson.M{"name": "hogehoge"}
	_, err = Collection.DeleteOne(context.Background(), deleteFilter)
	if err != nil {
		log.Fatal(err)
	}
}
