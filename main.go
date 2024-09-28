package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecentWork struct {
	Element     string
	Image       string
	Description string
	Opacity     string
}

var works = map[string]RecentWork{
	"celcoin": {
		Element:     "celcoin",
		Image:       "celcoin.svg",
		Description: "Infratech financeira para potencializar negócios",
		Opacity:     "opacity-100",
	},
	"symplicity": {
		Element:     "symplicity",
		Image:       "symplicity.webp",
		Description: "Streamline system-wide opportunities and increase student engagement",
		Opacity:     "opacity-100",
	},
}

func main() {
	server := gin.Default()
	server.Use(gin.Recovery())

	server.LoadHTMLGlob("view/**/*")
	server.Static("/static", "./static")

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":      "Mateus Werneck",
			"RecentWork": works,
		})
	})

	server.GET("/recent-work/logo/:name", func(c *gin.Context) {
		work := works[c.Param("name")]
		c.HTML(http.StatusOK, "logo.html", gin.H{"Opacity": "opacity-0", "Element": c.Param("name"), "Image": work.Image})
	})

	server.GET("/recent-work/summary/:name", func(c *gin.Context) {
		work := works[c.Param("name")]

		c.HTML(http.StatusOK, "logo-summary.html", gin.H{
			"Element":     c.Param("name"),
			"Description": work.Description,
		})
	})

	if err := server.Run(":9010"); err != nil {
		log.Fatalf("Server initialization failed: %v", err)
	}
}
