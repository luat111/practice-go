package userModule

import (
	"context"
	"errors"
	"example/web-service-gin/database"

	"github.com/goonode/mogo"
	"go.mongodb.org/mongo-driver/bson"
)

// Userservice is to handle user relation database query
type Userservice struct{}

// Create is to register new user
func (userservice Userservice) Create(user *(User)) error {
	db := database.GetConnection().Collection("user")

	result :=  db.FindOne(context.TODO(), bson.D{{"email", user.Email}})

	if result != nil {
		return errors.New("User existed!")
	}

	result, err := db.InsertOne(context.TODO(), user)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return vErr
	}
	return err
}

// Delete a user from database
func (userservice Userservice) Delete(email string) error {
	user, _ := userservice.FindByEmail(email)
	conn := database.GetConnection()
	defer conn.Session.Close()
	err := user.Remove()
	return err
}

// Find user
func (userservice Userservice) Find(user *(User)) (*User, error) {
	conn := database.GetConnection()
	defer conn.Session.Close()

	doc := mogo.NewDoc(User{}).(*(User))
	err := doc.FindOne(bson.M{"email": user.Email}, doc)

	if err != nil {
		return nil, err
	}
	return doc, nil
}

// Find user from email
func (userservice Userservice) FindByEmail(email string) (*User, error) {
	conn := database.GetConnection()
	defer conn.Session.Close()

	user := new(User)
	user.Email = email
	return userservice.Find(user)
}
