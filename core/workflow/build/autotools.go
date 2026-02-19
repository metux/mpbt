// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"fmt"
)

type AutotoolsBuilder struct {
	BuilderBase
}

func (ab *AutotoolsBuilder) RunPrepare() error {
	return ab.ExecInSourceDir([]string{"./autogen.sh"})
}

func (ab *AutotoolsBuilder) RunConfigure() error {
	return ab.ExecInSourceDir([]string{"./configure", fmt.Sprintf("--prefix=%s", ab.Package.GetInstallPrefix())})
}

func (ab *AutotoolsBuilder) RunBuild() error {
	return ab.ExecInSourceDir([]string{"make", fmt.Sprintf("-j%d", ab.Package.GetParallel())})
}

func (ab *AutotoolsBuilder) RunInstall() error {
	return ab.ExecInSourceDir([]string{"make", "install"})
}

func (ab *AutotoolsBuilder) RunClean() error {
	return ab.ExecInSourceDir([]string{"make", "clean"})
}
