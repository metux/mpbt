package core

import (
//	"log"
//	"os"
//	"strings"
)

type Solution struct {
	ComponentMapping map[string]string `yaml:"component-mapping"`
	Filename string `yaml:"_"`
}

func (c *Solution) LoadYaml(fn string) error {
	err := LoadYaml(fn, c)
	c.Filename = fn
	return err
}
