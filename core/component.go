package core

import (
	"log"
	"os"
	"strings"
)

type Component struct {
	Name     string `yaml:"name"`
	Provides string `yaml:"provides"`
	Type     string `yaml:"type"`
	Filename string `yaml:"_"`
}

func (c *Component) LoadYaml(fn string) error {
	err := LoadYaml(fn, c)
	c.Filename = fn
	return err
}

func LoadComponents(dirname string, list *[]Component) error {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		log.Printf("readdir() error on %s: %s\n", dirname, err)
		return err
	}

	for _, entry := range entries {
		n := entry.Name()
		if entry.IsDir() {
			LoadComponents(dirname+"/"+n, list)
		} else {
			if strings.HasSuffix(n, ".yaml") || strings.HasSuffix(n, ".yml") {
				comp := Component{}
				err = comp.LoadYaml(dirname + "/" + n)
				if err != nil {
					log.Printf("load error on %s: %s\n", dirname+"/"+n, err)
					return err
				}
				*list = append(*list, comp)
			}
		}
	}
	return nil
}
