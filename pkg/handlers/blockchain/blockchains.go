package blockchainhandlers

import (
	"log"
	"mirror-backend/pkg"
	"mirror-backend/pkg/services/blockchains"
	"mirror-backend/pkg/services/users"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateBlockchain(c *gin.Context, deps pkg.Dependencies) {
	email := c.GetString("email")
	user, err := users.GetUserSelf(c, deps.Repo, email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if len(user.Team.ApiKeys) == 0 {
		if err := users.CreateApiKey(c, deps.Repo, user); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		u, err := users.GetUserSelf(c, deps.Repo, email)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		user = u
	}

	blockchain, err := blockchains.CreateBlockchain(c, deps.RpcEngine, deps.Repo, user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, blockchain)
}

func EditBlockchain(c *gin.Context, deps pkg.Dependencies) {
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

	var newBlockchain pkg.Blockchain
	if err := c.BindJSON(&newBlockchain); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newBlockchain.ID = blockchainID
	if err := blockchains.UpdateBlockchain(c, deps.Repo, user, newBlockchain); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func DeleteBlockchain(c *gin.Context, deps pkg.Dependencies) {
	email := c.GetString("email")
	user, err := users.GetUserSelf(c, deps.Repo, email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(user.Team.ApiKeys) == 0 {
		if err := users.CreateApiKey(c, deps.Repo, user); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		u, err := users.GetUserSelf(c, deps.Repo, email)
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
	if err := blockchains.DeleteBlockchain(c, deps.RpcEngine, deps.Repo, user, blockchainID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Blockchain deleted"})
}

func CreateBlockchainSession(c *gin.Context, deps pkg.Dependencies) {
	key, _ := c.Get("key")
	apiKey := key.(pkg.ApiKey)

	userID := c.GetHeader("user_id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "user_id header is required"})
		return
	}

	blockchainID, err := blockchains.CreateBlockchainSession(c, deps.Repo, deps.RpcEngine, userID, apiKey)
	if err != nil {
		log.Println("Error creating session", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"url": "https://engine.mirror.ad/rpc/" + blockchainID.String(), "wsUrl": "wss://engine.mirror.ad/rpc/" + blockchainID.String()})
}
