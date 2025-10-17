package model

import (
	"github.com/metux/mpbt/core/util"
)

type Solution struct {
	PackageMapping map[string]string  `yaml:"package-mapping"`
	Filename         string           `yaml:"-"`
	Build            util.StringList  `yaml:"build"`
}

func (c *Solution) LoadYaml(fn string) error {
	err := util.LoadYaml(fn, c)
	c.Filename = fn
	return err
}

func (c *Solution) GetMapped(name string) string {
	if c.PackageMapping == nil {
		return name
	}
	if val, ok := c.PackageMapping[name]; ok {
		return val
	}
	return name
}
