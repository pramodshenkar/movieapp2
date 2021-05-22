package main

import (
	"fmt"
	"time"

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
