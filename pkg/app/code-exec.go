package app

import "github.com/gin-gonic/gin"

func (a *App) AttachCodeExecRoutes() {
	a.engine.POST("/code-exec/typescript", func(c *gin.Context) {
		var request struct {
			Code string `json:"code" binding:"required"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		output, err := a.tsRuntime.ExecuteCode(request.Code)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"output": output})
	})
}