package middleware

import (
	"../service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			tokenString := authHeader[len(BEARER_SCHEMA):]
			if len(tokenString) > 0 {
				token, err := service.ValidateToken(tokenString)
				if token != nil && token.Valid {
					claims := token.Claims.(jwt.MapClaims)
					fmt.Println(claims)
				} else {
					fmt.Println(err)
					c.AbortWithStatus(http.StatusUnauthorized)
				}
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
