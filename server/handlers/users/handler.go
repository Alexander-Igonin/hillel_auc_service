package users

import "hillel_go_auc/clients"

type Handler struct {
}

func NewHandler(clients *clients.Clients) *Handler {
	return &Handler{}
}
