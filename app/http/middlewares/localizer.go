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

		lang := "pt-br"

		if sessionLang := session.Get("user-lang"); sessionLang != nil {
			lang = sessionLang.(string)
		}

		if lang == "en" {
			localizer = i18n.NewLocalizer(bundle, language.English.String())
		}

		c.Set("localizer", localizer)

		c.Next()
	}
}
