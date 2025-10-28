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
	pkgName string
}

func (ab *MesonBuilder) Init(p *model.Package, cf api.Entry) {
	ab.Package = p
	ab.Config = cf
	ab.pkgName = ab.Package.GetName()
}

func (ab *MesonBuilder) RunPrepare() error {
	//	util.ExecCmd(ab.pkgName, []string{"rm", "-Rf", "__BUILD"}, ab.Package.GetSourceDir())
	return util.ExecCmd(ab.pkgName, []string{"mkdir", "-p", "__BUILD"}, ab.Package.GetSourceDir())
}

func (ab *MesonBuilder) RunConfigure() error {
	args := []string{"meson",
		"setup",
		"__BUILD",
		fmt.Sprintf("--prefix=%s", ab.Package.GetInstallPrefix())}

	meson_args := api.GetStrList(ab.Package, "meson-args")
	meson_extra_args := api.GetStrList(ab.Package, "meson-extra-args")

	args = append(args, meson_args...)
	args = append(args, meson_extra_args...)

	fmt.Printf("meson configure: source_dir=%s\n", ab.Package.GetSourceDir())
	return util.ExecCmd(ab.pkgName, args, ab.Package.GetSourceDir())
}

func (ab *MesonBuilder) RunBuild() error {
	fmt.Printf("meson build: source_dir=%s\n", ab.Package.GetSourceDir())
	return util.ExecCmd(ab.pkgName, []string{"meson", "compile"}, ab.Package.GetSourceDir()+"/__BUILD")
}

func (ab *MesonBuilder) RunInstall() error {
	fmt.Printf("meson install: source_dir=%s\n", ab.Package.GetSourceDir())
	util.ExecCmd(ab.pkgName, []string{"meson", "install"}, ab.Package.GetSourceDir()+"/__BUILD")
	return nil
}

func (ab *MesonBuilder) RunClean() error {
	return util.ExecCmd(ab.pkgName, []string{"rm", "-Rf", "__BUILD"}, ab.Package.GetSourceDir())
}
