package model

import (
	"log"
	"os"
	"strings"
)

type ProvidesMap map[string]ComponentMap

type Project struct {
	Components ComponentMap
	Provides   ProvidesMap
	Solution   Solution
	SourceRoot string
}

func (prj *Project) AddComponent(comp *Component) {
	if prj.Components == nil {
		prj.Components = make(ComponentMap)
	}
	prj.Components[comp.Name] = comp
	if prj.Provides == nil {
		prj.Provides = make(map[string]ComponentMap)
	}

	for _, prov := range comp.Provides {
		if val, ok := prj.Provides[prov]; ok {
			// already have it
			val[comp.Name] = comp
		} else {
			newlist := make(ComponentMap)
			newlist[comp.Name] = comp
			prj.Provides[prov] = newlist
		}
	}
}

func (prj *Project) LoadComponents(dirname string) error {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		n := entry.Name()
		if entry.IsDir() {
			prj.LoadComponents(dirname + "/" + n)
		} else {
			if strings.HasSuffix(n, ".yaml") || strings.HasSuffix(n, ".yml") {
				comp := Component{}
				if e := comp.LoadYaml(dirname + "/" + n); e != nil {
					return e
				}
				comp.SourceDir = prj.SourceRoot + "/" + comp.Name
				prj.AddComponent(&comp)
			}
		}
	}
	return nil
}

func (prj *Project) LoadSolution(fn string) error {
	return prj.Solution.LoadYaml(fn)
}

func (prj *Project) LookupComponent(name string) *Component {
	// apply mapping from solution
	name = prj.Solution.GetMapped(name)

	// try to find by exact package name
	if comp, ok := prj.Components[name]; ok {
		return comp
	}

	// try by provides
	if complist, ok := prj.Provides[name]; ok {
		if len(complist) == 1 {
			for _, v := range complist {
				return v
			}
		} else {
			log.Printf("ERR: multiple results\n")
			return nil
		}
	}

	log.Printf("NOT FOUND %s\n", name)
	return nil
}
