package response

import "github.com/cidery/cider/src/domain/model"

func NewWatcherListResponse(listeners []*model.Listener) *ListenerListResponse {
	res := &ListenerListResponse{
		Listener: make([]*LsitenerResponse, 0, len(listeners)),
	}
	for _, l := range listeners {
		res.Listener = append(res.Listener, NewListenerResponseFromWatcher(l))
	}

	return res
}

type ListenerListResponse struct {
	Listener []*LsitenerResponse
}
