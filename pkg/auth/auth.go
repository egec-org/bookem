package auth

import (
  "github.com/gin-gonic/gin"
  "go.uber.org/zap"
)

func AuthUser(c *gin.Context) {
  zap.S().Info("Generate token.")
  c.JSON(201, gin.H{
    "message": "Here is your token X",
  })
}
