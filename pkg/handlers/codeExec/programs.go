package codeexechandlers

import (
	"log"
	"mirror-backend/pkg"
	codeexec "mirror-backend/pkg/services/codeExec"

	"github.com/gin-gonic/gin"
)

func BuildAndDeployAnchor(c *gin.Context, deps pkg.Dependencies) {
	var request codeexec.BuildProgramRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := codeexec.BuildAndLoadProgram(
		c,
		request.Code,
		request.ProgramID,
		request.BlockchainID,
		deps.AnchorRuntime,
		deps.RpcEngine,
	)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Program built and loaded successfully"})
}
