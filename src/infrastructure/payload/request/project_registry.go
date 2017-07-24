package request

import (
	"github.com/satori/go.uuid"
)

type ProjectRegisterRequest struct {
	Name    string
	Locator string
	Watcher uuid.UUID
	Target  []*Target
}

type Target struct {
	Action   string
	Location string
}
