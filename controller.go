package main

import (
	"context"

	connectionhelper "github.com/pramodshenkar/movieapp2/connectionHelper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddMovie(movie Movie) (*mongo.InsertOneResult, error) {
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection("movies")
	res, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		return nil, err
	}
	return res, nil
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

func GetMoviesByName(id int) (Movie, error) {
	result := Movie{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
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

func UpdateMovie(movie Movie) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: movie.ID}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{{Key: "name", Value: movie.Name}, {Key: "budget", Value: movie.Budget}, {Key: "director", Value: movie.Director}}}}

	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)

	res, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteOne(id int) (*mongo.DeleteResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteAll() (*mongo.DeleteResult, error) {
	selector := bson.D{{}} // bson.D{{}} specifies 'all documents'
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	res, err := collection.DeleteMany(context.TODO(), selector)
	if err != nil {
		return nil, err
	}
	return res, nil
}
