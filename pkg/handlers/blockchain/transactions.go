package blockchainhandlers

import (
	"mirror-backend/pkg"
	"mirror-backend/pkg/services/blockchains"
	"mirror-backend/pkg/services/users"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTransactionLogs(c *gin.Context, deps pkg.Dependencies) {
	email := c.GetString("email")
	user, err := users.GetUserSelf(c, deps.Repo, email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	blockchainID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	logs, count, err := blockchains.GetTransactionLogs(c, deps.Repo, blockchainID, user, page, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"logs":  logs,
		"count": count,
	})
}
