package model

import (
	"time"
	"github.com/satori/go.uuid"
)

type Watcher struct {
	id           uuid.UUID
	class        string
	scope        string
	registeredAt time.Time
	lastUpdated  time.Time
}

func NewWatcher(id uuid.UUID, class, scope string, registeredAt time.Time) *Watcher {
	return &Watcher{id, class, scope, registeredAt, registeredAt}
}

func (w *Watcher) TrackUpdate(when time.Time) {
	w.lastUpdated = when
}

func (w *Watcher) Class() string {
	return w.class
}

func (w *Watcher) Scope() string {
	return w.scope
}
