package aucitems

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DeleteItemHandler godoc
// @Summary Delete one item
// @Description Delete one item
// @Produce json
// @Success 200 {object} DeleteResponseBody
// @Router /auc/delete [delete]
func (h *Handler) DeleteItemHandler(ctx echo.Context) error {
	itemID := ctx.Param("id")

	sqlStatement := `
		DELETE FROM items
		WHERE id = $1;`

	_, err := h.clients.Postgres.Exec(sqlStatement, itemID)
	if err != nil {
		return fmt.Errorf("failed to delete info: %w", err)
	}

	return ctx.NoContent(http.StatusOK)
}
