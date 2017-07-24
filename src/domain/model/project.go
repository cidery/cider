package model

type Project struct {
	name    string
	locator string
	watcher *Watcher
	targets []BuildTarget
}

func NewProject(name, locator string, watcher *Watcher, targets []BuildTarget) *Project {
	return &Project{
		name,
		locator,
		watcher,
		targets,
	}
}

func (p *Project) UpdateName(name string) {
	p.name = name
}

func (p *Project) UpdateBuildTargets(buildTargets []BuildTarget) {
	p.targets = buildTargets
}
