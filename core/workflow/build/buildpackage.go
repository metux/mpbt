package build

import (
	"fmt"
	"log"
	"os"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

func BuildPackage(pkg *model.Package, cf api.Entry) error {
	if !pkg.IsBuildable() {
		return nil
	}

	bs := pkg.GetBuildsystem()
	if bs == "meson" {
		return BuildWithBuilder(pkg, cf, &MesonBuilder{})
	}
	if bs == "autotools" {
		return BuildWithBuilder(pkg, cf, &AutotoolsBuilder{})
	}

	return fmt.Errorf("%s: no known build system defined: %s", pkg.GetName(), bs)
}

func BuildWithBuilder(pkg *model.Package, cf api.Entry, b model.IBuilder) error {
	b.Init(pkg, cf)

	pkgName := pkg.GetName()

	if _, err := os.Stat(pkg.GetSourceDir() + "/.DONE"); err == nil {
		log.Printf("[%s] Package already built\n", pkgName)
		return nil
	}

	if err := b.RunPrepare(); err != nil {
		log.Printf("[%s] Prepare error: %s\n", pkgName, err)
		return err
	}

	if err := b.RunConfigure(); err != nil {
		log.Printf("[%s] Configure error: %s\n", pkgName, err)
		return err
	}

	if err := b.RunBuild(); err != nil {
		log.Printf("[%s] Build error: %s\n", pkgName, err)
		return err
	}

	if err := b.RunInstall(); err != nil {
		log.Printf("[%s] Install error: %s\n", pkgName, err)
		return err
	}

	if err := util.ExecCmd([]string{"touch", ".DONE"}, pkg.GetSourceDir()); err != nil {
		log.Printf("[%s] Error: %s\n", pkgName, err)
		return err
	}

	return nil
}
