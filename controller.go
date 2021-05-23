package main

import (
	"context"

	"github.com/gin-gonic/gin"
	connectionhelper "github.com/pramodshenkar/movieapp2/connectionHelper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// func CreateMovie(task bson.D, c *gin.Context) error {
// 	client, err := connectionhelper.GetMongoClient()
// 	if err != nil {
// 		return err
// 	}
// 	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
// 	_, err = collection.InsertOne(context.TODO(), task)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func AddRecord(c *gin.Context, record bson.D, collectionName string) error {
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(connectionhelper.DB).Collection(collectionName)
	_, err = collection.InsertOne(context.TODO(), record)
	if err != nil {
		c.JSON(409, err)
		return err
	} else {
		c.JSON(200, "Record Added")
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

func GetAllRecords() ([]Movie, error) {
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

func UpdateRecord(movie Movie) error {
	filter := bson.D{primitive.E{Key: "_id", Value: movie.ID}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{{Key: "name", Value: movie.Name}, {Key: "budget", Value: movie.Budget}, {Key: "director", Value: movie.Director}}}}

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

// func DeleteOne(id string) error {
// 	filter := bson.D{primitive.E{Key: "_id", Value: id}}
// 	client, err := connectionhelper.GetMongoClient()
// 	if err != nil {
// 		// return err
// 		log.Println(err)
// 	}
// 	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
// 	result, err := collection.DeleteOne(context.TODO(), filter)

// 	// result := Movie{}
// 	// err = collection.FindOne(context.TODO(), filter).Decode(&result)

// 	if err != nil {
// 		// return err
// 		log.Println(err)

// 	}
// 	return nil
// }

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
