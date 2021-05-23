package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi")

	r := gin.Default()
	r.POST("/movie", AddMovieHandler)
	r.POST("/movies", AddManyMovieHandler)

	r.GET("/movie", GetMoviesByNameHandler)
	r.GET("/movies", GetAllMoviesHandler)

	// r.PUT("/movie/:id", UpdateMovieHandler)
	r.PUT("/movie", UpdateMovieHandler)
	r.DELETE("/movie", DeleteOneHandler)
	r.DELETE("/movies", DeleteAllHandler)

	// r.Run(":8000")
	r.Run()

}

//--------------- Add One Movie -------------

/*
http://localhost:8080/movie

{
			"ID":       1,
			"Name":     "Munnabhai MBBS",
			"Budget":   "10C",
			"Director": "Rajkumar Hirani"
}
*/
func AddMovieHandler(c *gin.Context) {

	var movie Movie
	c.Bind(&movie)
	res, err := AddMovie(movie)

	if err != nil {
		c.JSON(409, err)
	} else {
		c.JSON(200, res)
	}
}

//--------------- Add Multiple Movie -------------
/*
http://localhost:8080/movies

[
	{
		"ID": 2,
		"Name": "PK",
		"Budget": "10C",
		"Director": "Rajkumar Hirani",
	},
	{
		"ID": 3,
		"Name": "Happy new year",
		"Budget": "10C",
		"Director": "Rajkumar Hirani",
	},
	{
		"ID": 4,
		"Name": "Bahubali",
		"Budget": "10C",
		"Director": "Rajkumar Hirani",
	},
	{
		"ID": 5,
		"Name": "Saho",
		"Budget": "10C",
		"Director": "Rajkumar Hirani",
	}
]

*/
func AddManyMovieHandler(c *gin.Context) {

	var movies []Movie
	c.Bind(&movies)
	CreateMany(movies)
}

//--------------- Get Movie By Id -------------
/*
http://localhost:8080/movie?id=2
*/

func GetMoviesByNameHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	Movie, err := GetMoviesByName(id)

	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, Movie)
	}
}

//--------------- Get All Movies -------------
/*
http://localhost:8080/movies
*/

func GetAllMoviesHandler(c *gin.Context) {
	Movies, err := GetAllMovies()
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, Movies)
	}
}

//--------------- Get Movie By Id -------------
/*
http://localhost:8080/movie

{
			"ID":       1,
			"Name":     "Joker",
			"Budget":   "10C",
			"Director": "Ashutosh Gowarikar"
}
*/

func UpdateMovieHandler(c *gin.Context) {
	var movie Movie
	c.Bind(&movie)

	res, err := UpdateMovie(movie)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//-------------- Get Movie By Id -------------
/*
http://localhost:8080/movie?id=2
*/

func DeleteOneHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	res, err := DeleteOne(id)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//-------------- Get Movie By Id -------------
/*
http://localhost:8080/movies
*/
func DeleteAllHandler(c *gin.Context) {
	res, err := DeleteAll()
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}
