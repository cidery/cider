package model

import (
	"time"
	"github.com/satori/go.uuid"
)

type Listener struct {
	id            uuid.UUID
	class         string
	scope         string
	registeredAt  time.Time
	lastUpdatedAt time.Time
}

func NewListener(id uuid.UUID, class, scope string, registeredAt time.Time) *Listener {
	return &Listener{id, class, scope, registeredAt, registeredAt}
}

func (w *Listener) TrackUpdate(when time.Time) {
	w.lastUpdatedAt = when
}

func (w *Listener) Class() string {
	return w.class
}

func (w *Listener) Scope() string {
	return w.scope
}

func (w *Listener) Id() uuid.UUID {
	return w.id
}

func (w *Listener) RegisteredAt() time.Time {
	return w.registeredAt
}

func (w *Listener) LastUpdatedAt() time.Time {
	return w.lastUpdatedAt
}
