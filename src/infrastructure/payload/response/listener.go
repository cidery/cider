package response

import (
	"github.com/cidery/cider/src/domain/model"
	"time"
)

func NewListenerResponseFromWatcher(listener *model.Listener) *LsitenerResponse {
	return &LsitenerResponse{
		listener.Id().String(),
		listener.Class(),
		listener.Scope(),
		listener.RegisteredAt(),
		listener.LastUpdatedAt(),
	}
}

type LsitenerResponse struct {
	Id            string
	Class         string
	Scope         string
	RegisteredAt  time.Time
	LastUpdatedAt time.Time
}
