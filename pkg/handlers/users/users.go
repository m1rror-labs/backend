package userhandlers

import (
	"mirror-backend/pkg"
	"mirror-backend/pkg/services/users"

	"github.com/gin-gonic/gin"
)

func GetUserSelf(c *gin.Context, deps pkg.Dependencies) {
	email := c.GetString("email")
	user, err := users.GetUserSelf(c, deps.Repo, email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
