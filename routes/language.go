package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mateus-werneck/portifolio/app/tools"
)

func language(server *gin.Engine) {
	server.POST("/language/:lang", func(c *gin.Context) {
		lang := c.Param("lang")
		session := sessions.Default(c)

		session.Set("user-lang", lang)
		session.Save()

		if lang == "en-US" {
			tools.SetEnTransalator()
		}

		if lang == "pt-BR" {
			tools.SetPtBrTransaltor()
		}

		c.Header("HX-Location", "/")
		c.Status(http.StatusOK)
	})
}
