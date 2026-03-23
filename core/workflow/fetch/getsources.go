// SPDX-License-Identifier: AGPL-3.0-or-later
package fetch

import (
	"fmt"
	"log"

	"github.com/metux/mpbt/core/model"
)

func fetchPackage(prj *model.Project, name string, update bool) error {
	pkg := prj.LookupPackage(name)
	if pkg == nil {
		return fmt.Errorf("Cant resolve component %s\n", name)
	}

	if pkg.GetBool("@fetch-done", false) {
		return nil
	}

	for _, dep := range pkg.GetAllDeps() {
		if err := fetchPackage(prj, dep, update); err != nil {
			return err
		}
	}

	if err := FetchPackage(pkg, update); err != nil {
		return err
	}

	pkg.SetBool("@fetch-done", true)
	return nil
}

// FIXME: not honoring build flags yet
func FetchSource(prj *model.Project, update bool) error {
	if prj.GetSourceRoot() == "" {
		panic("prj.SourceRoot must not be empty")
	}

	for _, b := range prj.Solution.GetBuildList() {
		if err := fetchPackage(prj, b, update); err != nil {
			log.Printf("ERR on %s: %s\n", b, err)
			return err
		}
	}

	return nil
}
