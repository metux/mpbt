package main

import (
	"flag"
	"fmt"
)

var cfSolution string
var cfRootDir string
var args []string
var prjDefines MultiFlag = make(MultiFlag, 0)
var solDefines MultiFlag = make(MultiFlag, 0)

func main() {
	fmt.Printf("MPBT 0002\n")
	flag.StringVar(&cfSolution, "solution", "", "Solution config file")
	flag.StringVar(&cfRootDir, "root", ".", "Project root directory")
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
