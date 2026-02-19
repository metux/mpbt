// SPDX-License-Identifier: AGPL-3.0-or-later
package util

import (
	"github.com/metux/go-magicdict/api"
	"github.com/metux/go-magicdict/magic"
)

type SpecObj struct {
	magic.MagicDict
}

func (c *SpecObj) LoadYaml(fn string) error {
	d, err := magic.YamlLoad(fn, "")
	if err != nil {
		return err
	}
	c.MagicDict = d
	return nil
}

func (c *SpecObj) SetStr(name api.Key, val string) error {
	return api.SetStr(c, name, val)
}

func (c *SpecObj) SetDefaultStr(name api.Key, val string) error {
	return api.SetDefaultStr(c, name, val)
}

func (c *SpecObj) SetDefaultInt(name api.Key, val int) error {
	return api.SetDefaultInt(c, name, val)
}

func (c *SpecObj) GetStr(name api.Key) string {
	return api.GetStr(c, name)
}

func (c *SpecObj) GetStrList(name api.Key) []string {
	return api.GetStrList(c, name)
}

func (c *SpecObj) GetEntry(k api.Key) api.Entry {
	return api.GetEntry(c, k)
}
