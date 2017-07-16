package model

type Project struct {
	Name    string
	Watcher Watcher
	Targets []BuildTarget
}
