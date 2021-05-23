package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	connectionhelper "github.com/pramodshenkar/movieapp2/connectionHelper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	fmt.Println("Hi")

	r := gin.Default()
	r.POST("/movie", AddMovieHandler)
	r.POST("/movies", CreateManyMovieHandler)

	r.GET("/movie", GetMoviesByNameHandler)
	r.GET("/movies", GetAllMoviesHandler)

	// r.PUT("/movie/:id", UpdateRecordHandler)
	r.PUT("/movie", UpdateRecordHandler)
	r.DELETE("/movie", DeleteOneHandler)
	r.DELETE("/movies", DeleteAllHandler)

	// r.Run(":8000")
	r.Run()

}

func AddMovieHandler(c *gin.Context) {

	var movie Movie
	c.Bind(&movie)

	moviebson := bson.D{{Key: "_id", Value: movie.ID}, {Key: "name", Value: movie.Name}, {Key: "budget", Value: movie.Budget}, {Key: "director", Value: movie.Director}}
	err := AddRecord(c, moviebson, "movies")

	if err != nil {
		log.Println(err)
	}
}

func AddProducerHandler(c *gin.Context) {
	var producer Producer
	c.Bind(&producer)

	producerbson := bson.D{{Key: "_id", Value: producer.ID}, {Key: "name", Value: producer.Name}, {Key: "address", Value: producer.Address}}
	err := AddRecord(c, producerbson, "producer")

	if err != nil {
		log.Println(err)
	}
}

func CreateManyMovieHandler(c *gin.Context) {

	var movies = []Movie{
		{
			ID:       2,
			Name:     "PK",
			Budget:   "10C",
			Director: "Rajkumar Hirani",
		},
		{
			ID:       3,
			Name:     "Happy new year",
			Budget:   "10C",
			Director: "Rajkumar Hirani",
		},
		{
			ID:       4,
			Name:     "Bahubali",
			Budget:   "10C",
			Director: "Rajkumar Hirani",
		},
		{
			ID:       5,
			Name:     "Saho",
			Budget:   "10C",
			Director: "Rajkumar Hirani",
		}}

	CreateMany(movies)

}

func GetMoviesByNameHandler(c *gin.Context) {
	name, _ := http.Get("name")

	// Movie, _ := GetMoviesByName(name)
	c.JSON(200, name)
}

func GetAllMoviesHandler(c *gin.Context) {
	Movies, _ := GetAllRecords()
	c.JSON(200, Movies)
}

func UpdateRecordHandler(c *gin.Context) {
	var movie Movie
	c.Bind(&movie)

	// moviebson := bson.D{{Key: "_id", Value: movie.ID}, {Key: "name", Value: movie.Name}, {Key: "budget", Value: movie.Budget}, {Key: "director", Value: movie.Director}}
	// err := AddRecord(c, moviebson, "movies")

	// if err != nil {
	// 	log.Println(err)
	// }

	error := UpdateRecord(movie)
	c.JSON(200, error)

}

func DeleteOneHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		log.Println(err)
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, res)
}

func DeleteAllHandler(c *gin.Context) {
	movie := DeleteAll()
	c.JSON(200, movie)
}
