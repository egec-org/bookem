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
	"github.com/sarulabs/di"
	"go.uber.org/zap"
)

func initHttpServer(r *gin.Engine) {
	appointments.InitHttpHandlers(r)
	auth.InitHttpHandlers(r)
}

func createApp(
	config *config.AppConfig,
) di.Container {
	fmt.Println(config)
	builder, _ := di.NewBuilder()
	err := builder.Add(di.Def{
		Name:  "db-client",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return db.NewClient(config.Database)
		},
	})
	if err != nil {
		zap.S().Panic(err)
	}
	return builder.Build()
}

func main() {
	appConfigEnvVar := os.Getenv("APP_CONFIG")
	cache.BasicUsage()
	//var appContainer di.Container
	if appConfigEnvVar != "" {
		appConfig, err := config.ReadConfigPath(appConfigEnvVar)
		if err != nil {
			fmt.Println("Error looking for DB config")
			os.Exit(1)
		}
		fmt.Println(db.NewClient(appConfig.Database))

		//appContainer = createApp(appConfig)
		//fmt.Print(appContainer.Get("db-client"))
		//defer appContainer.Delete()
	} else {
		fmt.Println("Config path not set.")
		os.Exit(1)
	}
	r := gin.Default()

	initHttpServer(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
