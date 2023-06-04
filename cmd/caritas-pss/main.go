package main

import (
	"log"

	"github.com/rostis232/psycho/internal/app/repository"
	"github.com/rostis232/psycho/internal/pkg/app"
	"github.com/spf13/viper"
)

func main() {

  if err := initConfig(); err != nil {
    log.Fatalln("error while db config loading")
  }
  dbconf := repository.Config{
    Host: viper.GetString("db.host"),
    Port: viper.GetString("db.port"),
    Username: viper.GetString("db.username"),
    Password: viper.GetString("db.password"),
    DBName: viper.GetString("db.dbname"),
    SSLMode: viper.GetString("db.sslmode"),
  }

  app := app.NewApp(dbconf)

  app.Run(viper.GetString("app.port"))
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
