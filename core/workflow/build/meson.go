package build

import (
	"fmt"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

type MesonBuilder struct {
	Package *model.Package
	Config  api.Entry
}

func (ab *MesonBuilder) Init(p *model.Package, cf api.Entry) {
	ab.Package = p
	ab.Config = cf
}

func (ab *MesonBuilder) RunPrepare() error {
	//	util.ExecCmd([]string{"rm", "-Rf", "__BUILD"}, ab.Package.GetSourceDir())
	return util.ExecCmd([]string{"mkdir", "-p", "__BUILD"}, ab.Package.GetSourceDir())
}

func (ab *MesonBuilder) RunConfigure() error {
	args := []string{"meson",
		"setup",
		"__BUILD",
		fmt.Sprintf("--prefix=%s", ab.Package.GetInstallPrefix())}

	meson_args := api.GetStrList(ab.Config, "meson-args")
	if len(meson_args) == 0 {
		meson_args = api.GetStrList(ab.Package, "meson-args")
	}

	meson_extra_args := api.GetStrList(ab.Config, "meson-extra-args")
	if len(meson_extra_args) == 0 {
		meson_extra_args = api.GetStrList(ab.Package, "meson-extra-args")
	}

	args = append(args, meson_args...)
	args = append(args, meson_extra_args...)

	return util.ExecCmd(args, ab.Package.GetSourceDir())
}

func (ab *MesonBuilder) RunBuild() error {
	return util.ExecCmd([]string{"meson", "compile"}, ab.Package.GetSourceDir()+"/__BUILD")
}

func (ab *MesonBuilder) RunInstall() error {
	util.ExecCmd([]string{"meson", "install"}, ab.Package.GetSourceDir()+"/__BUILD")
	return nil
}

func (ab *MesonBuilder) RunClean() error {
	return util.ExecCmd([]string{"rm", "-Rf", "__BUILD"}, ab.Package.GetSourceDir())
}
