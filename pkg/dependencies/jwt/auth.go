package jwt

import (
	"errors"
	"log"
	"mirror-backend/pkg"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type DefaultAuth struct {
	key              []byte
	approvedCodeExec []string
}

func NewAuthMiddleware(key string, approvedCodeExec []string) pkg.Auth {
	return &DefaultAuth{
		key:              []byte(key),
		approvedCodeExec: approvedCodeExec,
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

func (a *DefaultAuth) ApiKey(userRepo pkg.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyStr := c.GetHeader("api_key")
		apiKey, err := uuid.Parse(apiKeyStr)
		if err != nil {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
		key, err := userRepo.ReadApiKey().ID(apiKey).ExecuteOne(c)
		if err != nil {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		c.Set("key", key)
	}
}

func (a *DefaultAuth) Team(userRepo pkg.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyStr := c.GetHeader("api_key")
		if apiKeyStr != "" {
			apiKey, err := uuid.Parse(apiKeyStr)
			if err != nil {
				c.Status(http.StatusForbidden)
				c.Abort()
				return
			}
			key, err := userRepo.ReadApiKey().ID(apiKey).WithTeam().ExecuteOne(c)
			if err != nil {
				c.Status(http.StatusForbidden)
				c.Abort()
				return
			}

			c.Set("team", *key.Team)
			return
		}

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

		user, err := userRepo.ReadUser().Email(email).WithTeam().ExecuteOne(c)
		if err != nil {
			log.Println("Error getting user", err)
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		c.Set("team", *user.Team)
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

func (a *DefaultAuth) CodeExec() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyStr := c.GetHeader("api_key")
		if !slices.Contains(a.approvedCodeExec, apiKeyStr) {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
	}
}
