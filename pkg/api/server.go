package http

import (
	"Assignment/pkg/api/handler"
	"Assignment/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHttp struct {
	engine *gin.Engine
}

func NewServerHttp(adminhandler *handler.AdminaHandler) *ServerHttp {
	engine := gin.New()
	engine.Use(gin.Logger())
	routes.Adminroutes(engine.Group("/admin"), adminhandler)

	return &ServerHttp{
		engine: engine,
	}
}
func (sh *ServerHttp) Start() {
	err := sh.engine.Run(":5000")
	if err != nil {
		log.Fatal("gin engine could not start")
	}
}
