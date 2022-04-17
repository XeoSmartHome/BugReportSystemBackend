package Repos

import (
	"BugReportSystemBackend/Database"
	"BugReportSystemBackend/Database/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepoType struct {
	collection *mongo.Collection
}

var UsersRepo = userRepoType{
	collection: Database.DB.Collection("User"),
}

func (repo *userRepoType) GetById(id primitive.ObjectID) (*Models.User, error) {
	var user Models.User
	filter := bson.M{
		"_id": id,
	}
	err := repo.collection.FindOne(nil, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepoType) GetByEmail(email string) (*Models.User, error) {
	var user Models.User
	filter := bson.M{
		"email": email,
	}
	err := repo.collection.FindOne(nil, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepoType) CreateUser(firstName string, lastName string, email string, password string, role string) (*Models.User, error) {
	user := Models.User{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PasswordHash: password,
		Role:         role,
	}
	result, err := repo.collection.InsertOne(nil, user)
	if err != nil {
		return nil, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return &user, err
}
