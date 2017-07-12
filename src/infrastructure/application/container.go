package application

import (
	"database/sql"
	"github.com/cidery/cider/src/infrastructure/database"
	"github.com/cidery/cider/src/infrastructure/controllers"
)

const (
	l_database           = "database"
	l_watcher_controller = "watcher_controller"
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
	service := c.createOrGet(
		l_database,
		func() (interface{}, error) {
			return database.NewDatabase(c.config.Database)
		},
	)

	return service.(*sql.DB)
}

func (c *container) WatcherController() *controllers.WatcherController {
	service := c.createOrGet(
		l_watcher_controller,
		func() (interface{}, error) {
			return controllers.NewWatcherController(), nil
		},
	)

	return service.(*controllers.WatcherController)
}

func (c *container) createOrGet(locator string, init constructor) interface{} {
	_, found := c.registry[locator]
	if false == found {
		service, err := init()
		if nil != err {
			panic(err)
		}

		c.registry[locator] = service
	}

	return c.registry[locator]
}
