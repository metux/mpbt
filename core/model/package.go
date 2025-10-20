package model

import (
	"github.com/metux/go-magicdict/api"
	"github.com/metux/go-magicdict/magic"
	"github.com/metux/mpbt/core/model/sources"
	"github.com/metux/mpbt/core/util"
)

type Package struct {
	magic.MagicDict
	Provides    util.StringList `yaml:"provides"`
	Sources     sources.Sources `yaml:"sources"`

	// internal only, not in YAML
	Filename      string `yaml:"-"`
	SourceDir     string `yaml:"-"`
	InstallPrefix string `yaml:"-"`
}

type PackageMap = map[string]*Package

func (c *Package) LoadYaml(fn string) error {
	c.Filename = fn
	if err := util.LoadYaml(fn, c); err != nil {
		return err
	}
	d, err := magic.YamlLoad(fn, "")
	if err != nil {
		return err
	}
	c.MagicDict = d
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
	return c.Sources.Git != nil
}

func (c Package) GetBuildsystem() string {
	return api.GetStr(c, "buildsystem")
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
