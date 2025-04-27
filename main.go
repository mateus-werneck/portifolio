package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mateus-werneck/portifolio/app/http/middlewares"
	"github.com/mateus-werneck/portifolio/app/storage"
	"github.com/mateus-werneck/portifolio/app/tools"
	"github.com/mateus-werneck/portifolio/routes"
	sloggin "github.com/samber/slog-gin"
)

func main() {
	godotenv.Load()

	server := gin.Default()
	server.Use(gin.Recovery())

	store := storage.NewSessionStore()

	server.Use(sessions.Sessions("guests", store))
	server.Use(middlewares.LocalizerMiddleware())

	server.Use(sloggin.NewWithConfig(tools.GinLogger, sloggin.Config{
		WithRequestBody:    true,
		WithRequestHeader:  true,
		WithResponseHeader: true,
	}))

	server.LoadHTMLGlob("view/**/*")
	server.Static("/static", "./static")

	routes.AppendRoutes(server)

	if err := server.Run(":2053"); err != nil {
		log.Fatalf("Server initialization failed: %v", err)
	}

	tools.GlobalLogger.Info("Server started", "Port", "2053")
}
