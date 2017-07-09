package model

import "time"

type Agent struct {
	Type         string
	Target       string
	RegisteredAt time.Time
	LastPingAt   time.Time
}
