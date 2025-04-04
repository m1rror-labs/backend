package app

import (
	userhandlers "mirror-backend/pkg/handlers/users"

	"github.com/gin-gonic/gin"
)

func (a *App) AttachUserRoutes() {
	a.engine.GET("/users/self", a.deps.Auth.User(), func(c *gin.Context) {
		userhandlers.GetUserSelf(c, a.deps)
	})

	a.engine.POST("/teams/api-keys", a.deps.Auth.User(), func(c *gin.Context) {
		userhandlers.GetApiKeys(c, a.deps)
	})
	a.engine.PUT("/teams/api-keys/:id", a.deps.Auth.User(), func(c *gin.Context) {
		userhandlers.UpdateApiKey(c, a.deps)
	})
	a.engine.DELETE("/teams/api-keys/:id", a.deps.Auth.User(), func(c *gin.Context) {
		userhandlers.DeleteApiKey(c, a.deps)
	})
}
