package http

import (
	"go-boiler/config"
	"go-boiler/docs"
	"net/http"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go-boiler/utils/log"

	"github.com/gin-gonic/gin"
)

const (
	DEVELOPMENT = "development"
)

type ServerHTTP struct {
	Config *config.Config
	gin    *gin.Engine
}

func InitializeServer(config *config.Config) *ServerHTTP {
	return &ServerHTTP{
		Config: config,
	}
}

// reference : https://gin-gonic.com/docs/
// reference : https://pkg.go.dev/github.com/gin-gonic/gin
// initialize server
func (server *ServerHTTP) SetupAndServe() {
	server.gin = gin.Default()
	server.setupSwagger()
	server.setupRoutes()
	err := server.gin.Run(":" + server.Config.Server.Port)
	if err != nil {
		log.ErrorWithStack(err)
	}
}

// reference : https://github.com/swaggo/gin-swagger
// setup swagger
func (server *ServerHTTP) setupSwagger() {
	if server.Config.Server.Env == DEVELOPMENT {
		docs.SwaggerInfo.BasePath = "/api/v1"
		server.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}

func (server *ServerHTTP) setupRoutes() {
	v1 := server.gin.Group("/api/v1")
	v1.GET("/health", CheckHealth)
}

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping check health
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Status OK
// @Router /health [get]
func CheckHealth(g *gin.Context) {
	g.JSON(http.StatusOK, "Status OK")
}
