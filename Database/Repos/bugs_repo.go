package Repos

import (
	"BugReportSystemBackend/Database"
	"BugReportSystemBackend/Database/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type bugsRepoType struct {
	collection *mongo.Collection
}

var BugsRepo = bugsRepoType{
	collection: Database.DB.Collection("Bugs"),
}

func (repo *bugsRepoType) GetAll() ([]Models.Bug, error) {
	cursor, err := repo.collection.Find(nil, bson.M{})
	if err != nil {
		return nil, err
	}
	var bugs []Models.Bug
	err = cursor.All(nil, &bugs)
	if err != nil {
		return nil, err
	}
	return bugs, nil
}

func (repo *bugsRepoType) InsertOne(bug *Models.Bug) error {
	inserted, err := repo.collection.InsertOne(nil, bug)
	if err != nil {
		bug.Id = inserted.InsertedID.(primitive.ObjectID)
	}
	return err
}
