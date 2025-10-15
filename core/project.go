package core

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type ProvidesMap map[string]ComponentMap

type Project struct {
	Components ComponentMap
	Provides ProvidesMap
	Solution Solution
}

func (db * Project) Add(comp *Component) {
	if db.Components == nil {
		db.Components = make(ComponentMap)
	}
	db.Components[comp.Name] = comp
	if db.Provides == nil {
		db.Provides = make(map[string]ComponentMap)
	}

	for _, prov := range comp.Provides {
		if val, ok := db.Provides[prov]; ok {
			// already have it
			val[comp.Name] = comp
		} else {
			newlist := make(ComponentMap)
			newlist[comp.Name] = comp
			db.Provides[prov] = newlist
		}
	}
}

func (db * Project) LoadComponents(dirname string) error {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		n := entry.Name()
		if entry.IsDir() {
			db.LoadComponents(dirname+"/"+n)
		} else {
			if strings.HasSuffix(n, ".yaml") || strings.HasSuffix(n, ".yml") {
				comp := Component{}
				err = comp.LoadYaml(dirname + "/" + n)
				if err != nil {
					return err
				}
				db.Add(&comp)
			}
		}
	}
	return nil
}

func (prj * Project) LoadSolution(fn string) error {
	return prj.Solution.LoadYaml(fn)
}

func (prj * Project) LookupComponent(name string) *Component {
	// apply mapping from solution
	name = prj.Solution.GetMapped(name)

	// try to find by exact package name
	if comp, ok := prj.Components[name]; ok {
		return comp
	}

	// try by provides
	if complist, ok := prj.Provides[name]; ok {
		if (len(complist) == 1) {
			for _, v := range complist {
				return v
			}
		} else {
			log.Printf("ERR: multiple results\n");
			return nil
		}
	}

	log.Printf("NOT FOUND %s\n", name)
	return nil
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

func (prj * Project) CloneComponent(comp * Component) error {
	log.Printf("cloning component %s\n%+v\n", comp.Name, comp)
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
