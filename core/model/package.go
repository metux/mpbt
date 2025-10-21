package model

import (
	"log"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/go-magicdict/magic"
	"github.com/metux/mpbt/core/model/sources"
	"github.com/metux/mpbt/core/util"
)

const (
	KeyPackageBuildsystem = api.Key("buildsystem")
	KeyPackageSourceDir = api.Key("source-dir")
)

type Package struct {
	magic.MagicDict

	// internal only, not in YAML
	SourceDir     string `yaml:"-"`
	InstallPrefix string `yaml:"-"`
	Git * sources.Git `yaml:"-"`
}

type PackageMap = map[string]*Package

func (c *Package) LoadYaml(fn string) error {
	if err := util.LoadYaml(fn, c); err != nil {
		return err
	}
	d, err := magic.YamlLoad(fn, "")
	if err != nil {
		return err
	}
	c.MagicDict = d
	api.SetStr(d, "@filename", fn)
	return nil
}

func (c Package) GetAllDeps() util.StringList {
	return append(c.GetBuildDepends(), c.GetDepends()...)
}

// tell wether the component should/can be built
// eg. "system" type has nothing to build at all
func (c Package) IsBuildable() bool {
	t := c.GetType()
	return t != "system" && t != "fetchonly"
}

func (c Package) IsFetchable() bool {
	return c.GetGit() != nil
}

func (c Package) GetBuildsystem() string {
	return api.GetStr(c, KeyPackageBuildsystem)
}

func (c Package) GetType() string {
	return api.GetStr(c, "type")
}

func (c Package) GetDepends() [] string {
	return api.GetStrList(c, "depends")
}

func (c Package) GetBuildDepends() [] string {
	return api.GetStrList(c, "build-depends")
}

func (c Package) GetName() string {
	return api.GetStr(c, "name")
}

func (c Package) SetName(n string) {
	api.SetStr(c, "name", n)
}

func (c Package) GetProvides() []string {
	return api.GetStrList(c, "provides")
}

func (pkg Package) GetGit() * sources.Git {
	if pkg.Git != nil {
		return pkg.Git
	}

	ent, err := pkg.Get("sources::git")
	if err != nil {
		log.Printf("[%s] failed getting git entry: %+v\n", pkg.GetName(), err)
	}

	if ent == nil {
		return nil
	}

	git := sources.Git {
		Url: api.GetStr(ent, "url"),
		Ref: api.GetStr(ent, "ref"),
		Depth: api.GetInt(ent, "depth", 0),
		Fetch: api.GetStrList(ent, "fetch"),
	}

	pkg.Git = &git
	return pkg.Git
}

func (pkg Package) GetSourceDir() string {
	return api.GetStr(pkg, KeyPackageSourceDir)
}
