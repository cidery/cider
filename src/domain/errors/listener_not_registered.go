package errors

import (
	"github.com/satori/go.uuid"
	"errors"
	"fmt"
)

type ListenerNotRegisteredError struct {
	error
	id uuid.UUID
}

func NewListenerNotRegisteredError(id uuid.UUID) *ListenerNotRegisteredError {
	return &ListenerNotRegisteredError{
		errors.New(fmt.Sprintf("Listener with id %s not registered", id.String())),
		id,
	}
}
