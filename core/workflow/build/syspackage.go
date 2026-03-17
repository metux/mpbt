// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"fmt"
	"log"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

func SysPackage(pkg *model.Package, cf api.Entry) error {
	name := pkg.GetName()

	// FIXME: pull this from config or env ?
	pkgconfig_cmd := []string{"pkg-config", "--modversion"}

	for _, pn := range pkg.GetStrList(model.Package_Key_PkgConfig) {
		out := util.ExecOut(append(pkgconfig_cmd, pn), "")
		if out == "" {
			return fmt.Errorf("[%s] missing pkg-config package: %s", name, pn)
		}
		log.Printf("[%s] pkg-config probe result: %s\n", name, out)
	}

	return nil
}
