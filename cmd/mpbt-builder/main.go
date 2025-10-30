// SPDX-License-Identifier: AGPL-3.0-or-later
package main

import (
	"flag"
	"fmt"

	"github.com/metux/mpbt/frontend"
)

var cfSolution string
var cfRootDir string
var cfWorkDir string
var args []string
var prjDefines MultiFlag = make(MultiFlag, 0)
var solDefines MultiFlag = make(MultiFlag, 0)

var Config frontend.BuilderConfig

func main() {
	fmt.Printf("MPBT 0003\n")
	flag.StringVar(&cfSolution, "solution", "", "Solution config file")
	flag.StringVar(&cfRootDir, "root", ".", "Project root directory")
	flag.StringVar(&cfWorkDir, "workdir", "", "Working directory")
	flag.Var(&prjDefines, "project-define", "define extra project variables")
	flag.Var(&solDefines, "solution-define", "define extra solution variables")
	flag.Parse()
	args = flag.Args()

	if cfSolution == "" {
		fmt.Printf("missing -solution <fn> parameter\n")
		helppage()
	}

	if cfRootDir == "" || len(args) == 0 {
		fmt.Printf("missing -root <dir> parameter\n")
		helppage()
	}

	if len(args) == 0 {
		helppage()
	}

	cmd := args[0]
	args = args[1:]

	if cmd == "build" {
		do_build()
	} else {
		helppage()
	}
}
