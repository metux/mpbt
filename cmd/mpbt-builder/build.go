// SPDX-License-Identifier: AGPL-3.0-or-later
package main

import (
	"fmt"
	"log"

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

	log.Printf("loading solution: %s\n", cfSolution)

	if err := prj.LoadSolution(cfSolution); err != nil {
		panic(fmt.Sprintf("failed loading solution: %s", err))
	}

	for k, v := range solDefines {
		prj.Solution.SetStr(api.Key(k), v)
	}

	if cfWorkDir != "" {
		prj.SetWorkdir(cfWorkDir)
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
