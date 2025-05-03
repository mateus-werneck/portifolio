package routes

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mateus-werneck/portifolio/app/builders"
	"github.com/mateus-werneck/portifolio/app/tools"
	"github.com/mateus-werneck/portifolio/app/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gopkg.in/mail.v2"
)

func contact(server *gin.Engine) {
	server.GET("/contact", func(c *gin.Context) {
		localizer := c.MustGet("localizer").(*i18n.Localizer)
		c.HTML(http.StatusOK, "contact.html", gin.H{
			"FormTitle":       localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Contact.Title"}),
			"FormDescription": localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Contact.Desc"}),
        	"Buttons": builders.HomePageButtons{
				Submit: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Buttons.Submit"}),
			},
			"ContactFields": map[string]string{
				"Name":    localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "ContactFields.Name"}),
				"Email":   localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "ContactFields.Email"}),
				"Message": localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "ContactFields.Message"}),
			},
		})
	})

	server.POST("/contact", func(c *gin.Context) {
		var validationErrors validator.ValidationErrors
		var formData types.ContactEmail

		localizer := c.MustGet("localizer").(*i18n.Localizer)

		err := c.ShouldBind(&formData)
		if err != nil {
			validationErrors = err.(validator.ValidationErrors)
		}

		formData.Errors = map[string]string{}

		for _, field := range validationErrors {
			formData.Errors[field.Field()] = field.Translate(tools.Translator)
		}

		contactFields := map[string]string{
			"Name":    localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "ContactFields.Name"}),
			"Email":   localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "ContactFields.Email"}),
			"Message": localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "ContactFields.Message"}),
		}

		response := gin.H{
			"ContactFields": contactFields,
			"Name":          formData.Name,
			"Sender":        formData.Email,
			"Body":          formData.Message,
		}

		if len(validationErrors) > 0 {
			tools.GlobalLogger.Error("ContactForm invalid input", "data", formData)
			response["FormErrors"] = formData.Errors

			c.HTML(http.StatusBadRequest, "contact-form.html", response)

			return
		}

		session := sessions.Default(c)
		qtdEmails := 0

		if qtd := session.Get(formData.Email); qtd != nil {
			qtdEmails = qtd.(int)
		}

		if qtdEmails >= 10 {
			response["Error"] = localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Errors.EmailLimit"})
			c.HTML(http.StatusBadRequest, "contact.html", response)
			return
		}

		email := mail.NewMessage()
		email.SetHeader("From", formData.Email)
		email.SetHeader("To", "werneck.mateus@gmail.com", "werneck.mateus@protonmail.com")
		email.SetHeader("Subject", "Me interessei no seu perfil - Mateus Werneck")
		email.SetBody("text/plain", formData.Message)

		port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
		d := mail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"))

		err = d.DialAndSend(email)
		if err != nil {
			tools.GlobalLogger.Error("SMTP sendEmail failed", "error", err)

			response["Error"] = localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Errors.SendEmail"})

			c.HTML(http.StatusBadRequest, "contact-form.html", response)

			return
		}

		session.Set(formData.Email, qtdEmails+1)
		session.Save()

		c.HTML(http.StatusOK, "contact-form.html", gin.H{})
	})
}
