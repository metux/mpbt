package build

import (
	"log"

	"github.com/metux/mpbt/core/model"
)

func BuildComponent(comp model.Component) error {
	log.Printf("building: %s\n", comp.Name)
	log.Printf("build system: %s\n", comp.BuildSystem)

	if comp.BuildSystem == "" {
		return nil
	}
	if comp.BuildSystem == "meson" {
		return BuildMeson(comp)
	}

	return nil
}

func BuildMeson(comp model.Component) error {
	log.Printf("building via meson: %s\n", comp.Name)
	return nil
}
