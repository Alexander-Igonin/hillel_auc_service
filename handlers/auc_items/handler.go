package aucitems

import "hillel_go_auc/clients"

type Handler struct {
	clients *clients.Clients
}

func NewHandler(clients *clients.Clients) *Handler {
	return &Handler{
		clients: clients,
	}
}
