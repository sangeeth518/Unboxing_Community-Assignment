package routes

import (
	"Assignment/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func Adminroutes(engine *gin.RouterGroup, adminhandler *handler.AdminaHandler) {
	engine.POST("/login", adminhandler.LoginHandler)
}
