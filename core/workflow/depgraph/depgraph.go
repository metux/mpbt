// SPDX-License-Identifier: AGPL-3.0-or-later
package depgraph

import (
	"fmt"
	"log"

	"github.com/metux/mpbt/core/model"
)

func graphPackage(prj *model.Project, pkg *model.Package) error {
	if pkg.GetBool("@visited", false) {
		return nil
	}
	pkg.SetBool("@visited", true)

	for _, dep := range pkg.GetAllDeps() {
		deppkg := prj.LookupPackage(dep)
		if deppkg == nil {
			return fmt.Errorf("Cant resolve component %s\n", dep)
		}

		if err := graphPackage(prj, deppkg); err != nil {
			return err
		}
		fmt.Printf("    \"%s\" -> \"%s\";\n", pkg.GetName(), deppkg.GetName())
	}

	return nil
}

// FIXME: not honoring build flags yet
func DepGraph(prj *model.Project) error {

	fmt.Printf("digraph A {\n")
	fmt.Printf("    rankdir=LR;\n")
	fmt.Printf("    nodesep=0.5;\n")
	fmt.Printf("    ranksep=1.0;\n")
	fmt.Printf("    node [shape=box, style=filled, fillcolor=lightyellow, fontsize=11, fontname=\"Arial\"];\n")
	fmt.Printf("    edge [arrowsize=0.8];\n")

	for _, b := range prj.Solution.GetBuildList() {
		pkg := prj.LookupPackage(b)
		if pkg == nil {
			return fmt.Errorf("Cant resolve component %s\n", b)
		}
		if err := graphPackage(prj, pkg); err != nil {
			log.Printf("ERR on %s: %s\n", b, err)
			return err
		}
	}

	fmt.Printf("}\n")

	return nil
}
