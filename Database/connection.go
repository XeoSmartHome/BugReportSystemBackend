package Database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var DB = ConnectToDatabase()

func ConnectToDatabase() *mongo.Database {
	connectionString := os.Getenv("MONGO_CONNECTION_STRING")
	client, _ := mongo.Connect(nil, options.Client().ApplyURI(connectionString))
	return client.Database("XeoSmartHome")
}
