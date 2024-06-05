package users

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type AddUserRequestBody struct {
	NameFull string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type AddUserResponseBody struct{}

func (h *Handler) AddUserHandler(ctx echo.Context) error {
	user := AddUserRequestBody{}
	err := ctx.Bind(&user)
	if err != nil {
		return fmt.Errorf("failed to bind user %w", err)
	}

	return nil
}
