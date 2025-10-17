package model

import (
	"log"

	"github.com/metux/go-magicdict/magic"
	"github.com/metux/go-magicdict/api"
)

type Solution struct {
	magic.MagicDict
	Filename         string
}

func (c *Solution) LoadYaml(fn string) error {
	c.Filename = fn
	d, err := magic.YamlLoad(fn, "")
	if err != nil {
		log.Printf("failed loading magic dict %s -> %s\n", fn, err)
		return err
	}
	c.MagicDict = d
	return nil
}

func (c *Solution) GetMapped(name string) string {
	p1 := api.GetStr(c, api.Key("package-mapping::"+name))
	if p1 == "" {
		log.Printf("not mapped: %s\n", name)
		return name
	}

	log.Printf("mapped %s => %s\n", name, p1)
	return p1
}

func (c *Solution) GetBuildList() [] string {
	return api.GetStrList(c, "build")
}
