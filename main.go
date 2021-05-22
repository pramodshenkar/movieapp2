package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi")

	r := gin.Default()
	r.POST("/movie", CreateMovieHandler)
	r.POST("/movies", CreateManyMovieHandler)
	r.GET("/movie", GetMoviesByNameHandler)
	r.GET("/movies", GetAllMoviesHandler)
	r.PUT("/movies", MarkCompletedHandler)
	r.DELETE("/movie", DeleteOneHandler)
	r.DELETE("/movies", DeleteAllHandler)

	r.Run(":8000")

}

func CreateMovieHandler(c *gin.Context) {

	var movie = Movie{
		ID:       1,
		Name:     "3 idiots",
		Budget:   "10C",
		Director: "Rajkumar Hirani",
	}

	CreateMovie(movie)

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
	name := ""
	Movie, _ := GetMoviesByName(name)
	c.JSON(200, Movie)
}

func GetAllMoviesHandler(c *gin.Context) {
	Movies, _ := GetAllMovies()
	c.JSON(200, Movies)
}

func MarkCompletedHandler(c *gin.Context) {
	name := "<httml>"
	err := MarkCompleted(name)
	c.JSON(200, err)
}

func DeleteOneHandler(c *gin.Context) {
	name := "<httml>"
	err := DeleteOne(name)
	c.JSON(200, err)
}

func DeleteAllHandler(c *gin.Context) {
	err := DeleteAll()
	c.JSON(200, err)
}
