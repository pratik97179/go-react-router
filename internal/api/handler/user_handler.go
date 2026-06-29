package handler

import (
	"errors"
	"net/http"

	"go-react-router/internal/api/response"
	"go-react-router/internal/domain/user"
	"go-react-router/internal/domain/user/usecase/commands"

	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	registerUser *commands.RegisterUserUseCase
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(
	registerUser *commands.RegisterUserUseCase,
) *UserHandler {
	return &UserHandler{
		registerUser: registerUser,
	}
}

// Register handles user registration.
func (h *UserHandler) Register(
	c *gin.Context,
) {
	var registration user.Registration

	if err := c.ShouldBindJSON(&registration); err != nil {
		response.JSON(
			c,
			http.StatusBadRequest,
			false,
			err.Error(),
			nil,
		)
		return
	}

	err := h.registerUser.Execute(
		c.Request.Context(),
		registration,
	)

	if err != nil {

		switch {

		case errors.Is(err, user.ErrEmailAlreadyExists):
			response.JSON(
				c,
				http.StatusConflict,
				false,
				err.Error(),
				nil,
			)

		case errors.Is(err, user.ErrEmailRequired),
			errors.Is(err, user.ErrInvalidEmail),
			errors.Is(err, user.ErrPasswordRequired),
			errors.Is(err, user.ErrFullNameRequired):

			response.JSON(
				c,
				http.StatusBadRequest,
				false,
				err.Error(),
				nil,
			)

		default:
			response.JSON(
				c,
				http.StatusInternalServerError,
				false,
				"Internal server error",
				nil,
			)
		}

		return
	}

	response.JSON(
		c,
		http.StatusOK,
		true,
		"Registration successful",
		nil,
	)
}
