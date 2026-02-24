// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"fmt"
	"os"
	"strconv"
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

	env := os.Environ()
	env = append(env, "DESTDIR="+ab.Package.GetDestdir())

	args = append(args, cmake_args...)
	args = append(args, cmake_extra_args...)

	return ab.ExecInBuildDirEnv(args, env)
}

func (ab *CMakeBuilder) RunBuild() error {
	return ab.ExecInBuildDir([]string{"cmake", "--build", ".", "--parallel", strconv.Itoa(ab.Package.GetParallel())})
}

func (ab *CMakeBuilder) RunInstall() error {
	return ab.ExecInBuildDir([]string{"cmake", "--install", "."})
}

func (ab *CMakeBuilder) RunClean() error {
	return ab.RemoveBuildDir()
}
