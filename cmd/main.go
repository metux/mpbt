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

//	for n,c := range prj.Components {
//		log.Printf("Component: %s => %+v\n", n, c)
//	}

//	for n,c := range prj.Provides {
//		log.Printf("Provides: %s => %+v\n", n, c)
//	}

	prj.LoadSolution("../cf/xlibre/solutions/devuan.yaml")

//	log.Printf("Solution: %+v\n", prj.Solution)
	prj.Resolve()
}
