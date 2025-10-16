package build

import (
	"fmt"
	"log"
	"os"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"

//	"github.com/metux/go-metabuild/util/jobs"
)

func BuildComponent(comp model.Package) error {
	if !comp.IsBuildable() {
		log.Printf("%s is not buildable\n", comp.Name)
		return nil
	}

	log.Printf("building: %s\n", comp.Name)
	log.Printf("build system: %s\n", comp.BuildSystem)

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
		fmt.Println("Package %s already built.", comp.Name)
		return nil
	}

	if err := b.RunPrepare(); err != nil {
		fmt.Println("Prepare error:", err)
		return err
	}

	if err := b.RunConfigure(); err != nil {
		fmt.Println("Configure error:", err)
		return err
	}

	if err := b.RunBuild(); err != nil {
		fmt.Println("Build error:", err)
		return err
	}

	if err := b.RunInstall(); err != nil {
		fmt.Println("Install error:", err)
		return err
	}

	if err := util.ExecCmd([]string{"touch", ".DONE"}, comp.SourceDir); err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}
