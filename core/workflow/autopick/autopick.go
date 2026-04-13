// SPDX-License-Identifier: AGPL-3.0-or-later
package autopick

import (
	"fmt"
	"log"

	"github.com/metux/mpbt/core/model"
)

const (
    DoneFlag = "@@autopick-visited"
)

func autopickPackage(prj *model.Project, pkg *model.Package) error {
	if pkg.GetBool(DoneFlag, false) {
		log.Printf("[%s] already done\n", pkg.GetName())
		return nil
	}
	pkg.SetBool(DoneFlag, true)
	log.Printf("[%s] doing autopick ...\n", pkg.GetName())

	for _, dep := range pkg.GetAllDeps() {
		deppkg := prj.LookupPackage(dep)
		if deppkg == nil {
			return fmt.Errorf("Cant resolve component %s\n", dep)
		}

		if err := autopickPackage(prj, deppkg); err != nil {
			return err
		}
	}

	return nil
}

func AutoPick(prj *model.Project) error {
	for _, b := range prj.Solution.GetBuildList() {
		pkg := prj.LookupPackage(b)
		if pkg == nil {
			return fmt.Errorf("Cant resolve component %s\n", b)
		}
		if err := autopickPackage(prj, pkg); err != nil {
			log.Printf("ERR on %s: %s\n", b, err)
			return err
		}
	}

	return nil
}
