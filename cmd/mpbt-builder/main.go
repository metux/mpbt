// SPDX-License-Identifier: AGPL-3.0-or-later
package main

import (
	"flag"
	"fmt"

	"github.com/metux/mpbt/core/util"
	"github.com/metux/mpbt/frontend"
)

func main() {
	prjDefines := make(util.MultiFlag, 0)
	solDefines := make(util.MultiFlag, 0)

	Config := frontend.BuildConfig{}

	flag.StringVar(&Config.SolutionFile, "solution", "", "Solution config file")
	flag.StringVar(&Config.RootDir, "root", ".", "Project root directory")
	flag.StringVar(&Config.WorkDir, "workdir", "", "Working directory")
	flag.Var(&prjDefines, "project-define", "define extra project variables")
	flag.Var(&solDefines, "solution-define", "define extra solution variables")
	flag.Parse()

	args := flag.Args()

	Config.SolutionDefines = solDefines
	Config.ProjectDefines = prjDefines

	if Config.SolutionFile == "" {
		fmt.Printf("missing -solution <fn> parameter\n")
		helppage()
	}

	if Config.RootDir == "" || len(args) == 0 {
		fmt.Printf("missing -root <dir> parameter\n")
		helppage()
	}

	if len(args) == 0 {
		helppage()
	}

	cmd := args[0]
	args = args[1:]

	if cmd == "build" {
		frontend.RunBuild(Config)
	} else {
		helppage()
	}
}
