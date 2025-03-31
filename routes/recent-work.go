package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateus-werneck/portifolio/app/builders"
	"github.com/mateus-werneck/portifolio/app/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func recentWork(server *gin.Engine) {
	server.GET("/recent-work/logo/:name", func(c *gin.Context) {
		work := types.FindWork(c.Param("name"))
		work.Opacity = "opacity-0"
		c.HTML(http.StatusOK, "logo.html", work)
	})

	server.GET("/recent-work/summary/:name", func(c *gin.Context) {
		work := types.FindWork(c.Param("name"))
		localizer := c.MustGet("localizer").(*i18n.Localizer)

		c.HTML(http.StatusOK, "logo-summary.html", gin.H{
			"Element":     work.Element,
			"Description": work.Desc(localizer),
			"Website":     work.Website,
			"Buttons": builders.HomePageButtons{
				Visit: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Buttons.Visit"}),
			},
		})
	})
}
