package main

import (
	"time"

	"github.com/adimurianto/go-gin-boilerplate/config"
	"github.com/adimurianto/go-gin-boilerplate/infra/database"
	"github.com/adimurianto/go-gin-boilerplate/infra/logger"
	"github.com/adimurianto/go-gin-boilerplate/migrations"
	"github.com/adimurianto/go-gin-boilerplate/routers"

	_ "github.com/adimurianto/go-gin-boilerplate/docs"

	"github.com/spf13/viper"
)

// @title						MASTER API
// @version					1.0
// @in							header
// @name						Authorization
func main() {
	// set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Jakarta")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
	psqlDBConf := config.DbConfiguration()

	if err := database.DbConnection(psqlDBConf); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	migrations.Migrate()

	router := routers.SetupRoute()
	logger.Fatalf("%v", router.Run(config.ServerConfig()))
}
