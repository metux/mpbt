package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
	"github.com/metux/mpbt/core/workflow/build"
	"github.com/metux/mpbt/core/workflow/fetch"
)

func abspath(p string) string {
	p2, _ := filepath.Abs(p)
	return p2
}

func main() {

	prj := model.Project{
		// FIXME: move this into the solution ?
		BuildMachine: util.ExecOut([]string{"gcc", "-dumpmachine"}),
		SourceRoot:   abspath("sources"),
		Prefix:       abspath("DESTDIR"),
	}

	pkgconf := fmt.Sprintf(
		"%s/share/pkgconfig:%s/lib/pkgconfig:%s/lib/%s/pkgconfig/",
		prj.Prefix,
		prj.Prefix,
		prj.Prefix,
		prj.BuildMachine)

	log.Printf("machine=%s\npkgconf=%s\n", prj.BuildMachine, pkgconf)

	os.Setenv("PKG_CONFIG_PATH", pkgconf)

	// FIXME: shall these also be defined in the solution ?
	err := prj.LoadPackages("../cf/xlibre/packages")
	if err != nil {
		log.Fatalf("error loading packages from %s\n", err)
	}

	prj.LoadSolution("../cf/xlibre/solutions/devuan.yaml")

	fetch.FetchSource(&prj)
	build.Build(&prj)
}
