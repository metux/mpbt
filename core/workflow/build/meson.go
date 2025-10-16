package build

import (
//	"fmt"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

type MesonBuilder struct {
	Package * model.Package
}

func (ab * MesonBuilder) SetPackage (p * model.Package) {
	ab.Package = p
}

func (ab * MesonBuilder) RunPrepare() error {
//	util.ExecCmd([]string{"rm", "-Rf", "__BUILD"}, ab.Package.SourceDir)
	return util.ExecCmd([]string{"mkdir", "-p", "__BUILD"}, ab.Package.SourceDir)
}

func (ab * MesonBuilder) RunConfigure() error {
	return util.ExecCmd([]string{"meson", "setup", "__BUILD"}, ab.Package.SourceDir)
}

func (ab * MesonBuilder) RunBuild() error {
	return util.ExecCmd([]string{"meson", "compile"}, ab.Package.SourceDir + "/__BUILD")
}

func (ab * MesonBuilder) RunInstall() error {
	util.ExecCmd([]string{"meson", "install"}, ab.Package.SourceDir + "/__BUILD")
	return nil
}

func (ab * MesonBuilder) RunClean() error {
	return util.ExecCmd([]string{"rm", "-Rf", "__BUILD"}, ab.Package.SourceDir)
}
