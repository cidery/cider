package request

import "github.com/satori/go.uuid"

type WatcherRegisterRequest struct {
	Id    uuid.UUID
	Class string
	Scope string
}

