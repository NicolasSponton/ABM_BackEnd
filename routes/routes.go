package routes

import (
	"ABM_Clientes/controllers/clientes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HandleRoutes(e *echo.Echo) {

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	r := e.Group("/abm")

	r.GET("/clientes", clientes.GetAll)
	r.GET("/clientes/search/:nombre", clientes.Search)
	r.GET("/clientes/:id", clientes.Get)
	r.POST("/clientes", clientes.Create)
	r.PUT("/clientes", clientes.Update)
	r.DELETE("/clientes/:id", clientes.Delete)

}
