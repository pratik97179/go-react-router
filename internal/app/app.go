package app

import (
	"go-react-router/internal/config"
	"go-react-router/internal/platform/database"

	"github.com/gin-gonic/gin"
	"github.com/supabase-community/supabase-go"
)

type Application struct {
	router *gin.Engine
	config  *config.Config
	db     *supabase.Client
}

func New() (*Application, error) {
	app := &Application{}

	if err := app.initializeConfig(); err != nil {
		return nil, err
	}

	if err := app.initializeDatabase(); err != nil {
		return nil, err
	}

	if err := app.initializeRouter(); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *Application) Run() error {
	return a.router.Run(":" + a.config.Server.Port)
}

func (a *Application) initializeConfig() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	a.config = cfg
	return nil
}

func (a *Application) initializeDatabase() error {
	db, err := database.New(a.config.Database.URL, a.config.Database.Key)
	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *Application) initializeRouter() error {
	router, err := NewRouter(a.db)
	if err != nil {
		return err
	}

	a.router = router
	return nil
}
