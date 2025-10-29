// SPDX-License-Identifier: AGPL-3.0-or-later
package model

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/go-magicdict/magic"
	"github.com/metux/mpbt/core/util"
)

const (
	KeyProjectInstallPrefix = "@installprefix"
	KeyProjectSourceRoot    = "@sourceroot"
	KeyProjectHomedir       = "@homedir"
	KeyProjectWorkdir       = "@workdir"
	KeyProjectMachine       = "@machine"
	KeyProjectRootDir       = "@rootdir"
	Project_Key_Solution    = "@SOLUTION"
)

type ProvidesMap map[string]PackageMap

type Project struct {
	magic.MagicDict
	Packages PackageMap
	Provides ProvidesMap
	Solution Solution
}

func (prj *Project) AddPackage(pkg *Package) {
	// init internal Package fields
	pkgName := pkg.GetName()
	pkg.SetProject(prj)

	if prj.Packages == nil {
		prj.Packages = make(PackageMap)
	}
	prj.Packages[pkgName] = pkg
	if prj.Provides == nil {
		prj.Provides = make(map[string]PackageMap)
	}

	for _, prov := range pkg.GetProvides() {
		if val, ok := prj.Provides[prov]; ok {
			// already have it
			val[pkgName] = pkg
		} else {
			newlist := make(PackageMap)
			newlist[pkgName] = pkg
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
				pkg, err := LoadPackageYaml(util.AppendPath(dirname, n), util.AppendPath(prefix, bn))
				if err != nil {
					return err
				}
				prj.AddPackage(pkg)
			}
		}
	}
	return nil
}

func (prj *Project) LoadSolution(fn string) error {
	if err := prj.Solution.LoadYaml(fn); err != nil {
		return err
	}

	prj.Solution.Put(Solution_Key_Project, prj)
	prj.Put(Project_Key_Solution, prj.Solution)

	log.Printf("project install-prefix=%s\n", api.GetStr(prj, KeyProjectInstallPrefix))
	log.Printf("solution install-prefix=%s\n", prj.Solution.GetStr(Solution_Key_InstallPrefix))

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

func (prj *Project) ApplyPackageConfigs() {
	for _, pkg := range prj.Packages {
		pkgName := pkg.GetName()
		if pconf := prj.Solution.GetPackageConfig(pkgName); pconf != nil {
			pkg := prj.Packages[pkgName]
			if pkg != nil {
				for _, key := range pconf.Keys() {
					value, err := pconf.Get(key)
					if err != nil {
						panic(fmt.Errorf("ApplyPackageConfigs: pkg=%s key=%s Get() error", pkgName, key, err))
					}
					if value != nil {
						pkg.Put(key, value)
					}
				}
			}
		}
	}
}

func (prj *Project) Init() {
	prj.MagicDict.Init()
	prj.SetMachine(util.ExecOut([]string{"gcc", "-dumpmachine"}))
	prj.SetRoot(".")
	api.SetDefaultStr(prj, KeyProjectWorkdir, "${" + KeyProjectRootDir + "}/WORK")
	api.SetDefaultStr(prj, KeyProjectSourceRoot, "${"+KeyProjectWorkdir+"}/sources")
	api.SetDefaultStr(prj, KeyProjectInstallPrefix, "${" + KeyProjectWorkdir + "}/DESTDIR")

	if home, err := os.UserHomeDir(); err == nil {
		api.SetDefaultStr(prj, KeyProjectHomedir, home)
	} else {
		panic(fmt.Sprintf("failed gettting homedir %s", err))
	}
}

func (prj *Project) SetWorkdir(wd string) {
	api.SetStr(prj, KeyProjectWorkdir, wd)
}

func (prj *Project) SetRoot(rootdir string) {
	r, _ := filepath.Abs(rootdir)
	api.SetStr(prj, KeyProjectRootDir, r)
}

func (prj *Project) SetMachine(machine string) {
	api.SetStr(prj, KeyProjectMachine, machine)
}

func (prj *Project) SetSourceRoot(dir string) {
	api.SetStr(prj, KeyProjectSourceRoot, dir)
}

func (prj *Project) GetSourceRoot() string {
	return api.GetStr(prj, KeyProjectSourceRoot)
}

func (prj *Project) PushEnv() {
	for _, k := range api.GetKeys(prj.Solution, "env") {
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
