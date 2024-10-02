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
	"github.com/gin-gonic/gin/binding"
	"github.com/go-mail/mail"
	"github.com/go-playground/locales/en"
	pt "github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/pt_BR"
	"github.com/joho/godotenv"
	"github.com/mateus-werneck/portifolio/app/storage"
	"github.com/mateus-werneck/portifolio/app/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/redis/go-redis/v9"
	"golang.org/x/text/language"
)

func main() {
	godotenv.Load()
	rdb := storage.NewRedis()

	trans := initTranslator()
	bundle := initLaguangueBundle()
	localizer := i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String())

	server := gin.Default()
	server.Use(gin.Recovery())

	server.LoadHTMLGlob("view/**/*")
	server.Static("/static", "./static")

	server.GET("/", func(c *gin.Context) {
		contactButton, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "ContactButton",
		})

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":         "Mateus Werneck",
			"RecentWork":    types.RecentWorks(),
			"ContactButton": contactButton,
		})
	})

	server.POST("/language/:lang", func(c *gin.Context) {
		lang := c.Param("lang")

		if lang == "en" {
			localizer = i18n.NewLocalizer(bundle, language.English.String())
		}

		if lang == "ptBr" {
			localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String())
		}

		c.Header("HX-Location", "/")
		c.Status(http.StatusOK)
	})

	server.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", gin.H{})
	})

	server.GET("/recent-work/logo/:name", func(c *gin.Context) {
		work := types.FindWork(c.Param("name"))
		work.Opacity = "opacity-0"
		c.HTML(http.StatusOK, "logo.html", work)
	})

	server.GET("/recent-work/summary/:name", func(c *gin.Context) {
		work := types.FindWork(c.Param("name"))
		c.HTML(http.StatusOK, "logo-summary.html", work)
	})

	server.POST("/contact", func(c *gin.Context) {
		var validationErrors validator.ValidationErrors
		var formData types.ContactEmail

		err := c.ShouldBind(&formData)

		if err != nil {
			validationErrors = err.(validator.ValidationErrors)
		}

		formData.Errors = map[string]string{}

		for _, field := range validationErrors {
			formData.Errors[field.Field()] = field.Translate(trans)
		}

		if len(formData.Errors) > 0 {
			c.HTML(http.StatusBadRequest, "contact-form.html", formData)
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

func initTranslator() ut.Translator {
	pt := pt.New()
	en := en.New()
	uni := ut.New(pt, pt, en)
	trans, _ := uni.GetTranslator("pt_BR")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		pt_BR.RegisterDefaultTranslations(v, trans)
	}

	return trans
}

func initLaguangueBundle() *i18n.Bundle {
	bundle := i18n.NewBundle(language.BrazilianPortuguese)
	bundle.MustLoadMessageFile("translations/en.json")
	bundle.MustLoadMessageFile("translations/pt-BR.json")

	return bundle
}
