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

type ComponentList = []Component

func (c *Component) LoadYaml(fn string) error {
	err := LoadYaml(fn, c)
	c.Filename = fn
	return err
}

type ComponentsDB struct {
	Components map[string]Component
	Provides map[string]ComponentList
}

func (db * ComponentsDB) Add(comp Component) {
	if db.Components == nil {
		db.Components = make(map[string]Component)
	}
	db.Components[comp.Name] = comp
	if db.Provides == nil {
		db.Provides = make(map[string]ComponentList)
	}
}

func (db * ComponentsDB) LoadComponents(dirname string) error {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		log.Printf("readdir() error on %s: %s\n", dirname, err)
		return err
	}

	for _, entry := range entries {
		n := entry.Name()
		if entry.IsDir() {
			db.LoadComponents(dirname+"/"+n)
		} else {
			if strings.HasSuffix(n, ".yaml") || strings.HasSuffix(n, ".yml") {
				comp := Component{}
				err = comp.LoadYaml(dirname + "/" + n)
				if err != nil {
					log.Printf("load error on %s: %s\n", dirname+"/"+n, err)
					return err
				}
				db.Add(comp)
			}
		}
	}
	return nil
}
