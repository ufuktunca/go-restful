package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	Name      string   `bson:"name"`
	Year      string   `bson:"year"`
	Directors []string `bson:"directors"`
	Writers   []string `bson:"writers"`
	BoxOffice `bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `bson:"budget"`
	Gross  uint64 `bson:"gross"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB successfully")
	collection := client.Database("appDB").Collection("movies")

	darkNight := Movie{
		Name:      "The Dark Night",
		Year:      "2008",
		Directors: []string{"Cristopher Nolan"},
		Writers:   []string{"Jonathan Nolan", "Cristopher Nolan"},
		BoxOffice: BoxOffice{
			Budget: 5454545,
			Gross:  44454545,
		},
	}

	_, err = collection.InsertOne(context.TODO(), darkNight)

	if err != nil {
		log.Fatal(err)
	}

	queryResult := &Movie{}
	filter := bson.M{"boxOffice.budget": bson.M{"$gt": 5200}}
	result := collection.FindOne(context.TODO(), filter)
	err = result.Decode(queryResult)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie:", queryResult)

	err = client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}
