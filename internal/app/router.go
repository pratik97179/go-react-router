package app

import (
	"go-react-router/internal/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewRouter creates and configures the application's HTTP router.
func NewRouter(
	db *pgxpool.Pool,
) *gin.Engine {

	router := gin.New()

	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	api := router.Group("/api/v1")

	routes.RegisterUserRoutes(
		api,
		db,
	)

	return router
}
