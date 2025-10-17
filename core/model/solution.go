package model

import (
	"log"

	"github.com/metux/mpbt/core/util"
	"github.com/metux/go-magicdict/api"
	"github.com/metux/go-magicdict/magic"
)

type Solution struct {
	magic.MagicDict
	Filename string
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
	return util.StrOr(api.GetStr(c, api.Key("package-mapping::"+name)), name)
}

func (c *Solution) GetBuildList() []string {
	return api.GetStrList(c, "build")
}
