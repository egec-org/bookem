package main

import (
	"fmt"
	"os"

	"github.com/egec-org/bookem/internal/cache"
	"github.com/egec-org/bookem/internal/config"
	"github.com/egec-org/bookem/internal/db"
	"github.com/egec-org/bookem/pkg/appointments"
	"github.com/egec-org/bookem/pkg/auth"
	"github.com/gin-gonic/gin"
)

func initHttpServer(r *gin.Engine) {
	appointments.InitHttpHandlers(r)
	auth.InitHttpHandlers(r)
}

func main() {
	appConfigEnvVar := os.Getenv("APP_CONFIG")
	cache.BasicUsage()
	if appConfigEnvVar != "" {
		appConfig, err := config.ReadConfigPath(appConfigEnvVar)
		if err != nil {
			fmt.Println("Error looking for DB config")
			os.Exit(1)
		}
		fmt.Println(db.NewClient(appConfig.Database))
	} else {
		os.Exit(1)
	}
	r := gin.Default()

	initHttpServer(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
