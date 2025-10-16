package main

import (
	"log"

//	"github.com/metux/mpbt/core"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/workflow/fetch"
	"github.com/metux/mpbt/core/workflow/build"
)

func main() {
	prj := model.Project{
		SourceRoot: "sources",
	}

	err := prj.LoadComponents("../cf/xlibre/components")
	if err != nil {
		log.Fatalf("error opening components directory: %s\n", err)
	}

	prj.LoadSolution("../cf/xlibre/solutions/devuan.yaml")

	fetch.FetchSource(&prj)
	build.Build(&prj)
}
