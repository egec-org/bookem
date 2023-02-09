package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/egec-org/bookem/internal/db"
)

type AppConfig struct {
	Database db.DBConfig `toml:"database"`
}

func ReadConfigPath(path string) (*AppConfig, error) {
	var appConfig AppConfig
	if strings.HasPrefix(path, "s3://") {
		fmt.Println("Pulling config file from s3")
		return &appConfig, nil
	}
	if _, err := os.Stat(path); err == nil {
		fmt.Println("File exists")
		_, err = toml.DecodeFile(path, &appConfig)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Unable to parse toml file")
			return &appConfig, err
		}
		return &appConfig, nil
	} else {
		msg := "File path invalid"
		fmt.Println(msg)
		return &appConfig, errors.New(msg)
	}
}
