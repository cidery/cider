package service

import (
	"github.com/cidery/cider/src/domain/model"
	"log"
	"github.com/satori/go.uuid"
	"github.com/cidery/cider/src/domain/errors"
	"sync"
)

type WatcherRegistry struct {
	timeProvider TimeProvider
	logger       *log.Logger

	watchers  map[string]*model.Listener
	targetMap map[string]string
	mu        *sync.Mutex
}

func NewWatcherRegistry(timeProvider TimeProvider, logger *log.Logger) *WatcherRegistry {
	return &WatcherRegistry{
		timeProvider: timeProvider,
		logger:       logger,
		watchers:     make(map[string]*model.Listener),
		targetMap:    make(map[string]string),
		mu:           &sync.Mutex{},
	}
}

func (wr *WatcherRegistry) GetWatcher(id uuid.UUID) (*model.Listener, error) {
	watcher, found := wr.watchers[id.String()]

	if false == found {
		return nil, errors.NewListenerNotRegisteredError(id)
	}

	return watcher, nil
}

func (wr *WatcherRegistry) RegisterWatcher(id uuid.UUID, class, scope string) error {
	target := wr.getTarget(class, scope)

	wr.mu.Lock()
	defer wr.mu.Unlock()

	if w, found := wr.watchers[id.String()]; true == found {
		if target == wr.getWatcherTarget(w) {
			w.TrackUpdate(wr.timeProvider.Now())
		} else {
			err := errors.NewWatcherTargetMismatchError(id, wr.getWatcherTarget(w), target)
			wr.logger.Println(err.Error())
			return err
		}
	} else {
		if wid, found := wr.targetMap[target]; true == found {
			wr.logger.Printf("Listener %s found for target %s, removing in favour of %s", wid, target, id.String())
			delete(wr.targetMap, wid)
		}

		wr.watchers[id.String()] = model.NewListener(id, class, scope, wr.timeProvider.Now())
		wr.targetMap[target] = id.String()
	}

	return nil
}

func (wr *WatcherRegistry) Watchers() []*model.Listener {
	result := make([]*model.Listener, 0, len(wr.watchers))

	for _, value := range wr.watchers {
		result = append(result, value)
	}

	return result
}

func (wr *WatcherRegistry) getTarget(class, scope string) string {
	return class + "." + scope
}

func (wr *WatcherRegistry) getWatcherTarget(w *model.Listener) string {
	return wr.getTarget(w.Class(), w.Scope())
}
