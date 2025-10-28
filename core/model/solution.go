package model

import (
	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/util"
)

const (
	Solution_Key_Project = "@PROJECT"
	Solution_Key_InstallPrefix = "@installprefix"
)

type Solution struct {
	SpecObj
}

func (c *Solution) GetMapped(name string) string {
	return util.StrOr(c.GetStr(api.Key("package-mapping::"+name)), name)
}

func (c *Solution) GetBuildList() []string {
	return c.GetStrList("build")
}

func (c *Solution) GetPackageSpecDirs() []string {
	return c.GetStrList("packages")
}

func (c *Solution) GetPackageConfig(pkgname string) api.Entry {
	return c.GetEntry(api.Key("package-config::"+pkgname))
}

func (c *Solution) SetProject(prj * Project) {
	c.Put(Solution_Key_Project, prj)
}

func (c *Solution) LoadYaml(fn string) error {
	if err := c.SpecObj.LoadYaml(fn); err != nil {
		return err
	}

	// initialize some default keys
	api.SetDefaultStr(c, Solution_Key_InstallPrefix, "${"+Solution_Key_Project+"::"+KeyProjectInstallPrefix+"}")
	return nil
}
