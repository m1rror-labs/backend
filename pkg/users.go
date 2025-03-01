package pkg

import "github.com/gin-gonic/gin"

type Auth interface {
	User() gin.HandlerFunc
}
