package model

type BuildTarget struct {
	Action   string
	Location string
}

func NewBuildTarget(action, location string) BuildTarget {
	return BuildTarget{
		action,
		location,
	}
}
