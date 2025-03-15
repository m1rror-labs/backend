package app

import (
	codeexec "mirror-backend/pkg/handlers/code-exec"

	"github.com/gin-gonic/gin"
)

func (a *App) AttachCodeExecRoutes() {
	a.engine.POST("/code-exec/typescript", func(c *gin.Context) {
		var request codeexec.ExecuteCodeRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		output, err := codeexec.RunCode(c, request.Code, a.tsRuntime)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"output": output})
	})
}
