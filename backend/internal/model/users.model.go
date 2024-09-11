package model

import (
	"backend/internal/genproto/users"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in MongoDB
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Surname  string             `bson:"surname"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

func ConvertMongoToGrpc(user User) (*users.User, error) {
	id := user.ID.Hex() // Convert ObjectID to hex string
	return &users.User{
		Id:       id,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func ConvertGrpcToMongo(user *users.User) (*User, error) {
	objectID, err := primitive.ObjectIDFromHex(user.Id) // Convert string back to ObjectID
	if err != nil {
		return nil, errors.New("invalid ObjectID format")
	}
	return &User{
		ID:       objectID,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
