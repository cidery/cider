package model

import (
	"time"
	"github.com/satori/go.uuid"
)

type Watcher struct {
	id            uuid.UUID
	class         string
	scope         string
	registeredAt  time.Time
	lastUpdatedAt time.Time
}

func NewWatcher(id uuid.UUID, class, scope string, registeredAt time.Time) *Watcher {
	return &Watcher{id, class, scope, registeredAt, registeredAt}
}

func (w *Watcher) TrackUpdate(when time.Time) {
	w.lastUpdatedAt = when
}

func (w *Watcher) Class() string {
	return w.class
}

func (w *Watcher) Scope() string {
	return w.scope
}

func (w *Watcher) Id() uuid.UUID {
	return w.id
}

func (w *Watcher) RegisteredAt() time.Time {
	return w.registeredAt
}

func (w *Watcher) LastUpdatedAt() time.Time {
	return w.lastUpdatedAt
}
