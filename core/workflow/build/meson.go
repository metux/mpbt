package build

import (
	"fmt"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
	"github.com/metux/go-magicdict/api"
)

type MesonBuilder struct {
	Package *model.Package
	Config api.Entry
}

func (ab *MesonBuilder) Init(p *model.Package, cf api.Entry) {
	ab.Package = p
	ab.Config = cf
}

func (ab *MesonBuilder) RunPrepare() error {
	//	util.ExecCmd([]string{"rm", "-Rf", "__BUILD"}, ab.Package.SourceDir)
	return util.ExecCmd([]string{"mkdir", "-p", "__BUILD"}, ab.Package.SourceDir)
}

func (ab *MesonBuilder) RunConfigure() error {
	args := []string{"meson",
		"setup",
		"__BUILD",
		fmt.Sprintf("--prefix=%s", ab.Package.InstallPrefix)}

	args = append(args, api.GetStrList(ab.Config, "meson-args")...)

	return util.ExecCmd(args, ab.Package.SourceDir)
}

func (ab *MesonBuilder) RunBuild() error {
	return util.ExecCmd([]string{"meson", "compile"}, ab.Package.SourceDir+"/__BUILD")
}

func (ab *MesonBuilder) RunInstall() error {
	util.ExecCmd([]string{"meson", "install"}, ab.Package.SourceDir+"/__BUILD")
	return nil
}

func (ab *MesonBuilder) RunClean() error {
	return util.ExecCmd([]string{"rm", "-Rf", "__BUILD"}, ab.Package.SourceDir)
}
