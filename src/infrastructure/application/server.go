package application

import (
	"github.com/joeshaw/envdecode"
	"github.com/gin-gonic/gin"
)

type server struct {
	container *container
	router    *gin.Engine
}

func NewApplication() *server {
	config := config{}
	envdecode.Decode(&config)

	server := &server{
		container: newContainer(config),
		router:    gin.Default(),
	}

	server.container.WatcherController().Bind(server.router)
	server.container.ProjectController().Bind(server.router)

	return server
}

func (s *server) Run() {
	s.router.Run()
}
