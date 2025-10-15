package main

import (
	"log"

	"github.com/metux/mpbt/core"
)

func main() {
	prj := core.Project{}

	err := prj.LoadComponents("../cf/xlibre/components")
	if err != nil {
		log.Fatalf("error opening components directory: %s\n", err)
	}

	prj.LoadSolution("../cf/xlibre/solutions/devuan.yaml")
	prj.Resolve()
}
