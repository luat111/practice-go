package userModule

import (
	// "time"
)

//User struct is to handle user data
type User struct {	
	Id                 string `bson:"_id,omitempty"`
	Email              string `bson:"email,omitempty"`
	Password           string `bson:"password,omitempty"`
	Name               string `bson:"name,omitempty"`
	// CreatedAt          *time.Time
	// UpdatedAt          *time.Time
	// VerifiedAt 		   *time.Time
}
