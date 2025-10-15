package core

import (
	"fmt"
	"log"

	"github.com/metux/mpbt/core/model"
)

type Project struct {
	model.Project
}

// resolve what we need to clone
func (prj * Project) ResolvePkg(name string) error {
	comp := prj.LookupComponent(name)
	if comp == nil {
		return fmt.Errorf("Cant resolve component %s\n", name)
	}

	for _, dep := range comp.BuildDepend {
		if err := prj.ResolvePkg(dep); err != nil {
			return err
		}
	}

	for _, dep := range comp.Depend {
		if err := prj.ResolvePkg(dep); err != nil {
			return err
		}
	}

	log.Printf("HANDLING: %s\n", name)
	return prj.CloneComponent(comp)
}

func (prj * Project) CloneComponent(comp * model.Component) error {
	log.Printf("cloning component %s\n%+v\n", comp.Name, comp)
	if (comp.Sources == nil) {
		log.Printf("no need to clone it\n");
	} else {
		log.Println("should clone it\n");
	}

	return nil
}

// FIXME: yet need to check for recursion and feature flags
func (prj * Project) Resolve() {
	for _,b := range prj.Solution.Build {
		if err := prj.ResolvePkg(b); err != nil {
			log.Printf("ERR on %s: %s\n", b, err)
		}
	}
}
