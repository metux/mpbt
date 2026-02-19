// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"fmt"
)

type CMakeBuilder struct {
	BuilderBase
}

func (ab *CMakeBuilder) RunPrepare() error {
	return ab.MakeBuildDir()
}

func (ab *CMakeBuilder) RunConfigure() error {
	args := []string{"cmake",
		ab.Package.GetSourceDir(),
		fmt.Sprintf("-DCMAKE_INSTALL_PREFIX=%s", ab.Package.GetInstallPrefix())}
	cmake_args := ab.Package.GetStrList("cmake-args")
	cmake_extra_args := ab.Package.GetStrList("cmake-extra-args")

	args = append(args, cmake_args...)
	args = append(args, cmake_extra_args...)

	return ab.ExecInBuildDir(args)
}

func (ab *CMakeBuilder) RunBuild() error {
	return ab.ExecInBuildDir([]string{"make", fmt.Sprintf("-j%d", ab.Package.GetParallel())})
}

func (ab *CMakeBuilder) RunInstall() error {
	return ab.ExecInBuildDir([]string{"make", "install"})
}

func (ab *CMakeBuilder) RunClean() error {
	return ab.RemoveBuildDir()
}
