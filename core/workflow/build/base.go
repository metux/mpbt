// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

type BuilderBase struct {
	Package *model.Package
	Config  api.Entry
	PkgName string
}

func (b *BuilderBase) Init(p *model.Package, cf api.Entry) {
	b.Package = p
	b.Config = cf
	b.PkgName = p.GetName()
}

func (b BuilderBase) ExecInSourceDir(cmdline []string) error {
	return util.ExecCmd(b.PkgName, cmdline, b.Package.GetSourceDir())
}

func (b BuilderBase) ExecInBuildDir(cmdline []string) error {
	return util.ExecCmd(b.PkgName, cmdline, b.Package.GetBuildDir())
}

func (b BuilderBase) MakeBuildDir() error {
	return b.ExecInSourceDir([]string{"mkdir", "-p", b.Package.GetBuildDir()})
}

func (b BuilderBase) RemoveBuildDir() error {
	return b.ExecInSourceDir([]string{"rm", "-Rf", b.Package.GetBuildDir()})
}
