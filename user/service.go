package userModule

import (
	"context"
	"errors"
	"example/web-service-gin/database"
	"example/web-service-gin/utils"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (userservice UserService) GetOne(id string) (*User, error) {
	db := database.GetConnection().Collection("user")

	var user User
	err := db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return nil, err
	}

	if user.Id != "" {
		return nil, errors.New("Not found record!")
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

	if user.Id == "" {
		return nil, errors.New("Not found record!")
	}

	return &user, nil
}

func (userservice UserService) GetJwtToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	secretKey := utils.EnvVar("SECRET", "")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
