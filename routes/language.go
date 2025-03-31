package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func language(server *gin.Engine) {
	server.POST("/language/:lang", func(c *gin.Context) {
		lang := c.Param("lang")
		session := sessions.Default(c)

		session.Set("user-lang", lang)
		session.Save()

		c.Header("HX-Location", "/")
		c.Status(http.StatusOK)
	})
}
