package main

import (
	"log"

//	"github.com/metux/mpbt/core"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/workflow"
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

	workflow.FetchSource(&prj)
}
