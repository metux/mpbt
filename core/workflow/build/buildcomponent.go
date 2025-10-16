package build

import (
	"fmt"
	"log"
	"os"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"

//	"github.com/metux/go-metabuild/util/jobs"
)

func BuildComponent(comp model.Component) error {
	if !comp.IsBuildable() {
		log.Printf("%s is not buildable\n", comp.Name)
		return nil
	}

	log.Printf("building: %s\n", comp.Name)
	log.Printf("build system: %s\n", comp.BuildSystem)

	if comp.BuildSystem == "meson" {
		return BuildMeson(comp)
	}
	if comp.BuildSystem == "autotools" {
		return BuildAutotools(comp)
	}

	return fmt.Errorf("%s: no known build system defined: %s", comp.Name, comp.BuildSystem)
}

func BuildMeson(comp model.Component) error {
	log.Printf("building via meson: %s\n", comp.Name)
	return nil
}

func BuildAutotools(comp model.Component) error {
	if _, err := os.Stat(comp.SourceDir + "/.DONE"); err == nil {
		fmt.Println("Package %s already built.", comp.Name)
		return nil
	}

	if err := util.ExecCmd([]string{"./autogen.sh"}, comp.SourceDir); err != nil {
		fmt.Println("Error:", err)
		return err
	}

	if err := util.ExecCmd([]string{"./configure", fmt.Sprintf("--prefix=%s", comp.InstallPrefix)}, comp.SourceDir); err != nil {
		fmt.Println("Error:", err)
		return err
	}

	if err := util.ExecCmd([]string{"make"}, comp.SourceDir); err != nil {
		fmt.Println("Error:", err)
		return err
	}

	if err := util.ExecCmd([]string{"make", "install"}, comp.SourceDir); err != nil {
		fmt.Println("Error:", err)
		return err
	}

	if err := util.ExecCmd([]string{"touch", ".DONE"}, comp.SourceDir); err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}
