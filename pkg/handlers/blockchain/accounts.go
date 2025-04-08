package blockchainhandlers

import (
	"encoding/base64"
	"mirror-backend/pkg"
	"mirror-backend/pkg/services/blockchains"
	"mirror-backend/pkg/services/users"
	"strconv"
	"time"

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

func SetAccountStateFromRecentTransactions(c *gin.Context, deps pkg.Dependencies) {
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

	if err := blockchains.SetAccountStateFromRecentTransactions(
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

type AccountBase64 struct {
	ID           uuid.UUID `bun:"id,pk,type:uuid"`
	CreatedAt    time.Time `json:"created_at"`
	BlockchainID uuid.UUID `json:"blockchain_id"`
	Address      string    `json:"address"`
	Lamports     uint      `json:"lamports"`
	Data         string    `json:"data"`
	Owner        string    `json:"owner"`
	Executable   bool      `json:"executable"`
	RentEpoch    uint      `json:"rentEpoch"`
}

func GetAccounts(c *gin.Context, deps pkg.Dependencies) {
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
	if !blockchains.UserCanAccessBlockchain(user, blockchainID, deps.Repo) {
		c.JSON(403, gin.H{"error": "User does not have access to this blockchain"})
		return
	}

	accounts, count, err := blockchains.GetAccounts(c, deps.Repo, blockchainID, page, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	accountsBase64 := []AccountBase64{}
	for _, account := range accounts {
		data := base64.StdEncoding.EncodeToString(account.Data)
		accountsBase64 = append(accountsBase64, AccountBase64{
			ID:           account.ID,
			CreatedAt:    account.CreatedAt,
			BlockchainID: account.BlockchainID,
			Address:      account.Address,
			Lamports:     account.Lamports,
			Data:         data,
			Owner:        account.Owner,
			Executable:   account.Executable,
			RentEpoch:    account.RentEpoch,
		})
	}

	c.JSON(200, gin.H{
		"accounts": accountsBase64,
		"count":    count,
	})
}
