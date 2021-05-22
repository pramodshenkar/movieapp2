package main

import (
	"context"

	connectionhelper "github.com/pramodshenkar/movieapp2/connectionHelper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMovie(task Movie) error {
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	_, err = collection.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}
	return nil
}

func CreateMany(list []Movie) error {
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	_, err = collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return err
	}
	return nil
}

func GetMoviesByName(name string) (Movie, error) {
	result := Movie{}
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetAllMovies() ([]Movie, error) {
	filter := bson.D{{}}
	movies := []Movie{}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return movies, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return movies, findError
	}
	for cur.Next(context.TODO()) {
		t := Movie{}
		err := cur.Decode(&t)
		if err != nil {
			return movies, err
		}
		movies = append(movies, t)
	}
	cur.Close(context.TODO())
	if len(movies) == 0 {
		return movies, mongo.ErrNoDocuments
	}
	return movies, nil
}

func MarkCompleted(name string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "completed", Value: true},
	}}}

	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)

	_, err = collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOne(name string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAll() error {
	selector := bson.D{{}} // bson.D{{}} specifies 'all documents'
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	_, err = collection.DeleteMany(context.TODO(), selector)
	if err != nil {
		return err
	}
	return nil
}
