package model

type Project struct {
	name     string
	locator  string
	listener *Listener
	targets  []BuildTarget
}

func NewProject(name, locator string, listener *Listener, targets []BuildTarget) *Project {
	return &Project{
		name,
		locator,
		listener,
		targets,
	}
}

func (p *Project) UpdateName(name string) {
	p.name = name
}

func (p *Project) UpdateBuildTargets(buildTargets []BuildTarget) {
	p.targets = buildTargets
}
