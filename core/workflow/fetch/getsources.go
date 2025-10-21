package fetch

import (
	"fmt"
	"log"

	"github.com/metux/mpbt/core/model"
)

func fetchPackage(prj *model.Project, name string) error {
	pkg := prj.LookupPackage(name)
	if pkg == nil {
		return fmt.Errorf("Cant resolve component %s\n", name)
	}

	for _, dep := range pkg.GetAllDeps() {
		if err := fetchPackage(prj, dep); err != nil {
			return err
		}
	}

	if pkg.GetGit() == nil {
		return nil
	}

	return ClonePackage(*pkg, prj.Solution.GetPackageConfig(pkg.GetName()))
}

// FIXME: not honoring build flags yet
func FetchSource(prj *model.Project) error {
	if prj.GetSourceRoot() == "" {
		panic("prj.SourceRoot must not be empty")
	}

	for _, b := range prj.Solution.GetBuildList() {
		if err := fetchPackage(prj, b); err != nil {
			log.Printf("ERR on %s: %s\n", b, err)
			return err
		}
	}

	return nil
}
