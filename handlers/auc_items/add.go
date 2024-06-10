package aucitems

import (
	"fmt"
	"hillel_go_auc/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AddItemHandler godoc
// @Summary Add an item
// @Description Add an item to the auc
// @Produce json
// @Success 200 {object} AddResponseBody
// @Router /auc/new [post]
func (h *Handler) AddItemHandler(ctx echo.Context) error {
	item := models.Item{}

	err := ctx.Bind(&item)
	if err != nil {
		return fmt.Errorf("failed to bind user %w", err)
	}

	sqlStatement := `
		INSERT INTO items (user_id, item_name, price)
		VALUES ($1, $2, $3)`

	_, err = h.clients.Postgres.Exec(sqlStatement, item.UserID, item.Name, item.Price)
	if err != nil {
		return fmt.Errorf("failed to add item: %w", err)
	}

	return ctx.JSON(http.StatusOK, item)
}
