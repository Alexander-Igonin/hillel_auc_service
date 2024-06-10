package aucitems

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ListResponseBody struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"item_name"`
	Price  int    `json:"price"`
}

// GetItemsListHandler godoc
// @Summary Show list of items
// @Description Show list of items
// @Produce json
// @Success 200 {object} ListResponseBody
// @Router /auc/list [get]
func (h *Handler) GetItemsListHandler(ctx echo.Context) error {
	rows, err := h.clients.Postgres.Query("SELECT * FROM items")
	if err != nil {
		return fmt.Errorf("failed to query item: %w", err)
	}
	defer rows.Close()

	var items []ListResponseBody

	for rows.Next() {
		var item ListResponseBody
		if err := rows.Scan(&item.ID, &item.UserID, &item.Name, &item.Price); err != nil {
			return fmt.Errorf("failed to scan item: %w", err)
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("some error in queried rows: %w", err)
	}

	return ctx.JSON(http.StatusOK, items)
}
