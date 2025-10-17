package model

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/metux/mpbt/core/util"
)

type ProvidesMap map[string]PackageMap

type Project struct {
	Packages     PackageMap
	Provides     ProvidesMap
	Solution     Solution
	SourceRoot   string
	Prefix       string
	BuildMachine string
	HostMachine  string
}

func (prj *Project) AddPackage(comp *Package) {
	// init internal Package fields
	comp.SourceDir, _ = filepath.Abs(prj.SourceRoot + "/" + comp.Name)
	comp.InstallPrefix, _ = filepath.Abs(prj.Prefix)

	if prj.Packages == nil {
		prj.Packages = make(PackageMap)
	}
	prj.Packages[comp.Name] = comp
	if prj.Provides == nil {
		prj.Provides = make(map[string]PackageMap)
	}

	for _, prov := range comp.Provides {
		if val, ok := prj.Provides[prov]; ok {
			// already have it
			val[comp.Name] = comp
		} else {
			newlist := make(PackageMap)
			newlist[comp.Name] = comp
			prj.Provides[prov] = newlist
		}
	}
}

func (prj *Project) LoadPackages(dirname string, prefix string) error {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		n := entry.Name()
		if entry.IsDir() {
			if err := prj.LoadPackages(util.AppendPath(dirname, n), util.AppendPath(prefix, n)); err != nil {
				return err
			}
		} else {
			ext := filepath.Ext(n)
			bn := strings.TrimSuffix(n, ext)
			if ext == ".yaml" || ext == ".yml" {
				pkg := Package{}
				if e := pkg.LoadYaml(util.AppendPath(dirname, n)); e != nil {
					return e
				}
				pkg.Name = util.AppendPath(prefix, bn)
				prj.AddPackage(&pkg)
			}
		}
	}
	return nil
}

func (prj *Project) LoadSolution(fn string) error {
	return prj.Solution.LoadYaml(fn)
}

func (prj *Project) LookupPackage(name string) *Package {
	// apply mapping from solution
	name = prj.Solution.GetMapped(name)

	// try to find by exact package name
	if comp, ok := prj.Packages[name]; ok {
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
