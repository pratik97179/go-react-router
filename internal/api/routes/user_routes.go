package routes

import (
	"go-react-router/internal/api/handler"
	"go-react-router/internal/domain/user/usecase/commands"
	"go-react-router/internal/platform/database"
	"go-react-router/internal/platform/security"

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

	registerUserUseCase := commands.NewRegisterUserUseCase(
		userRepository,
		passwordHasher,
	)

	userHandler := handler.NewUserHandler(
		registerUserUseCase,
	)

	users := router.Group("/users")
	{
		users.POST(
			"/register",
			userHandler.Register,
		)
	}
}
