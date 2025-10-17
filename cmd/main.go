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
		SourceRoot:   abspath("../BUILD/sources"),
		Prefix:       abspath("../BUILD/DESTDIR"),
	}

	// FIXME: shall these also be defined in the solution ?
	if err := prj.LoadPackages("../cf/xlibre/packages", ""); err != nil {
		panic(fmt.Sprintf("error loading packages from %s\n", err))
	}

	prj.LoadSolution("../cf/xlibre/solutions/devuan.yaml")

	pkgconf := fmt.Sprintf(
		"%s/share/pkgconfig:%s/lib/pkgconfig:%s/lib/%s/pkgconfig/",
		prj.Prefix,
		prj.Prefix,
		prj.Prefix,
		prj.BuildMachine)

	log.Printf("machine=%s\npkgconf=%s\n", prj.BuildMachine, pkgconf)

	aclocal := fmt.Sprintf("%s/share/aclocal", prj.Prefix)

	os.Setenv("PKG_CONFIG_PATH", pkgconf)
	os.Setenv("ACLOCAL_PATH", aclocal)

	if err := fetch.FetchSource(&prj); err != nil {
		panic(fmt.Sprintf("fetch failed: %s\n", err))
	}
	if err := build.Build(&prj); err != nil {
		panic(fmt.Sprintf("build failed: %s\n", err))
	}
}
