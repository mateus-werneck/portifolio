package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mateus-werneck/portifolio/app/builders"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func home(server *gin.Engine) {
	server.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)

		language := session.Get("user-lang").(string)
		localizer := c.MustGet("localizer").(*i18n.Localizer)

		pageData := builders.NewHomePage().
			SetTitle("Mateus Werneck").
			SetLanguage(language).
			SetLocalizer(localizer).
			Build()

		c.HTML(http.StatusOK, "index.html", pageData)
	})
}
