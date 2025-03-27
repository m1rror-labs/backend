package app

import (
	"log"
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

		output, logs, err := codeexec.RunCode(c, request.Code, a.tsRuntime, a.repo)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error(), "output": output, "logs": logs})
			return
		}

		c.JSON(200, gin.H{"output": output, "logs": logs})
	})
	a.engine.POST("/code-exec/rust", func(c *gin.Context) {
		var request codeexec.ExecuteCodeRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		output, logs, err := codeexec.RunCode(c, request.Code, a.rustRuntime, a.repo)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"error": err.Error(), "output": output, "logs": logs})
			return
		}

		c.JSON(200, gin.H{"output": output, "logs": logs})
	})
}
