package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(db *pgxpool.Pool) (*gin.Engine, error) {
	router := gin.New()

	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	return router, nil
}
