package main

import (
	"fmt"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
	"github.com/metux/mpbt/core/workflow/build"
	"github.com/metux/mpbt/core/workflow/fetch"
)

func main() {
	prj := model.MakeProject()
	prj.SetRoot("../")

	if err := prj.LoadSolution(util.AppendPath(prj.GetRoot(), "cf/xlibre/solutions/devuan.yaml")); err != nil {
		panic(fmt.Sprintf("failed loading solution: %s", err))
	}

	prj.PushEnv()

	if err := fetch.FetchSource(&prj); err != nil {
		panic(fmt.Sprintf("fetch failed: %s\n", err))
	}

	if err := build.Build(&prj); err != nil {
		panic(fmt.Sprintf("build failed: %s\n", err))
	}
}
