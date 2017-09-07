package errors

import (
	"github.com/satori/go.uuid"
	"fmt"
	"errors"
)

type WatcherTargetMismatchError struct {
	error
	id             uuid.UUID
	targetExist    string
	targetRegister string
}

func NewWatcherTargetMismatchError(id uuid.UUID, targetExist, targetRegister string) *WatcherTargetMismatchError {
	return &WatcherTargetMismatchError{
		errors.New(fmt.Sprintf(
			"Watcher %s has target %s, trying to register %s",
			id.String(),
			targetExist,
			targetRegister,
		)),
		id,
		targetExist,
		targetRegister,
	}
}
