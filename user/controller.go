package userModule

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func (userController *UserController) Login(c *gin.Context) {

	var loginInfo User
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	//TODO
	userservice := UserService{}
	user, errf := userservice.FindByEmail(loginInfo.Email)
	if errf != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		c.AbortWithStatusJSON(402, gin.H{"error": "Email or password is invalid."})
		return
	}

	token, err := userservice.GetJwtToken(user.Email, user.Id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	//-------
	c.JSON(200, gin.H{
		"token": token,
	})
}

// Profile is to provide current user info
func (userController *UserController) Profile(c *gin.Context) {
	user := c.MustGet("user").(*(User))

	c.JSON(200, gin.H{
		"id":        user.Id,
		"user_name": user.Name,
		"email":     user.Email,
	})
}

type SignupInfo struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
}

// Signup is for user signup
func (userController *UserController) Signup(c *gin.Context) {

	var info SignupInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}

	user := User{}
	user.Email = info.Email
	hash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
		return
	}

	user.Password = string(hash)
	user.Name = info.Name
	userservice := UserService{}

	result, err := userservice.Create(&user)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, result)
	}
	return
}

func (userController *UserController) Update(c *gin.Context) {
	id, convertErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convertErr != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid ID"})
	}

	var payloadUpdate User
	if err := c.ShouldBindJSON(&payloadUpdate); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	userservice := UserService{}
	updatedUser, updateErr := userservice.Update(id, &payloadUpdate)
	if updateErr != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": updateErr.Error()})
		return
	} else {
		c.JSON(200, updatedUser)
	}

	return
}

func (userController *UserController) Remove(c *gin.Context) {
	id, convertErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convertErr != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid ID"})
	}

	userservice := UserService{}
	result, delErr := userservice.Remove(id)
	if delErr != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": delErr.Error()})
		return
	} else {
		c.JSON(200, result)
	}

	return
}
