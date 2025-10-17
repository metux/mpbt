package model

type IBuilder interface {
	SetPackage(p *Package)
	RunPrepare() error // eg. autogen.sh
	RunConfigure() error
	RunBuild() error
	RunInstall() error
}
