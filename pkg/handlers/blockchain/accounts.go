package blockchainhandlers

import (
	"mirror-backend/pkg"
	"mirror-backend/pkg/services/blockchains"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetMainnetAccountState(c *gin.Context, deps pkg.Dependencies) {
	blockchainId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blockchain ID"})
		return
	}

	type RequestBody struct {
		Accounts      []string `json:"accounts"`
		Label         *string  `json:"label,omitempty"`
		TokenMintAuth *string  `json:"token_mint_auth,omitempty"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if err := blockchains.SetMainnetAccountState(
		c,
		deps.RpcEngine,
		deps.AccountRetriever,
		blockchainId,
		requestBody.Accounts,
		requestBody.Label,
		requestBody.TokenMintAuth,
	); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Mainnet account state updated successfully"})
}

func SetProgramOwnedAccountState(c *gin.Context, deps pkg.Dependencies) {
	blockchainId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blockchain ID"})
		return
	}

	type RequestBody struct {
		Account string `json:"account"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if err := blockchains.SetProgramOwnedAccountState(
		c,
		deps.RpcEngine,
		deps.AccountRetriever,
		blockchainId,
		requestBody.Account,
	); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Mainnet account state updated successfully"})
}
