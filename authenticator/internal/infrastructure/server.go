package infrastructure

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	Router *gin.Engine
}

func NewServer(port string, mode string) Server {
	switch mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	server := Server{
		port:   port,
		Router: gin.Default(),
	}

	server.Router.Use(cors.Default())

	return server
}

func (s *Server) Start() {
	s.Router.Run(":" + s.port)
}
