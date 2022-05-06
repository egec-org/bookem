package config

import (
  "strings"
  "fmt"
  "os"

  "github.com/BurntSushi/toml"
  "github.com/egec-org/bookem/internal/db"
)

type AppConfig struct {
  Database db.DBConfig `toml:"database"`
}

func ReadConfigPath (path string) *AppConfig {
  var appConfig AppConfig
  if strings.HasPrefix(path, "s3://") {
    fmt.Println("Pulling config file from s3")
  }
  if _, err := os.Stat(path); err == nil {
    fmt.Println("File exists")
    _, err = toml.DecodeFile(path, &appConfig)
    if err != nil {
      fmt.Println(err)
      fmt.Println("Unable to parse toml file")
    }
    return &appConfig
  } else {
    fmt.Println("File path invalid.")
  }
  return &appConfig

}
