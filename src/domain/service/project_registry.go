package service

import (
	"log"
	"sync"
	"github.com/cidery/cider/src/domain/model"
	"github.com/satori/go.uuid"
	"fmt"
)

type ProjectRegistry struct {
	timeProvider    TimeProvider
	watcherRegistry *WatcherRegistry
	logger          *log.Logger

	projects map[string]*model.Project
	mu       *sync.Mutex
}

func NewProjectRegistry(timeProvider TimeProvider, watcherRegistry *WatcherRegistry, logger *log.Logger) *ProjectRegistry {
	return &ProjectRegistry{
		timeProvider:    timeProvider,
		watcherRegistry: watcherRegistry,
		logger:          logger,

		projects: make(map[string]*model.Project),
		mu:       &sync.Mutex{},
	}
}

func (p *ProjectRegistry) RegisterProject(name, locator string, watcherId uuid.UUID, buildTargets []model.BuildTarget) error {
	watcher, err := p.watcherRegistry.GetWatcher(watcherId)

	if nil != err {
		return err
	}

	p.mu.Lock()

	key := fmt.Sprintf("%s-%s-%s", watcher.Class(), watcher.Scope(), locator)
	project, found := p.projects[key]

	if false == found {
		p.projects[key] = model.NewProject(name, locator, watcher, buildTargets)
	} else {
		project.UpdateName(name)
		project.UpdateBuildTargets(buildTargets)
	}

	p.mu.Unlock()

	return nil
}
