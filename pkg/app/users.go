package app

import (
	"mirror-backend/pkg"
	"mirror-backend/pkg/handlers/users"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *App) AttachUserRoutes() {
	a.engine.GET("/users/self", a.auth.User(), func(c *gin.Context) {
		email := c.GetString("email")
		user, err := users.GetUserSelf(c, a.repo, email)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	})

	a.engine.POST("/teams/api-keys", a.auth.User(), func(c *gin.Context) {
		email := c.GetString("email")
		user, err := users.GetUserSelf(c, a.repo, email)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err := users.CreateApiKey(c, a.repo, user); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	})
	a.engine.PUT("/teams/api-keys/:id", a.auth.User(), func(c *gin.Context) {
		email := c.GetString("email")
		user, err := users.GetUserSelf(c, a.repo, email)
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

		if err := users.UpdateApiKey(c, a.repo, user, newKey); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	})
	a.engine.DELETE("/teams/api-keys/:id", a.auth.User(), func(c *gin.Context) {
		email := c.GetString("email")
		user, err := users.GetUserSelf(c, a.repo, email)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		apiKeyID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := users.DeleteApiKey(c, a.repo, user, apiKeyID); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	})
}
