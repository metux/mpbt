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

	// internal only, not in YAML
	Filename string `yaml:"-"`
	CloneDir string `yaml:"-"`
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
