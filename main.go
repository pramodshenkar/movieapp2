package main

import (
	"context"
	"fmt"
	"time"

	connectionhelper "github.com/pramodshenkar/movieapp2/connectionHelper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi")

	r := gin.Default()
	r.POST("/issue", CreateIssueHandler)
	r.POST("/issues", CreateManyIssueHandler)
	r.GET("/issue", GetIssuesByCodeHandler)
	r.GET("/issues", GetAllIssuesHandler)
	r.PUT("/issues", MarkCompletedHandler)
	r.DELETE("/issue", DeleteOneHandler)
	r.DELETE("/issues", DeleteAllHandler)

	r.Run(":8000")

}

func CreateIssueHandler(c *gin.Context) {

	var issue = Issue{
		ID:          1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       "issue1",
		Code:        "<html>",
		Description: "Error",
		Completed:   false,
	}

	CreateIssue(issue)

}

func CreateManyIssueHandler(c *gin.Context) {

	var issues = []Issue{
		{
			ID:          2,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       "issue1",
			Code:        "<head>",
			Description: "Error",
			Completed:   false,
		},
		{
			ID:          3,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       "issue1",
			Code:        "<body>",
			Description: "Error",
			Completed:   false,
		},
		{
			ID:          4,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       "issue1",
			Code:        "<h1>",
			Description: "Error",
			Completed:   false,
		},
		{
			ID:          5,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       "issue1",
			Code:        "<h1>",
			Description: "Error",
			Completed:   false,
		}}

	CreateMany(issues)

}

func GetIssuesByCodeHandler(c *gin.Context) {
	code := "<httml>"
	Issue, _ := GetIssuesByCode(code)
	c.JSON(200, Issue)
}

func GetAllIssuesHandler(c *gin.Context) {
	Issues, _ := GetAllIssues()
	c.JSON(200, Issues)
}

func MarkCompletedHandler(c *gin.Context) {
	code := "<httml>"
	err := MarkCompleted(code)
	c.JSON(200, err)
}

func DeleteOneHandler(c *gin.Context) {
	code := "<httml>"
	err := DeleteOne(code)
	c.JSON(200, err)
}

func DeleteAllHandler(c *gin.Context) {
	err := DeleteAll()
	c.JSON(200, err)
}

//CreateIssue - Insert a new document in the collection.
func CreateIssue(task Issue) error {
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	//Perform InsertOne operation & validate against the error.
	_, err = collection.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

func CreateMany(list []Issue) error {
	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	//Perform InsertMany operation & validate against the error.
	_, err = collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

//GetIssuesByCode - Get All issues for collection
func GetIssuesByCode(code string) (Issue, error) {
	result := Issue{}
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "code", Value: code}}
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return result, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	//Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	//Return result without any error.
	return result, nil
}

//GetAllIssues - Get All issues for collection
func GetAllIssues() ([]Issue, error) {
	//Define filter query for fetching specific document from collection
	filter := bson.D{{}} //bson.D{{}} specifies 'all documents'
	issues := []Issue{}
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return issues, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return issues, findError
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := Issue{}
		err := cur.Decode(&t)
		if err != nil {
			return issues, err
		}
		issues = append(issues, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(issues) == 0 {
		return issues, mongo.ErrNoDocuments
	}
	return issues, nil
}

// MarkCompleted - MarkCompleted
func MarkCompleted(code string) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "code", Value: code}}

	//Define updater for to specifiy change to be updated.
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "completed", Value: true},
	}}}

	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)

	//Perform UpdateOne operation & validate against the error.
	_, err = collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

//DeleteOne - Get All issues for collection
func DeleteOne(code string) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "code", Value: code}}
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	//Perform DeleteOne operation & validate against the error.
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

//DeleteAll - Get All issues for collection
func DeleteAll() error {
	//Define filter query for fetching specific document from collection
	selector := bson.D{{}} // bson.D{{}} specifies 'all documents'
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	//Perform DeleteMany operation & validate against the error.
	_, err = collection.DeleteMany(context.TODO(), selector)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}
