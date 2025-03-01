package jwt

import (
	"errors"
	"log"
	"mirror-backend/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type DefaultAuth struct {
	key []byte
}

func NewAuthMiddleware(key string) pkg.Auth {
	return &DefaultAuth{
		key: []byte(key),
	}
}

func (a *DefaultAuth) User() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken, err := getBearerToken(c)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		claims := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(bearerToken, &claims, func(t *jwt.Token) (interface{}, error) {
			return a.key, nil
		})
		if err != nil {
			log.Println("Error parsing claims", err)
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			log.Println("No email in claims", claims)
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		c.Set("email", email)
	}
}

func getBearerToken(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		token := c.Query("token")
		if token == "" {
			return "", errors.New("no auth header")
		}
		authHeader = token
	}
	authStrs := strings.Split(authHeader, "Bearer ")
	if len(authStrs) != 2 {
		return "", errors.New("doesn't include Bearer")
	}
	return authStrs[1], nil
}
