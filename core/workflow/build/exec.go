// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

type ExecBuilder struct {
	Package *model.Package
	Config  api.Entry
	pkgName string
}

func (eb *ExecBuilder) Init(p *model.Package, cf api.Entry) {
	eb.Package = p
	eb.Config = cf
	eb.pkgName = eb.Package.GetName()
}

func (eb *ExecBuilder) RunPrepare() error {
	return eb.doExec("prepare")
}

func (eb *ExecBuilder) RunConfigure() error {
	return eb.doExec("configure")
}

func (eb *ExecBuilder) RunBuild() error {
	return eb.doExec("build")
}

func (eb *ExecBuilder) RunInstall() error {
	return eb.doExec("install")
}

func (eb *ExecBuilder) RunClean() error {
	return eb.doExec("clean")
}

func (eb *ExecBuilder) doExec(stage string) error {
	cmdline := eb.Package.GetStrList(api.Key("commands::" + stage))
	if len(cmdline) > 0 {
		return util.ExecCmd(eb.Package.GetName(), cmdline, eb.Package.GetSourceDir())
	}
	return nil
}
