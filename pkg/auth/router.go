package auth

import (
  "github.com/gin-gonic/gin"
)

func InitHttpHandlers(r *gin.Engine) {
  r.POST("/access/token", AuthUser)
}
