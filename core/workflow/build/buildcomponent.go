package build

import (
	"fmt"
	"log"
	"os"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

func BuildPackage(comp model.Package) error {
	if !comp.IsBuildable() {
		log.Printf("%s is not buildable\n", comp.Name)
		return nil
	}

	if comp.BuildSystem == "meson" {
		return BuildWithBuilder(comp, &MesonBuilder{Package: &comp})
	}
	if comp.BuildSystem == "autotools" {
		return BuildWithBuilder(comp, &AutotoolsBuilder{Package: &comp})
	}

	return fmt.Errorf("%s: no known build system defined: %s", comp.Name, comp.BuildSystem)
}

func BuildWithBuilder(pkg model.Package, b model.IBuilder) error {
	if _, err := os.Stat(pkg.SourceDir + "/.DONE"); err == nil {
		log.Printf("[%s] Package already built\n", pkg.Name)
		return nil
	}

	if err := b.RunPrepare(); err != nil {
		log.Printf("[%s] Prepare error: %s\n", pkg.Name, err)
		return err
	}

	if err := b.RunConfigure(); err != nil {
		log.Printf("[%s] Configure error: %s\n", pkg.Name, err)
		return err
	}

	if err := b.RunBuild(); err != nil {
		log.Printf("[%s] Build error: %s\n",pkg.Name,  err)
		return err
	}

	if err := b.RunInstall(); err != nil {
		log.Printf("[%s] Install error: %s\n", pkg.Name, err)
		return err
	}

	if err := util.ExecCmd([]string{"touch", ".DONE"}, pkg.SourceDir); err != nil {
		log.Printf("[%s] Error: %s\n", pkg.Name, err)
		return err
	}

	return nil
}
