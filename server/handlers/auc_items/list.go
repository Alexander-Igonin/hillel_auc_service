package aucitems

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ListResponseBody struct{}

// GetItemsListHandler godoc
// @Summary Show list of items
// @Description Show list of items
// @Produce json
// @Success 200 {object} ListResponseBody
// @Router /auc/list [get]
func (h *Handler) GetItemsListHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, ListResponseBody{})
}
