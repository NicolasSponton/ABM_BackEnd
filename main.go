package main

import (
	"ABM_Clientes/database"
	"ABM_Clientes/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	database.InitConnection()
	routes.HandleRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
