package app

import (
	"mirror-backend/pkg"
	"mirror-backend/pkg/handlers/blockchains"
	"mirror-backend/pkg/handlers/users"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *App) AttachBlockchainRoutes() {
	a.engine.POST("/blockchains", a.auth.User(), func(c *gin.Context) {
		email := c.GetString("email")
		user, err := users.GetUserSelf(c, a.repo, email)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if len(user.Team.ApiKeys) == 0 {
			if err := users.CreateApiKey(c, a.repo, user); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			u, err := users.GetUserSelf(c, a.repo, email)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			user = u
		}

		blockchain, err := blockchains.CreateBlockchain(c, a.rpcEngine, a.repo, user)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, blockchain)
	})
	a.engine.PUT("/blockchains/:id", a.auth.User(), func(c *gin.Context) {
		email := c.GetString("email")
		user, err := users.GetUserSelf(c, a.repo, email)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		blockchainID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var newBlockchain pkg.Blockchain
		if err := c.BindJSON(&newBlockchain); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		newBlockchain.ID = blockchainID
		if err := blockchains.UpdateBlockchain(c, a.repo, user, newBlockchain); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	})
	a.engine.DELETE("/blockchains/:id", a.auth.User(), func(c *gin.Context) {
		email := c.GetString("email")
		user, err := users.GetUserSelf(c, a.repo, email)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if len(user.Team.ApiKeys) == 0 {
			if err := users.CreateApiKey(c, a.repo, user); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			u, err := users.GetUserSelf(c, a.repo, email)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			user = u
		}

		blockchainID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := blockchains.DeleteBlockchain(c, a.rpcEngine, a.repo, user, blockchainID); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Blockchain deleted"})
	})

	a.engine.GET("/blockchains/:id/transactions/logs", a.auth.User(), func(c *gin.Context) {
		email := c.GetString("email")
		user, err := users.GetUserSelf(c, a.repo, email)
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
		logs, count, err := blockchains.GetTransactionLogs(c, a.repo, blockchainID, user, page, limit)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"logs":  logs,
			"count": count,
		})

	})
}
