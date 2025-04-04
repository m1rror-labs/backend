package userhandlers

import (
	"mirror-backend/pkg"
	"mirror-backend/pkg/services/users"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetApiKeys(c *gin.Context, deps pkg.Dependencies) {
	email := c.GetString("email")
	user, err := users.GetUserSelf(c, deps.Repo, email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := users.CreateApiKey(c, deps.Repo, user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func UpdateApiKey(c *gin.Context, deps pkg.Dependencies) {
	email := c.GetString("email")
	user, err := users.GetUserSelf(c, deps.Repo, email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	apiKeyID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var newKey pkg.ApiKey
	if err := c.BindJSON(&newKey); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	newKey.ID = apiKeyID

	if err := users.UpdateApiKey(c, deps.Repo, user, newKey); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func DeleteApiKey(c *gin.Context, deps pkg.Dependencies) {
	email := c.GetString("email")
	user, err := users.GetUserSelf(c, deps.Repo, email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	apiKeyID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := users.DeleteApiKey(c, deps.Repo, user, apiKeyID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}
