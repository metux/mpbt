package build

import (
	"fmt"
	"log"

	"github.com/metux/mpbt/core/model"
)

func buildPackage(prj *model.Project, name string) error {
	pkg := prj.LookupPackage(name)
	if pkg == nil {
		return fmt.Errorf("Cant resolve component %s\n", name)
	}

	for _, dep := range pkg.GetAllDeps() {
		if err := buildPackage(prj, dep); err != nil {
			return err
		}
	}

	return BuildPackage(pkg, prj.Solution.GetPackageConfig(pkg.GetName()))
}

// FIXME: not honoring build flags yet
func Build(prj *model.Project) error {
	if prj.GetSourceRoot() == "" {
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
