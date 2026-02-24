// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"os"

	"github.com/metux/go-magicdict/api"
)

type ExecBuilder struct {
	BuilderBase
}

func (eb *ExecBuilder) RunPrepare() error {
	return eb.doExec("prepare", os.Environ())
}

func (eb *ExecBuilder) RunConfigure() error {
	return eb.doExec("configure", os.Environ())
}

func (eb *ExecBuilder) RunBuild() error {
	return eb.doExec("build", os.Environ())
}

func (eb *ExecBuilder) RunInstall() error {
	env := os.Environ()
	env = append(env, "DESTDIR="+eb.Package.GetDestdir())

	return eb.doExec("install", env)
}

func (eb *ExecBuilder) RunClean() error {
	return eb.doExec("clean", os.Environ())
}

func (eb *ExecBuilder) doExec(stage string, env []string) error {
	cmdline := eb.Package.GetStrList(api.Key("commands::" + stage))
	if len(cmdline) > 0 {
		return eb.ExecInSourceDirEnv(cmdline, env)
	}
	return nil
}
