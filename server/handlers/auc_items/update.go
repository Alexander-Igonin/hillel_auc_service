package aucitems

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UpdateResponseBody struct{}

// UpdateItemHandler godoc
// @Summary Update one item
// @Description Update one item
// @Produce json
// @Success 200 {object} UpdateResponseBody
// @Router /auc/update [put]
func (h *Handler) UpdateItemHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, UpdateResponseBody{})
}
