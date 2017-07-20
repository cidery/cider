package response

import (
	"github.com/cidery/cider/src/domain/model"
	"time"
)

func NewWatcherResponseFromWatcher(w *model.Watcher) *WatcherResponse {
	return &WatcherResponse{
		w.Id().String(),
		w.Class(),
		w.Scope(),
		w.RegisteredAt(),
		w.LastUpdatedAt(),
	}
}

type WatcherResponse struct {
	Id            string
	Class         string
	Scope         string
	RegisteredAt  time.Time
	LastUpdatedAt time.Time
}
