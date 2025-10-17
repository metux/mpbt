package build

import (
	"fmt"
	"log"

	"github.com/metux/mpbt/core/model"
)

func buildPackage(prj *model.Project, name string) error {
	comp := prj.LookupPackage(name)
	if comp == nil {
		return fmt.Errorf("Cant resolve component %s\n", name)
	}

	for _, dep := range comp.GetAllDeps() {
		if err := buildPackage(prj, dep); err != nil {
			return err
		}
	}

	return BuildPackage(*comp)
}

// FIXME: not honoring build flags yet
func Build(prj *model.Project) error {
	if prj.SourceRoot == "" {
		panic("prj.SourceRoot must not be empty")
	}

	for _, b := range prj.Solution.GetBuildList() {
		if err := buildPackage(prj, b); err != nil {
			log.Printf("BUILD ERR on %s: %s\n", b, err)
			return err
		}
	}

	return nil
}
