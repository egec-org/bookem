package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"time"
)

func ValidateUser(c *gin.Context) {
	zap.S().Info("Generate token.")
	token, err := generateJWT()
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"message": "Unable to get token",
		})
		return
	}
	c.JSON(201, gin.H{
		"token": token,
	})
}

func generateJWT() (string, error) {
	var sampleSecretKey = []byte("SecretYouShouldHide")
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
