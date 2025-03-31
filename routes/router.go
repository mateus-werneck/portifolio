package routes

import "github.com/gin-gonic/gin"

func AppendRoutes(server *gin.Engine) {
	home(server)
	contact(server)
	language(server)
	recentWork(server)
}
