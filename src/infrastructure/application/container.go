package application

import (
	"database/sql"
	"github.com/cidery/cider/src/infrastructure/database"
	"github.com/cidery/cider/src/infrastructure/controllers"
	"github.com/cidery/cider/src/domain/service"
	"log"
	"os"
)

const (
	l_project_controller = "project_controller"
	l_watcher_controller = "watcher_controller"

	l_project_registry = "project_registry"
	l_watcher_registry = "watcher_registry"

	l_database      = "database"
	l_logger        = "logger"
	l_time_provider = "time_provider"
)

type container struct {
	config   config
	registry map[string]interface{}
}

type constructor func() (interface{}, error)

func newContainer(config config) *container {
	return &container{
		config:   config,
		registry: make(map[string]interface{}),
	}
}

func (c *container) Database() *sql.DB {
	s := c.createOrGet(
		l_database,
		func() (interface{}, error) {
			return database.NewDatabase(c.config.Database)
		},
	)

	return s.(*sql.DB)
}

func (c *container) WatcherController() *controllers.WatcherController {
	s := c.createOrGet(
		l_watcher_controller,
		func() (interface{}, error) {
			return controllers.NewWatcherController(c.WatcherRegistry()), nil
		},
	)

	return s.(*controllers.WatcherController)
}

func (c *container) ProjectController() *controllers.ProjectController {
	s := c.createOrGet(
		l_project_controller,
		func() (interface{}, error) {
			return controllers.NewProjectController(c.ProjectRegistry()), nil
		},
	)

	return s.(*controllers.ProjectController)
}

func (c *container) WatcherRegistry() *service.WatcherRegistry {
	s := c.createOrGet(
		l_watcher_registry,
		func() (interface{}, error) {
			return service.NewWatcherRegistry(c.TimeProvider(), c.Logger()), nil
		},
	)

	return s.(*service.WatcherRegistry)
}

func (c *container) ProjectRegistry() *service.ProjectRegistry {
	s := c.createOrGet(
		l_project_registry,
		func() (interface{}, error) {
			return service.NewProjectRegistry(c.TimeProvider(), c.WatcherRegistry(), c.Logger()), nil
		},
	)

	return s.(*service.ProjectRegistry)
}

func (c *container) Logger() *log.Logger {
	s := c.createOrGet(
		l_logger,
		func() (interface{}, error) {
			return log.New(os.Stderr, "", log.LstdFlags), nil
		},
	)

	return s.(*log.Logger)
}

func (c *container) TimeProvider() *service.RealTimeProvider {
	s := c.createOrGet(
		l_time_provider,
		func() (interface{}, error) {
			return &service.RealTimeProvider{}, nil
		},
	)

	return s.(*service.RealTimeProvider)
}

func (c *container) createOrGet(locator string, init constructor) interface{} {
	_, found := c.registry[locator]
	if false == found {
		s, err := init()
		if nil != err {
			panic(err)
		}

		c.registry[locator] = s
	}

	return c.registry[locator]
}
