package routes

import (
	"go-react-router/internal/api/handler"
	"go-react-router/internal/application/identity/command"
	"go-react-router/internal/infrastructure/database"
	"go-react-router/internal/infrastructure/security"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// RegisterUserRoutes registers all user related routes.
func RegisterUserRoutes(
	router *gin.RouterGroup,
	db *pgxpool.Pool,
) {
	userRepository := database.NewUserRepository(db)

	passwordHasher := security.NewBcryptPasswordHasher()

	registerUserCommandHandler := command.NewRegisterUserCommandHandler(
		userRepository,
		passwordHasher,
	)

	userHandler := handler.NewUserHandler(
		registerUserCommandHandler,
	)

	users := router.Group("/users")
	{
		users.POST(
			"/register",
			userHandler.Register,
		)
	}
}
