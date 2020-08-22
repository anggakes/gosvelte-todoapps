package handlers

import (
	"github.com/anggakes/gosvelte/src/backend/pkg/providers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoute(e *echo.Echo, prv *providers.Provider) {

	e.Use(middleware.CORS())

	todoHndlr := &TodoHandler{prv.Todo}
	todo := e.Group("/todos")
	todo.Add("GET", "/:id/", todoHndlr.Get)
	todo.Add("POST", "/", todoHndlr.Create)
	todo.Add("DELETE", "/:id/", todoHndlr.Delete)
	todo.Add("PUT", "/:id/", todoHndlr.Update)
	todo.Add("GET", "/", todoHndlr.List)

}
