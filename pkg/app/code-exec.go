package app

import (
	codeexechandlers "mirror-backend/pkg/handlers/codeExec"

	"github.com/gin-gonic/gin"
)

func (a *App) AttachCodeExecRoutes() {
	a.engine.POST("/code-exec/typescript", a.deps.Auth.CodeExec(), func(c *gin.Context) {
		codeexechandlers.ExecuteTypescript(c, a.deps)
	})
	a.engine.POST("/code-exec/rust", a.deps.Auth.CodeExec(), func(c *gin.Context) {
		codeexechandlers.ExecuteRust(c, a.deps)
	})

	a.engine.POST("/code-exec/programs/anchor", a.deps.Auth.CodeExec(), func(c *gin.Context) {
		codeexechandlers.BuildAndDeployAnchor(c, a.deps)
	})
	a.engine.POST("/code-exec/programs/anchor/test", a.deps.Auth.CodeExec(), func(c *gin.Context) {
		codeexechandlers.BuildAndTestAnchor(c, a.deps)
	})

	// a.engine.POST("/code-exec/load-test", func(c *gin.Context) {
	// 	codeexechandlers.LoadTest(c, a.deps)
	// })
}
