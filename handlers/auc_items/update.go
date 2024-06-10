package aucitems

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UpdateRequestBody struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"item_name"`
	Price  int    `json:"price"`
}

// UpdateItemHandler godoc
// @Summary Update one item
// @Description Update one item
// @Produce json
// @Success 200 {object} UpdateResponseBody
// @Router /auc/update [put]
func (h *Handler) UpdateItemHandler(ctx echo.Context) error {
	newItem := UpdateRequestBody{}

	err := ctx.Bind(&newItem)
	if err != nil {
		return fmt.Errorf("failed to bind user %w", err)
	}

	sqlStatement := `
		UPDATE items
		SET user_id = $2, item_name = $3, price = $4
		WHERE id = $1;`

	_, err = h.clients.Postgres.Exec(sqlStatement, newItem.ID, newItem.UserID, newItem.Name, newItem.Price)
	if err != nil {
		return fmt.Errorf("failed to execute query %w", err)
	}
	return ctx.JSON(http.StatusOK, newItem)
}
