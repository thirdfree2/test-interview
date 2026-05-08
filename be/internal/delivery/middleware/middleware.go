package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("my-secret-key")

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "missing authorization header",
			})
			return
		}

		// Bearer xxxxx
		splitToken := strings.Split(authHeader, "Bearer ")

		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid authorization format",
			})
			return
		}

		tokenString := splitToken[1]

		// parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// check signing method
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}

		// get claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token claims",
			})
			return
		}

		// set user data into context
		c.Set("user_id", claims["user_id"])
		c.Set("name", claims["name"])

		c.Next()
	}
}