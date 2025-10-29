// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"fmt"
	"runtime"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

type AutotoolsBuilder struct {
	Package *model.Package
	pkgName string
}

func (ab *AutotoolsBuilder) Init(p *model.Package, cf api.Entry) {
	ab.Package = p
	ab.pkgName = p.GetName()
}

func (ab *AutotoolsBuilder) RunPrepare() error {
	return util.ExecCmd(ab.pkgName, []string{"./autogen.sh"}, ab.Package.GetSourceDir())
}

func (ab *AutotoolsBuilder) RunConfigure() error {
	return util.ExecCmd(ab.pkgName, []string{"./configure", fmt.Sprintf("--prefix=%s", ab.Package.GetInstallPrefix())}, ab.Package.GetSourceDir())
}

func (ab *AutotoolsBuilder) RunBuild() error {
	return util.ExecCmd(
		ab.pkgName,
		[]string{"make", fmt.Sprintf("-j%d", runtime.NumCPU())},
		ab.Package.GetSourceDir())
}

func (ab *AutotoolsBuilder) RunInstall() error {
	return util.ExecCmd(ab.pkgName, []string{"make", "install"}, ab.Package.GetSourceDir())
}

func (ab *AutotoolsBuilder) RunClean() error {
	return util.ExecCmd(ab.pkgName, []string{"make", "clean"}, ab.Package.GetSourceDir())
}
