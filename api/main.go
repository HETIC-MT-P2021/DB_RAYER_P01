// Web application written in Go & ELM used to manage a photo album
package main

import (
	"fmt"

	r "github.com/HETIC-MT-P2021/DB_RAYER_P01/api/router"

	"github.com/HETIC-MT-P2021/DB_RAYER_P01/api/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

// initConfig set all configuration for the project
func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./api")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

// main launch all part of the project
func main() {
	// setup env
	initConfig()

	db.Connect()

	e := echo.New()

	r.InitRoutes(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.RequestID())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":80"))
}
