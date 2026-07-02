package handler

import (
	"errors"
	"net/http"

	"go-react-router/internal/api/response"
	"go-react-router/internal/application/identity/command"
	"go-react-router/internal/domain/identity"
	"go-react-router/internal/domain/identity/aggregate"

	"github.com/gin-gonic/gin"
)

// UserHandler handles identity-related HTTP requests.
type UserHandler struct {
	registerUser *command.RegisterUserCommandHandler
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(
	registerUser *command.RegisterUserCommandHandler,
) *UserHandler {
	return &UserHandler{
		registerUser: registerUser,
	}
}

// Register handles identity registration.
func (h *UserHandler) Register(
	c *gin.Context,
) {
	var registration aggregate.Registration

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

		case errors.Is(err, identity.ErrEmailAlreadyExists):
			response.JSON(
				c,
				http.StatusConflict,
				false,
				err.Error(),
				nil,
			)

		case errors.Is(err, identity.ErrEmailRequired),
			errors.Is(err, identity.ErrInvalidEmail),
			errors.Is(err, identity.ErrPasswordRequired),
			errors.Is(err, identity.ErrFullNameRequired):

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
