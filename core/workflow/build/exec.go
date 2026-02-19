// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"github.com/metux/go-magicdict/api"
)

type ExecBuilder struct {
	BuilderBase
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
		return eb.ExecInSourceDir(cmdline)
	}
	return nil
}
