package userModule

import (
	"context"
	"errors"
	"example/web-service-gin/database"
	"example/web-service-gin/utils"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Create(user *User) (*User, error)
	GetOne(id string) (*User, error)
	FindByEmail(email string) (*User, error)
}

type UserService struct{}

func (userservice UserService) Create(user *(User)) (*User, error) {
	db := database.GetConnection().Collection("user")

	var checkExisted User
	err := db.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&checkExisted)

	if checkExisted.Id != "" {
		return nil, errors.New("User existed!")
	}

	result, err := db.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	user.Id = result.InsertedID.(primitive.ObjectID).Hex()

	return user, nil
}

func (userservice UserService) GetOne(id primitive.ObjectID) (*User, error) {
	db := database.GetConnection().Collection("user")

	var user User
	err := db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userservice UserService) FindByEmail(email string) (*User, error) {
	db := database.GetConnection().Collection("user")

	var user User
	err := db.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userservice UserService) Update(id primitive.ObjectID, payload *User) (*User, error) {
	db := database.GetConnection().Collection("user")

	var checkExisted User
	err := db.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&checkExisted)

	if err != nil {
		return nil, err
	}

	if payload.Password == "" {
		payload.Password = checkExisted.Password
	} else {
		hashStr, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.MinCost)
		if err != nil {
			return nil, err
		}

		payload.Password = string(hashStr)
	}

	if payload.Email == "" {
		payload.Email = checkExisted.Email
	}

	if payload.Name == "" {
		payload.Name = checkExisted.Name
	}
	
	update := bson.D{{Key: "$set", Value: bson.D{
		primitive.E{Key: "email", Value: payload.Email},
		primitive.E{Key: "name", Value: payload.Name},
		primitive.E{Key: "password", Value: payload.Password},
	}}}

	_, updatErr := db.UpdateOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}, update)
	if updatErr != nil {
		return nil, updatErr
	}

	payload.Id = checkExisted.Id

	return payload, nil
}

func (userservice UserService) Remove(id primitive.ObjectID) (string, error) {
	db := database.GetConnection().Collection("user")

	_, err := db.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}})

	if err != nil {
		return "Some thing wrong", err
	} else {
		return "Delete succeed", nil
	}
}

func (userservice UserService) GetJwtToken(email string, id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
	})
	secretKey := utils.EnvVar("SECRET", "")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
