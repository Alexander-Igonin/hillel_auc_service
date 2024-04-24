package server

import (
	"hillel_go_auc/server/handlers"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e        echo.Echo
	handlers handlers.Handlers
}

func NewServer() *Server {
	return &Server{
		e:        *echo.New(),
		handlers: *handlers.NewHandlers(),
	}
}
