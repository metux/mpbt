package model

import (
	"log"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/go-magicdict/magic"
	"github.com/metux/mpbt/core/model/sources"
	"github.com/metux/mpbt/core/util"
)

const (
	KeyPackageBuildDepends = "build-depends"
	KeyPackageBuildsystem  = "buildsystem"
	KeyPackageDepends      = "depends"
	KeyPackageFilename     = "@filename"
	KeyPackageName         = "name"
	KeyPackageProject      = "@PROJECT"
	KeyPackageProvides     = "provides"
	KeyPackageSourceDir    = "source-dir"
	KeyPackageType         = "type"
	KeyPackageInstallPrefix = "install-prefix"
)

type Package struct {
	magic.MagicDict

	// internal only, not in YAML
//	InstallPrefix string       `yaml:"-"`
	cacheGit           *sources.Git `yaml:"-"`
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
	api.SetStr(d, KeyPackageFilename, fn)
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
	return api.GetStr(c, KeyPackageType)
}

func (c Package) GetDepends() []string {
	return api.GetStrList(c, KeyPackageDepends)
}

func (c Package) GetBuildDepends() []string {
	return api.GetStrList(c, KeyPackageBuildDepends)
}

func (c Package) GetName() string {
	return api.GetStr(c, KeyPackageName)
}

func (c Package) SetName(n string) {
	api.SetStr(c, KeyPackageName, n)
}

func (c Package) GetProvides() []string {
	return api.GetStrList(c, KeyPackageProvides)
}

func (pkg Package) GetGit() *sources.Git {
	if pkg.cacheGit != nil {
		return pkg.cacheGit
	}

	ent, err := pkg.Get("sources::git")
	if err != nil {
		log.Printf("[%s] failed getting git entry: %+v\n", pkg.GetName(), err)
	}

	if ent == nil {
		return nil
	}

	git := sources.Git{
		Url:   api.GetStr(ent, "url"),
		Ref:   api.GetStr(ent, "ref"),
		Depth: api.GetInt(ent, "depth", 0),
		Fetch: api.GetStrList(ent, "fetch"),
	}

	pkg.cacheGit = &git
	return pkg.cacheGit
}

func (pkg Package) GetSourceDir() string {
	return api.GetStr(pkg, KeyPackageSourceDir)
}

func (pkg Package) SetSourceDir(src string) error {
	return api.SetStr(pkg, KeyPackageSourceDir, src)
}

func (pkg Package) GetInstallPrefix() string {
	return api.GetStr(pkg, KeyPackageInstallPrefix)
}
