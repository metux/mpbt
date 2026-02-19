// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"fmt"
	"strconv"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

type MesonBuilder struct {
	Package *model.Package
	Config  api.Entry
	pkgName string
}

func (ab *MesonBuilder) Init(p *model.Package, cf api.Entry) {
	ab.Package = p
	ab.Config = cf
	ab.pkgName = ab.Package.GetName()
}

func (ab *MesonBuilder) RunPrepare() error {
	//	util.ExecCmd(ab.pkgName, []string{"rm", "-Rf", ab.Package.GetBuildDir()}, ab.Package.GetSourceDir())
	return util.ExecCmd(ab.pkgName, []string{"mkdir", "-p", ab.Package.GetBuildDir()}, ab.Package.GetSourceDir())
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

	return util.ExecCmd(ab.pkgName, args, ab.Package.GetSourceDir())
}

func (ab *MesonBuilder) RunBuild() error {
	return util.ExecCmd(
		ab.pkgName,
		[]string{"meson", "compile", "-j", strconv.Itoa(ab.Package.GetParallel())},
		ab.Package.GetBuildDir())
}

func (ab *MesonBuilder) RunInstall() error {
	return util.ExecCmd(ab.pkgName, []string{"meson", "install"}, ab.Package.GetBuildDir())
}

func (ab *MesonBuilder) RunClean() error {
	return util.ExecCmd(ab.pkgName, []string{"rm", "-Rf", ab.Package.GetBuildDir()}, ab.Package.GetSourceDir())
}
