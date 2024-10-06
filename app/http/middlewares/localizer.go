package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mateus-werneck/portifolio/app/tools"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func LocalizerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		bundle := tools.Bundle
		localizer := i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String())

		lang := c.GetHeader("Accept-Language")

		if lang == "" {
			lang = "pt-BR"
		}

		if sessionLang := session.Get("user-lang"); sessionLang != nil {
			lang = sessionLang.(string)
		}

		if lang == "en-US" {
			localizer = i18n.NewLocalizer(bundle, language.English.String())
		}

		session.Set("user-lang", lang)
		session.Save()

		c.Set("localizer", localizer)

		c.Next()
	}
}
