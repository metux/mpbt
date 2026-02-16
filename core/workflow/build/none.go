// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
)

type NoneBuilder struct {
	Package *model.Package
	Config  api.Entry
	pkgName string
}

func (ab *NoneBuilder) Init(p *model.Package, cf api.Entry) {
	ab.Package = p
	ab.Config = cf
	ab.pkgName = ab.Package.GetName()
}

func (ab *NoneBuilder) RunPrepare() error {
	return nil
}

func (ab *NoneBuilder) RunConfigure() error {
	return nil
}

func (ab *NoneBuilder) RunBuild() error {
	return nil
}

func (ab *NoneBuilder) RunInstall() error {
	return nil
}

func (ab *NoneBuilder) RunClean() error {
	return nil
}
