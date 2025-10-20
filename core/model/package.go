package model

import (
	"github.com/metux/go-magicdict/magic"
	"github.com/metux/mpbt/core/model/sources"
	"github.com/metux/mpbt/core/util"
)

type Package struct {
	magic.MagicDict
	Name        string          `yaml:"name"`
	Provides    util.StringList `yaml:"provides"`
	Type        string          `yaml:"type"`
	BuildDepend util.StringList `yaml:"build-depends"`
	Depend      util.StringList `yaml:"depends"`
	Sources     sources.Sources `yaml:"sources"`
	BuildSystem string          `yaml:"buildsystem"`

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
	return append(c.BuildDepend, c.Depend...)
}

// tell wether the component should/can be built
// eg. "system" type has nothing to build at all
func (c Package) IsBuildable() bool {
	return c.Type != "system" && c.Type != "fetchonly"
}

func (c Package) IsFetchable() bool {
	return c.Sources.Git != nil
}
