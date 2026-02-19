// SPDX-License-Identifier: AGPL-3.0-or-later
package model

import (
	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/util"
)

const (
	Solution_Key_Project        = "@PROJECT"
	Solution_Key_InstallPrefix  = "install-prefix"
	Solution_Key_PackageMapping = "package-mapping"
	Solution_Key_Build          = "build"
	Solution_Key_Packages       = "packages"
	Solution_Key_PackageConfig  = "package-config"
)

type Solution struct {
	util.SpecObj
}

func (c *Solution) GetMapped(name string) string {
	return util.StrOr(c.GetStr(api.Key(Solution_Key_PackageMapping+"::"+name)), name)
}

func (c *Solution) GetBuildList() []string {
	return c.GetStrList(Solution_Key_Build)
}

func (c *Solution) GetPackageSpecDirs() []string {
	return c.GetStrList(Solution_Key_Packages)
}

func (c *Solution) GetPackageConfig(pkgname string) api.Entry {
	return c.GetEntry(api.Key(Solution_Key_PackageConfig + "::" + pkgname))
}

func (c *Solution) SetProject(prj *Project) {
	c.Put(Solution_Key_Project, prj)
}

func (c *Solution) LoadYaml(fn string) error {
	if err := c.SpecObj.LoadYaml(fn); err != nil {
		return err
	}

	// initialize some default keys
	c.SetDefaultStr(Solution_Key_InstallPrefix, "${"+Solution_Key_Project+"::"+Project_Key_InstallPrefix+"}")
	return nil
}
