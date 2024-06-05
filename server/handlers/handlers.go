package handlers

import (
	"hillel_go_auc/clients"
	aucitems "hillel_go_auc/server/handlers/auc_items"
	"hillel_go_auc/server/handlers/users"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Auc  *aucitems.Handler
	User *users.Handler
}

func NewHandlers(clients *clients.Clients) *Handlers {
	return &Handlers{
		Auc:  aucitems.NewHandler(clients),
		User: users.NewHandler(clients),
	}
}

func (h Handlers) RegisterRouts(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	aucApi := e.Group("/auc")

	aucApi.GET("/item/:id", h.Auc.GetItemHandler)
	aucApi.GET("/list", h.Auc.GetItemsListHandler)
	aucApi.POST("/new", h.Auc.AddItemHandler)
	aucApi.PUT("/update", h.Auc.UpdateItemHandler)
	aucApi.DELETE("/delete", h.Auc.DeleteItemHandler)

	userApi := e.Group("/user")

	userApi.POST("/add", h.User.AddUserHandler)
}
