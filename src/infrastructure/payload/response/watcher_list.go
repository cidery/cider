package response

import "github.com/cidery/cider/src/domain/model"

func NewWatcherListResponse(watchers []*model.Watcher) *WatcherListResponse {
	res := &WatcherListResponse{
		Watchers: make([]*WatcherResponse, 0, len(watchers)),
	}
	for _, w := range watchers {
		res.Watchers = append(res.Watchers, NewWatcherResponseFromWatcher(w))
	}

	return res
}

type WatcherListResponse struct {
	Watchers []*WatcherResponse
}
