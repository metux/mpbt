package build

import (
	"fmt"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
	"github.com/metux/go-magicdict/api"
)

type AutotoolsBuilder struct {
	Package *model.Package
	Config api.Entry
}

func (ab *AutotoolsBuilder) Init(p *model.Package, cf api.Entry) {
	ab.Package = p
	ab.Config = cf
}

func (ab *AutotoolsBuilder) RunPrepare() error {
	return util.ExecCmd([]string{"./autogen.sh"}, ab.Package.SourceDir)
}

func (ab *AutotoolsBuilder) RunConfigure() error {
	return util.ExecCmd([]string{"./configure", fmt.Sprintf("--prefix=%s", ab.Package.InstallPrefix)}, ab.Package.SourceDir)
}

func (ab *AutotoolsBuilder) RunBuild() error {
	return util.ExecCmd([]string{"make"}, ab.Package.SourceDir)
}

func (ab *AutotoolsBuilder) RunInstall() error {
	return util.ExecCmd([]string{"make", "install"}, ab.Package.SourceDir)
}

func (ab *AutotoolsBuilder) RunClean() error {
	return util.ExecCmd([]string{"make", "clean"}, ab.Package.SourceDir)
}
