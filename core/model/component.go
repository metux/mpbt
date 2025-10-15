package model

import (
	"github.com/metux/mpbt/core/model/sources"
	"github.com/metux/mpbt/core/util"
)

type Component struct {
	Name     string `yaml:"name"`
	Provides util.StringList `yaml:"provides"`
	Type     string `yaml:"type"`
	Filename string `yaml:"_"`
	BuildDepend util.StringList `yaml:"build-depends"`
	Depend util.StringList `yaml:"depends"`
	Sources * sources.Sources `yaml:"sources"`
}

type ComponentMap = map[string]*Component

func (c *Component) LoadYaml(fn string) error {
	err := util.LoadYaml(fn, c)
	c.Filename = fn
	return err
}
