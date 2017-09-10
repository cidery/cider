package main

import "github.com/cidery/cider/src/infrastructure/application"

var (
	version   = "dev build"
	commit    = "unknown"
	buildTIme = "unknown"
)

func main() {
	application.NewApplication().Run()
}
