package model

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/metux/go-magicdict/magic"
	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/util"
)

type ProvidesMap map[string]PackageMap

type Project struct {
	magic.MagicDict
	Packages     PackageMap
	Provides     ProvidesMap
	Solution     Solution
//	SourceRoot   string
	Prefix       string
}

func (prj *Project) AddPackage(comp *Package) {
	// init internal Package fields
	comp.SourceDir, _ = filepath.Abs(prj.GetSourceRoot() + "/" + comp.Name)
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
	if err := prj.Solution.LoadYaml(fn); err != nil {
		return err
	}

	prj.Solution.Put("@PROJECT", prj)
	return nil
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

func (prj *Project) Init() {
	prj.MagicDict.Init()
	prj.SetMachine(util.ExecOut([]string{"gcc", "-dumpmachine"}))
	prj.SetRoot(".")
	prj.SetWorkdir("${@rootdir}/BUILD")
	prj.SetSourceRoot("${@workdir}/sources")
}

func (prj *Project) SetWorkdir(wd string) {
	api.SetStr(prj, "@workdir", wd)
}

func (prj *Project) SetRoot(rootdir string) {
	r, _ := filepath.Abs(rootdir)
	api.SetStr(prj, "@rootdir", r)
}

func (prj *Project) SetMachine(machine string) {
	api.SetStr(prj, "@machine", machine)
}

func (prj *Project) SetSourceRoot(dir string) {
	api.SetStr(prj, "@sourceroot", dir)
}

func (prj *Project) GetSourceRoot() string {
	return api.GetStr(prj, "@sourceroot")
}
