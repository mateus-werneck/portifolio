package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-mail/mail"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type RecentWork struct {
	Element     string
	Image       string
	Description string
	Website     string
	Opacity     string
}

var works = map[string]RecentWork{
	"celcoin": {
		Element:     "celcoin",
		Image:       "celcoin.svg",
		Description: "Infratech financeira para potencializar negócios",
		Website:     "https://www.celcoin.com.br",
		Opacity:     "opacity-100",
	},
	"symplicity": {
		Element:     "symplicity",
		Image:       "symplicity.webp",
		Description: "Streamline system-wide opportunities and increase student engagement",
		Website:     "https://www.symplicity.com",
		Opacity:     "opacity-100",
	},
}

type ContactEmail struct {
	Name   string `form:"name" binding:"required,min=3,alpha"`
	Sender string `form:"email" binding:"required,email"`
	Body   string `form:"message" binding:"required,min=10,max=400"`
}

func main() {
	godotenv.Load()

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

	server.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", gin.H{})
	})

	server.POST("/contact", func(c *gin.Context) {
		var formData ContactEmail

		err := c.ShouldBind(&formData)
		validationErrors := err.(validator.ValidationErrors)

		if len(validationErrors) > 0 {
			c.HTML(http.StatusBadRequest, "contact-form.html", gin.H{
				"errors": validationErrors,
			})
			return
		}

		email := mail.NewMessage()
		email.SetHeader("From", formData.Sender)
		email.SetHeader("To", "werneck.mateus@gmail.com", "werneck.mateus@protonmail.com")
		email.SetHeader("Subject", "Me interessei no seu perfil - Mateus Werneck")
		email.SetBody("text/plain", formData.Body)

		port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
		d := mail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"))

		err = d.DialAndSend(email)
		if err != nil {
			slog.Error("SMTP sendEmail failed", "error", err)
			c.HTML(http.StatusBadRequest, "contact-form.html", gin.H{})
			return
		}

		c.HTML(http.StatusOK, "contact-form.html", gin.H{})
	})

	server.GET("/recent-work/logo/:name", func(c *gin.Context) {
		work := works[c.Param("name")]
		work.Opacity = "opacity-0"
		c.HTML(http.StatusOK, "logo.html", work)
	})

	server.GET("/recent-work/summary/:name", func(c *gin.Context) {
		work := works[c.Param("name")]
		c.HTML(http.StatusOK, "logo-summary.html", work)
	})

	if err := server.Run(":9010"); err != nil {
		log.Fatalf("Server initialization failed: %v", err)
	}
}
