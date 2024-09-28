package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(gin.Recovery())

	server.LoadHTMLGlob("view/**/*")
	server.Static("/static", "./static")

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":             "Mateus Werneck",
			"CelcoinOpacity":    "opacity-100",
			"SymplicityOpacity": "opacity-100",
		})
	})

	server.GET("/recent-work/celcoin/logo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "celcoin-logo.html", gin.H{"CelcoinOpacity": "opacity-0"})
	})

	server.GET("/recent-work/celcoin/animation", func(c *gin.Context) {
		c.HTML(http.StatusOK, "celcoin-animation.html", nil)
	})

	server.GET("/recent-work/symplicity/logo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "symplicity-logo.html", gin.H{"SymplicityOpacity": "opacity-0"})
	})

	server.GET("/recent-work/symplicity/animation", func(c *gin.Context) {
		c.HTML(http.StatusOK, "symplicity-animation.html", nil)
	})

	if err := server.Run(":9010"); err != nil {
		log.Fatalf("Server initialization failed: %v", err)
	}
}
