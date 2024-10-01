package main

import (
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-mail/mail"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
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
	rdb := NewRedis()

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

	server.GET("/recent-work/logo/:name", func(c *gin.Context) {
		work := works[c.Param("name")]
		work.Opacity = "opacity-0"
		c.HTML(http.StatusOK, "logo.html", work)
	})

	server.GET("/recent-work/summary/:name", func(c *gin.Context) {
		work := works[c.Param("name")]
		c.HTML(http.StatusOK, "logo-summary.html", work)
	})

	server.POST("/contact", func(c *gin.Context) {
		var validationErrors validator.ValidationErrors
		var formData ContactEmail

		err := c.ShouldBind(&formData)

		if err != nil {
			validationErrors = err.(validator.ValidationErrors)
		}

		if len(validationErrors) > 0 {
			c.HTML(http.StatusBadRequest, "contact-form.html", gin.H{
				"errors": validationErrors,
			})
			return
		}

		sentEmails, err := rdb.Get(c, formData.Sender).Int()

		if errors.Is(err, redis.Nil) {
			rdb.Set(c, formData.Sender, sentEmails+1, time.Duration(time.Second*86400))
		}

		if err != nil && !errors.Is(err, redis.Nil) {
			slog.Error("Failed to find sender on redis", "sender", formData.Sender, "error", err.Error())
		}

		if sentEmails >= 10 {
			c.HTML(http.StatusBadRequest, "contact-.html", gin.H{
				"errors": []string{"Limite de emails atingido."},
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

	if err := server.Run(":9010"); err != nil {
		log.Fatalf("Server initialization failed: %v", err)
	}
}
