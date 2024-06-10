package aucitems

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetResponseBody struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"item_name"`
	Price  int    `json:"price"`
}

// GetItemHandler godoc
// @Summary Show one item
// @Description Show one item
// @Produce json
// @Success 200 {object} GetResponseBody
// @Router /auc/item/:id [get]
func (h *Handler) GetItemHandler(ctx echo.Context) error {
	itemID := ctx.Param("id")

	row := h.clients.Postgres.QueryRow(`SELECT * FROM items WHERE id=$1;`, itemID)

	var item GetResponseBody

	err := row.Scan(&item.ID, &item.UserID, &item.Name, &item.Price)
	if err == sql.ErrNoRows {
		return fmt.Errorf("no items found")
	}
	if err != nil {
		return fmt.Errorf("failed to scan item: %w", err)
	}

	return ctx.JSON(http.StatusOK, item)
}
