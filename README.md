# bookem

This service is built on Go. It provides the backend APIs for the
stylist and customers.

A toml config file is needed for the DB. You can follow this file below.
```
[database]
host = "127.0.0.1"
user = "root"
password = "password"
name = "sample"

```

You will need to set this env var `APP_CONFIG` to point to the config file path.
For example, if your config file lives in etc/app_config.toml, then the env var
will be set like this.
`export APP_CONFIG=etc/app_config.toml`

You can run the API server by running
go run cmd/server/main.go
