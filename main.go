// Web application written in Go & ELM used to manage a photo album
package main

import (
	c "github.com/HETIC-MT-P2021/DB_RAYER_P01/config"
	db "github.com/HETIC-MT-P2021/DB_RAYER_P01/database"
	r "github.com/HETIC-MT-P2021/DB_RAYER_P01/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// main launch all part of the project
func main() {

	c.InitConfig()

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
