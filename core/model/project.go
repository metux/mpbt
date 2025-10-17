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
}

func (prj *Project) AddPackage(pkg *Package) {
	// init internal Package fields
	pkg.SourceDir = util.AppendPath(prj.GetSourceRoot(), pkg.Name)
	pkg.InstallPrefix = prj.GetInstallPrefix()

	if prj.Packages == nil {
		prj.Packages = make(PackageMap)
	}
	prj.Packages[pkg.Name] = pkg
	if prj.Provides == nil {
		prj.Provides = make(map[string]PackageMap)
	}

	for _, prov := range pkg.Provides {
		if val, ok := prj.Provides[prov]; ok {
			// already have it
			val[pkg.Name] = pkg
		} else {
			newlist := make(PackageMap)
			newlist[pkg.Name] = pkg
			prj.Provides[prov] = newlist
		}
	}
}

func (prj *Project) LoadPackages(dirname string, prefix string) error {
	log.Printf("[PROJECT] loading packages dir=%s prefix=%s\n", dirname, prefix)
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

func (prj *Project) LoadSolutionYAML(fn string) error {
	if err := prj.Solution.LoadYaml(fn); err != nil {
		return err
	}

	prj.Solution.Put("@PROJECT", prj)
	return nil
}

func (prj *Project) LoadSolution(fn string) error {
	if err := prj.LoadSolutionYAML(fn); err != nil {
		return nil
	}

	pkglist := prj.Solution.GetPackageSpecDirs()
	for _, p := range pkglist {
		if err := prj.LoadPackages(p, ""); err != nil {
			return err
		}
	}

	return nil
}


func (prj *Project) LookupPackage(name string) *Package {
	// apply mapping from solution
	name = prj.Solution.GetMapped(name)

	// try to find by exact package name
	if pkg, ok := prj.Packages[name]; ok {
		return pkg
	}

	// try by provides
	if pkglist, ok := prj.Provides[name]; ok {
		if len(pkglist) == 1 {
			for _, v := range pkglist {
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
	prj.SetInstallPrefix("${@workdir}/DESTDIR")
}

func (prj *Project) SetWorkdir(wd string) {
	api.SetStr(prj, "@workdir", wd)
}

func (prj *Project) SetRoot(rootdir string) {
	r, _ := filepath.Abs(rootdir)
	api.SetStr(prj, "@rootdir", r)
}

func (prj *Project) GetRoot() string {
	return api.GetStr(prj, "@rootdir")
}

func (prj *Project) SetMachine(machine string) {
	api.SetStr(prj, "@machine", machine)
}

func (prj *Project) SetSourceRoot(dir string) {
	api.SetStr(prj, "@sourceroot", dir)
}

func (prj *Project) SetInstallPrefix(dir string) {
	api.SetStr(prj, "@installprefix", dir)
}

func (prj *Project) GetSourceRoot() string {
	return api.GetStr(prj, "@sourceroot")
}

func (prj *Project) GetInstallPrefix() string {
	return api.GetStr(prj, "@installprefix")
}

func (prj *Project) PushEnv() {
	for _,k := range api.GetKeys(prj.Solution, "env") {
		val := api.GetStr(prj.Solution, api.Key("env::"+string(k)))
		log.Printf("[PROJECT] ENV: %s=%s\n", string(k), val)
		os.Setenv(string(k), val)
	}
}

func MakeProject() Project {
	prj := Project{}
	prj.Init()
	return prj
}
