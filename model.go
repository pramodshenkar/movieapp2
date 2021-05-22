package main

//Issue - struct to map with mongodb documents
type Movie struct {

	// ID          primitive.ObjectID `bson:"_id"`
	// CreatedAt time.Time `bson:"created_at"`
	// UpdatedAt time.Time `bson:"updated_at"`
	ID       int    `bson:"_id"`
	Name     string `bson:"name"`
	Budget   string `bson:"budget"`
	Director string `bson:"director"`
	// Producers string `bson:"producers"`
	// Actors    string `bson:"actors"`
}
