package build

import (
	"github.com/metux/mpbt/core/model"
)

type IBuilder interface {
	SetPackage(p * model.Package)
	RunPrepare() error // eg. autogen.sh
	RunConfigure() error
	RunBuild() error
	RunInstall() error
}
