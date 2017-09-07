package model

import "time"

type Worker struct {
	Type         string
	Target       string
	RegisteredAt time.Time
	LastPingAt   time.Time
}
