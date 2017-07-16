package service

import "time"

type TimeProvider interface {
	Now() time.Time
}

type RealTimeProvider struct {
}

func (rtp *RealTimeProvider) Now() time.Time {
	return time.Now()
}
