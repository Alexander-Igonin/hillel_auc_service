package aucitems

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AddResponseBody struct{}

// AddItemHandler godoc
// @Summary Add an item
// @Description Add an item to the auc
// @Produce json
// @Success 200 {object} AddResponseBody
// @Router /auc/new [post]
func (h *Handler) AddItemHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, AddResponseBody{})
}
