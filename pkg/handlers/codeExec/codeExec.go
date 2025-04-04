package codeexechandlers

import (
	"log"
	"mirror-backend/pkg"
	codeexec "mirror-backend/pkg/services/codeExec"

	"github.com/gin-gonic/gin"
)

func ExecuteTypescript(c *gin.Context, deps pkg.Dependencies) {
	var request codeexec.ExecuteCodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	output, logs, err := codeexec.RunCode(c, request.Code, deps.TsRuntime, deps.Repo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "output": output, "logs": logs})
		return
	}

	c.JSON(200, gin.H{"output": output, "logs": logs})
}

func ExecuteRust(c *gin.Context, deps pkg.Dependencies) {
	var request codeexec.ExecuteCodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	output, logs, err := codeexec.RunCode(c, request.Code, deps.RustRuntime, deps.Repo)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": err.Error(), "output": output, "logs": logs})
		return
	}

	c.JSON(200, gin.H{"output": output, "logs": logs})
}
