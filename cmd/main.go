package main

import (
	"hillel_go_auc/clients"
	"hillel_go_auc/config"
	"hillel_go_auc/docs"
	"hillel_go_auc/server/handlers"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// @title           AucServiceAPI
// @version         1.0
// @description     hillel auc project.
// @termsOfService  http://swagger.io/terms/

// @BasePath /
func main() {
	docs.SwaggerInfo.Host = ""

	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	server := echo.New()
	clients := clients.NewClients()
	handlers := handlers.NewHandlers(clients)
	handlers.RegisterRouts(server)

	err = server.Start(cfg.Port)
	if err != nil {
		logrus.Fatal(err)
	}
}
