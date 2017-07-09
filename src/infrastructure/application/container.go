package application

import (
	"database/sql"
	"github.com/cidery/cider/src/infrastructure/database"
)

const (
	l_database = "database"
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

func (c *container) createOrGet(locator string, init constructor) interface{} {
	_, found := c.registry[locator]
	if false == found {
		service, err := init()
		if nil != err {
			panic(err)
		}

		c.registry[l_database] = service
	}

	return c.registry[l_database]
}
