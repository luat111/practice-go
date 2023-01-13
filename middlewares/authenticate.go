package middlewares

import (
	"example/web-service-gin/user"
	"example/web-service-gin/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Authentication is for auth middleware
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authentication")
		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "Authentication header is missing",
			})
			return
		}
		temp := strings.Split(authHeader, "Bearer")
		if len(temp) < 2 {
			c.AbortWithStatusJSON(400, gin.H{"error": "Invalid token"})
			return
		}
		tokenString := strings.TrimSpace(temp[1])
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// 	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			// }
			secretKey := utils.EnvVar("TOKEN_KEY", "")
			return []byte(secretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
			userservice := userModule.Userservice{}
			user, err := userservice.FindByEmail(email)
			if err != nil {
				c.AbortWithStatusJSON(402, gin.H{
					"error": "User not found",
				})
				return
			}
			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "Token is not valid",
			})
			return
		}
	}
}
