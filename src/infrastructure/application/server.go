package application

import "github.com/joeshaw/envdecode"

type server struct {
	container *container
}

func NewApplication() *server {
	config := config{}
	envdecode.Decode(&config)

	return &server{
		container: newContainer(config),
	}
}

func (s *server) Run() {

}
