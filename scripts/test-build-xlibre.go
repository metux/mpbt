// SPDX-License-Identifier: AGPL-3.0-or-later
package main

import (
	"github.com/metux/mpbt/frontend"
)

func main() {
	Config := frontend.BuildConfig{
		SolutionFile:    "cf/xlibre/solutions/devuan.yaml",
		RootDir:         ".",
		WorkDir:         "WORK",
		ProjectDefines:  make(map[string]string),
		SolutionDefines: make(map[string]string)}

	Config.ProjectDefines["xlibre-git"] = "git@github.com:X11Libre"

	frontend.RunBuild(Config)
}
