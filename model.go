package main

//Issue - struct to map with mongodb documents
type Movie struct {
	ID       int    `bson:"_id"`
	Name     string `bson:"name"`
	Budget   string `bson:"budget"`
	Director string `bson:"director"`
}

type Producer struct {
	ID      int    `bson:"_id"`
	Name    string `bson:"name"`
	Address string `bson:"address"`
}
