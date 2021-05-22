package main

import (
	"context"

	connectionhelper "github.com/pramodshenkar/movieapp2/connectionHelper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateIssue(task Issue) error {
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

func CreateMany(list []Issue) error {
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

func GetIssuesByCode(code string) (Issue, error) {
	result := Issue{}
	filter := bson.D{primitive.E{Key: "code", Value: code}}
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

func GetAllIssues() ([]Issue, error) {
	filter := bson.D{{}}
	issues := []Issue{}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return issues, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return issues, findError
	}
	for cur.Next(context.TODO()) {
		t := Issue{}
		err := cur.Decode(&t)
		if err != nil {
			return issues, err
		}
		issues = append(issues, t)
	}
	cur.Close(context.TODO())
	if len(issues) == 0 {
		return issues, mongo.ErrNoDocuments
	}
	return issues, nil
}

func MarkCompleted(code string) error {
	filter := bson.D{primitive.E{Key: "code", Value: code}}

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

func DeleteOne(code string) error {
	filter := bson.D{primitive.E{Key: "code", Value: code}}
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
