package main

import (
	"fmt"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/workflow/build"
	"github.com/metux/mpbt/core/workflow/fetch"
)

func do_build() {
	prj := model.MakeProject()
	prj.SetRoot(cfRootDir)

	for k, v := range prjDefines {
		api.SetStr(prj, api.Key(k), v)
	}

	if err := prj.LoadSolution(cfSolution); err != nil {
		panic(fmt.Sprintf("failed loading solution: %s", err))
	}

	for k, v := range solDefines {
		api.SetStr(prj.Solution, api.Key(k), v)
	}

	prj.ApplyPackageConfigs()
	prj.PushEnv() // FIXME: should be done per exec

	if err := fetch.FetchSource(&prj); err != nil {
		panic(fmt.Sprintf("fetch failed: %s\n", err))
	}

	if err := build.Build(&prj); err != nil {
		panic(fmt.Sprintf("build failed: %s\n", err))
	}
}
