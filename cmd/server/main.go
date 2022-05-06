package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/egec-org/bookem/pkg/appointments"
	"github.com/egec-org/bookem/pkg/auth"
	"github.com/egec-org/bookem/internal/db"
	"github.com/egec-org/bookem/internal/config"
)

func initHttpServer(r *gin.Engine) {
	appointments.InitHttpHandlers(r)
	auth.InitHttpHandlers(r)
}

func main() {
	var appConfig *config.AppConfig
	appConfigEnvVar := os.Getenv("APP_CONFIG")

	if appConfigEnvVar != "" {
		appConfig = config.ReadConfigPath(appConfigEnvVar)
		fmt.Println(appConfig.Database.User)
		db.NewClient(appConfig.Database)
	} else {
		fmt.Println("Config path not set.")
		os.Exit(1)
	}
	r := gin.Default()


	initHttpServer(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
