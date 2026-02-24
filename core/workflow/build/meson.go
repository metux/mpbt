// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"fmt"
	"strconv"
)

type MesonBuilder struct {
	BuilderBase
}

func (ab *MesonBuilder) RunPrepare() error {
	return ab.MakeBuildDir()
}

func (ab *MesonBuilder) RunConfigure() error {
	args := []string{"meson",
		"setup",
		ab.Package.GetBuildDir(),
		fmt.Sprintf("--prefix=%s", ab.Package.GetInstallPrefix())}

	meson_args := ab.Package.GetStrList("meson-args")
	meson_extra_args := ab.Package.GetStrList("meson-extra-args")

	args = append(args, meson_args...)
	args = append(args, meson_extra_args...)

	return ab.ExecInSourceDir(args)
}

func (ab *MesonBuilder) RunBuild() error {
	return ab.ExecInBuildDir([]string{"meson", "compile", "-j", strconv.Itoa(ab.Package.GetParallel())})
}

func (ab *MesonBuilder) RunInstall() error {
	return ab.ExecInBuildDir([]string{"meson", "install", "--destdir", ab.Package.GetDestdir()})
}

func (ab *MesonBuilder) RunClean() error {
	return ab.RemoveBuildDir()
}
