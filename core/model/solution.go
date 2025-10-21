package model

import (
	"github.com/metux/go-magicdict/api"
	"github.com/metux/go-magicdict/magic"
	"github.com/metux/mpbt/core/util"
)

const (
	KeySolutionProject = "@PROJECT"
)

type Solution struct {
	magic.MagicDict
}

func (c *Solution) LoadYaml(fn string) error {
	d, err := magic.YamlLoad(fn, "")
	if err != nil {
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

func (c *Solution) GetPackageSpecDirs() []string {
	return api.GetStrList(c, "packages")
}

func (c *Solution) GetPackageConfig(pkgname string) api.Entry {
	return api.GetEntry(c, api.Key("package-config::"+pkgname))
}
