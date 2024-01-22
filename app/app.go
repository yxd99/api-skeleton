package app

import (
	"api-skeleton/app/db"
	"api-skeleton/app/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error to load the file .env %v", err)
	}

	db.Connet()

	router := gin.Default()
	application := &App{
		Router: router,
	}
	application.configureRoutes()
	return application
}

func (app *App) configureRoutes() {
	pingHandler := handlers.NewPingHandler()
	greetHandler := handlers.NewGreetHandler()
	routes := app.Router.Group("api")
	routes.GET("/ping", pingHandler.Ping)
	routes.GET("/greet", greetHandler.Greet)
	app.Router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Route not found"})
	})
}

func (app *App) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	address := ":" + port
	return app.Router.Run(address)
}
