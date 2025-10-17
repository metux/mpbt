package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/metux/go-magicdict/api"
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

	rootdir := abspath("../")
	machine := util.ExecOut([]string{"gcc", "-dumpmachine"})

	prj := model.Project{
		// FIXME: move this into the solution ?
		BuildMachine: machine,
		SourceRoot:   util.AppendPath(rootdir, "BUILD/sources"),
		Prefix:       util.AppendPath(rootdir, "BUILD/DESTDIR"),
	}

	// FIXME: shall these also be defined in the solution ?
	if err := prj.LoadPackages(util.AppendPath(rootdir, "cf/xlibre/packages"), ""); err != nil {
		panic(fmt.Sprintf("error loading packages from %s\n", err))
	}

	prj.LoadSolution(util.AppendPath(rootdir, "cf/xlibre/solutions/devuan.yaml"))
	api.SetStr(prj.Solution, "@rootdir", rootdir)
	api.SetStr(prj.Solution, "@workdir", util.AppendPath(rootdir, "BUILD"))
	api.SetStr(prj.Solution, "@machine", machine)

	for _,k := range api.GetKeys(prj.Solution, "env") {
		val := api.GetStr(prj.Solution, api.Key("env::"+string(k)))
		log.Printf("key=%s --> %s\n", string(k), val)
		os.Setenv(string(k), val)
	}

	if err := fetch.FetchSource(&prj); err != nil {
		panic(fmt.Sprintf("fetch failed: %s\n", err))
	}

	if err := build.Build(&prj); err != nil {
		panic(fmt.Sprintf("build failed: %s\n", err))
	}
}
