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

func BuildWithBuilder(comp model.Package, b IBuilder) error {
	if _, err := os.Stat(comp.SourceDir + "/.DONE"); err == nil {
		log.Printf("[%s] Package already built\n", comp.Name)
		return nil
	}

	if err := b.RunPrepare(); err != nil {
		log.Println("Prepare error:", err)
		return err
	}

	if err := b.RunConfigure(); err != nil {
		log.Println("Configure error:", err)
		return err
	}

	if err := b.RunBuild(); err != nil {
		log.Println("Build error:", err)
		return err
	}

	if err := b.RunInstall(); err != nil {
		log.Println("Install error:", err)
		return err
	}

	if err := util.ExecCmd([]string{"touch", ".DONE"}, comp.SourceDir); err != nil {
		log.Println("Error:", err)
		return err
	}

	return nil
}
