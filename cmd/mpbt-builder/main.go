package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/workflow/build"
	"github.com/metux/mpbt/core/workflow/fetch"
)

var cfSolution string
var cfRootDir string
var args []string

func absdir(str string) string {
	s, e := filepath.Abs(str)
	if e != nil {
		panic(e)
	}
	return s
}

func helppage() {
	fmt.Printf("Usage: %s -solution <fn> -root <dir> [command...]\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Printf("Available commands:\n")
	fmt.Printf("    build           pull sources (once) and run build\n")
	os.Exit(1)
}

func main() {
	flag.StringVar(&cfSolution, "solution", "", "Solution config file")
	flag.StringVar(&cfRootDir, "root", ".", "Project root directory")
	flag.Parse()
	args = flag.Args()

	if cfSolution == "" || cfRootDir == "" {
		helppage()
	}

	prj := model.MakeProject()
	prj.SetRoot(absdir(cfRootDir))

	if err := prj.LoadSolution(absdir(cfSolution)); err != nil {
		panic(fmt.Sprintf("failed loading solution: %s", err))
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
