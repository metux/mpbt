package core

import (
	"log"
	"os"
	"strings"
)
type Project struct {
	Components ComponentMap
	Provides map[string]ComponentMap
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
		// FIXME: add support for multiple ones
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
		log.Printf("readdir() error on %s: %s\n", dirname, err)
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
					log.Printf("load error on %s: %s\n", dirname+"/"+n, err)
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

func (prj * Project) Resolve() {
	for _,b := range prj.Solution.Build {
		log.Printf("need to build: %s\n", b)
	}
}
