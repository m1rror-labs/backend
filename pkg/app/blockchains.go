package app

import (
	"context"
	"mirror-backend/pkg"
	blockchainhandlers "mirror-backend/pkg/handlers/blockchain"
	"mirror-backend/pkg/services/blockchains"

	"github.com/gin-gonic/gin"
)

func (a *App) AttachBlockchainRoutes() {
	a.engine.POST("/blockchains", a.deps.Auth.User(), func(c *gin.Context) {
		blockchainhandlers.CreateBlockchain(c, a.deps)
	})
	a.engine.PUT("/blockchains/:id", a.deps.Auth.User(), func(c *gin.Context) {
		blockchainhandlers.EditBlockchain(c, a.deps)
	})
	a.engine.DELETE("/blockchains/:id", a.deps.Auth.User(), func(c *gin.Context) {
		blockchainhandlers.DeleteBlockchain(c, a.deps)
	})
	a.engine.POST("/blockchains/sessions", a.deps.Auth.ApiKey(a.deps.Repo), func(c *gin.Context) {
		blockchainhandlers.CreateBlockchainSession(c, a.deps)
	})

	a.engine.GET("/blockchains/:id/transactions/logs", a.deps.Auth.User(), func(c *gin.Context) {
		blockchainhandlers.GetTransactionLogs(c, a.deps)
	})

	a.engine.GET("/blockchains/:id/accounts", a.deps.Auth.User(), func(c *gin.Context) {
		blockchainhandlers.GetAccounts(c, a.deps)
	})
	a.engine.POST("/blockchains/:id/accounts/mainnet", a.deps.Auth.Team(a.deps.Repo), func(c *gin.Context) {
		blockchainhandlers.SetMainnetAccountState(c, a.deps)
	})
	// a.engine.POST("/blockchains/:id/accounts/mainnet/program-owned-accounts", a.deps.Auth.User(), func(c *gin.Context) {
	// 	blockchainhandlers.SetProgramOwnedAccountState(c, a.deps)
	// })

	if a.env != "dev" {
		go pkg.ProtectedFunc(func() {
			blockchains.ExpireBlockchains(context.Background(), a.deps.RpcEngine)
		})
	}
}
