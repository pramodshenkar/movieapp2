package main

import (
	"time"
)

//Issue - struct to map with mongodb documents
type Issue struct {
	// ID          primitive.ObjectID `bson:"_id"`
	ID          int       `bson:"_id"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
	Title       string    `bson:"title"`
	Code        string    `bson:"code"`
	Description string    `bson:"description"`
	Completed   bool      `bson:"completed"`
}
