package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-mail/mail"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/mateus-werneck/portifolio/app/http/middlewares"
	"github.com/mateus-werneck/portifolio/app/storage"
	"github.com/mateus-werneck/portifolio/app/tools"
	"github.com/mateus-werneck/portifolio/app/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func main() {
	godotenv.Load()

	server := gin.Default()
	server.Use(gin.Recovery())

	store := storage.NewSessionStore()
	server.Use(sessions.Sessions("guests", store))
	server.Use(middlewares.LocalizerMiddleware())

	server.LoadHTMLGlob("view/**/*")
	server.Static("/static", "./static")

	server.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)

		lang := "pt-br"
		changeLang := "en"

		if userLang := session.Get("user-lang"); userLang != nil {
			lang = userLang.(string)
		}

		if lang == "en" {
			changeLang = "pt-br"
		}

		slog.Info("UserLang", "language", lang)

		var localizer *i18n.Localizer

		if locale, ok := c.Get("localizer"); ok {
			localizer = locale.(*i18n.Localizer)
		}

		contactButton, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "ContactButton",
		})

		languageOption := "Inglês"
		languageFlag := "/static/images/us.svg"

		if lang == "en" {
			languageOption = "Português"
			languageFlag = "/static/images/br.svg"
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":         "Mateus Werneck",
			"RecentWork":    types.RecentWorks(),
			"ContactButton": contactButton,
			"ChangeLang":    changeLang,
			"LanguageName":  languageOption,
			"LanguageFlag":  languageFlag,
		})
	})

	server.POST("/language/:lang", func(c *gin.Context) {
		lang := c.Param("lang")
		session := sessions.Default(c)

		session.Set("user-lang", lang)
		session.Save()

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
			formData.Errors[field.Field()] = field.Translate(tools.Translator)
		}

		if len(formData.Errors) > 0 {
			c.HTML(http.StatusBadRequest, "contact-form.html", formData)
			return
		}

		session := sessions.Default(c)
		qtdEmails := 0

		if qtd := session.Get(formData.Sender); qtd != nil {
			qtdEmails = qtd.(int)
		}

		if qtdEmails == 0 {
			session.Set(formData.Sender, qtdEmails+1)
			session.Save()
		}

		if qtdEmails >= 10 {
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
