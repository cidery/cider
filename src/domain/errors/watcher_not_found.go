package errors

import (
	"github.com/satori/go.uuid"
	"errors"
	"fmt"
)

type WatcherNotRegisteredError struct {
	error
	id uuid.UUID
}

func NewWatcherNotRegisteredError(id uuid.UUID) *WatcherNotRegisteredError {
	return &WatcherNotRegisteredError{
		errors.New(fmt.Sprintf("Watcher with id %s not registered", id.String())),
		id,
	}
}
