// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"fmt"
	"runtime"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

type CMakeBuilder struct {
	Package *model.Package
	Config  api.Entry
	pkgName string
}

func (ab *CMakeBuilder) Init(p *model.Package, cf api.Entry) {
	ab.Package = p
	ab.Config = cf
	ab.pkgName = ab.Package.GetName()
}

func (ab *CMakeBuilder) RunPrepare() error {
	//	util.ExecCmd(ab.pkgName, []string{"rm", "-Rf", ab.Package.GetBuildDir()}, ab.Package.GetSourceDir())
	return util.ExecCmd(ab.pkgName, []string{"mkdir", "-p", ab.Package.GetBuildDir()}, ab.Package.GetSourceDir())
}

func (ab *CMakeBuilder) RunConfigure() error {
	args := []string{"cmake",
		ab.Package.GetSourceDir(),
		fmt.Sprintf("-DCMAKE_INSTALL_PREFIX=%s", ab.Package.GetInstallPrefix())}
	cmake_args := ab.Package.GetStrList("cmake-args")
	cmake_extra_args := ab.Package.GetStrList("cmake-extra-args")

	args = append(args, cmake_args...)
	args = append(args, cmake_extra_args...)

	return util.ExecCmd(ab.pkgName, args, ab.Package.GetBuildDir())
}

func (ab *CMakeBuilder) RunBuild() error {
	return util.ExecCmd(
		ab.pkgName,
		[]string{"make", fmt.Sprintf("-j%d", runtime.NumCPU())},
		ab.Package.GetBuildDir())
}

func (ab *CMakeBuilder) RunInstall() error {
	return util.ExecCmd(ab.pkgName, []string{"make", "install"}, ab.Package.GetBuildDir())
}

func (ab *CMakeBuilder) RunClean() error {
	return util.ExecCmd(ab.pkgName, []string{"rm", "-Rf", ab.Package.GetBuildDir()}, ab.Package.GetSourceDir())
}
