// SPDX-License-Identifier: AGPL-3.0-or-later
package frontend

import (
	"fmt"
	"log"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/workflow/build"
	"github.com/metux/mpbt/core/workflow/fetch"
)

func RunBuild(cf BuildConfig) {
	prj := model.MakeProject()
	prj.SetRoot(cf.RootDir)

	for k, v := range cf.ProjectDefines {
		api.SetStr(prj, api.Key(k), v)
	}

	log.Printf("loading solution: %s\n", cf.SolutionFile)

	if err := prj.LoadSolution(cf.SolutionFile); err != nil {
		panic(fmt.Sprintf("failed loading solution: %s", err))
	}

	for k, v := range cf.SolutionDefines {
		prj.Solution.SetStr(api.Key(k), v)
	}

	if cf.WorkDir != "" {
		prj.SetWorkdir(cf.WorkDir)
	}

	log.Printf("applying package config ...\n")

	prj.ApplyPackageConfigs()
	prj.PushEnv() // FIXME: should be done per exec

	log.Printf("fetching sources ...\n")
	if err := fetch.FetchSource(&prj); err != nil {
		panic(fmt.Sprintf("fetch failed: %s\n", err))
	}

	log.Printf("building ...\n")
	if err := build.Build(&prj); err != nil {
		panic(fmt.Sprintf("build failed: %s\n", err))
	}
}
