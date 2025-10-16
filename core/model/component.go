package model

import (
	"github.com/metux/mpbt/core/model/sources"
	"github.com/metux/mpbt/core/util"
)

type Component struct {
	Name        string          `yaml:"name"`
	Provides    util.StringList `yaml:"provides"`
	Type        string          `yaml:"type"`
	BuildDepend util.StringList `yaml:"build-depends"`
	Depend      util.StringList `yaml:"depends"`
	Sources     sources.Sources `yaml:"sources"`
	BuildSystem string          `yaml:"buildsystem"`

	// internal only, not in YAML
	Filename string `yaml:"-"`
	SourceDir string `yaml:"-"`
	InstallPrefix string `yaml:"-"`
}

type ComponentMap = map[string]*Component

func (c *Component) LoadYaml(fn string) error {
	err := util.LoadYaml(fn, c)
	c.Filename = fn
	return err
}

func (c Component) GetAllDeps() util.StringList {
	return append(c.BuildDepend, c.Depend...)
}

// tell wether the component should/can be built
// eg. "system" type has nothing to build at all
func (c Component) IsBuildable() bool {
	return c.Type != "system" && c.Type != "fetchonly"
}

func (c Component) IsFetchable() bool {
	return c.Sources.Git != nil
}
