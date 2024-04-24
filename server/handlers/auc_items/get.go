package aucitems

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetResponseBody struct{}

// GetItemHandler godoc
// @Summary Show one item
// @Description Show one item
// @Produce json
// @Success 200 {object} GetResponseBody
// @Router /auc/item/:id [get]
func (h *Handler) GetItemHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, GetResponseBody{})
}
