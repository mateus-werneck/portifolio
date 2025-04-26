package middlewares

import (
	"strings"

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
		langParts := strings.Split(lang, ",")

		langName := langParts[0]

		if langName == "" {
			langName = "pt-BR"
		}

		if sessionLang := session.Get("user-lang"); sessionLang != nil {
			langName = sessionLang.(string)
		}

		if langName == "en-US" {
			localizer = i18n.NewLocalizer(bundle, language.English.String())
			tools.SetEnTransalator()
		}

		if langName == "pt-BR" {
			tools.SetPtBrTransaltor()
		}

		session.Set("user-lang", langName)
		session.Save()

		c.Set("localizer", localizer)
	}
}
