package build

import (
	"fmt"
	"log"

	"github.com/metux/mpbt/core/model"
)

func buildComponent(prj * model.Project, name string) error {
	comp := prj.LookupComponent(name)
	if comp == nil {
		return fmt.Errorf("Cant resolve component %s\n", name)
	}

	for _, dep := range comp.GetAllDeps() {
		log.Printf("[%s] DEP: %s\n", comp.Name, dep)
		if err := buildComponent(prj, dep); err != nil {
			return err
		}
	}

	log.Printf("[%s] building component\n", name)
	return BuildComponent(*comp)
}

// FIXME: not honoring build flags yet
func Build(prj * model.Project) error {
	if prj.SourceRoot == "" {
		panic("prj.SourceRoot must not be empty")
	}

	for _, b := range prj.Solution.Build {
		if err := buildComponent(prj, b); err != nil {
			log.Printf("ERR on %s: %s\n", b, err)
		}
	}

	return nil
}
