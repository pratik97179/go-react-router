package app

import (
	"github.com/gin-gonic/gin"
	"github.com/supabase-community/supabase-go"
)

func NewRouter(db *supabase.Client) (*gin.Engine, error) {
	router := gin.New()

	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	return router, nil
}
