package aucitems

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DeleteResponseBody struct{}

// DeleteItemHandler godoc
// @Summary Delete one item
// @Description Delete one item
// @Produce json
// @Success 200 {object} DeleteResponseBody
// @Router /auc/delete [delete]
func (h *Handler) DeleteItemHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, DeleteResponseBody{})
}
