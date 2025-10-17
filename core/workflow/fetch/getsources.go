package fetch

import (
	"fmt"
	"log"

	"github.com/metux/mpbt/core/model"
)

func fetchComponent(prj * model.Project, name string) error {
	comp := prj.LookupPackage(name)
	if comp == nil {
		return fmt.Errorf("Cant resolve component %s\n", name)
	}

	for _, dep := range comp.GetAllDeps() {
		if err := fetchComponent(prj, dep); err != nil {
			return err
		}
	}

	if comp.Sources.Git == nil {
		return nil
	}

	return CloneComponent(*comp)
}

// FIXME: not honoring build flags yet
func FetchSource(prj * model.Project) error {
	if prj.SourceRoot == "" {
		panic("prj.SourceRoot must not be empty")
	}

	for _, b := range prj.Solution.Build {
		if err := fetchComponent(prj, b); err != nil {
			log.Printf("ERR on %s: %s\n", b, err)
		}
	}

	return nil
}
