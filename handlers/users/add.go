package users

import (
	"fmt"
	"hillel_go_auc/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AddUserHandler(ctx echo.Context) error {
	user := models.User{}

	err := ctx.Bind(&user)
	if err != nil {
		return fmt.Errorf("failed to bind user %w", err)
	}

	sqlStatement := `
		INSERT INTO users (name_full, email, phone, user_password)
		VALUES ($1, $2, $3, $4)`

	_, err = h.clients.Postgres.Exec(sqlStatement, user.NameFull, user.Email, user.Phone, user.Password)
	if err != nil {
		panic(err)
	}

	return ctx.JSON(http.StatusOK, user)
}
