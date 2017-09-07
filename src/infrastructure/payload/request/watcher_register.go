package request

import "github.com/satori/go.uuid"

type ListenerRegisterRequest struct {
	Id    uuid.UUID
	Class string
	Scope string
}

