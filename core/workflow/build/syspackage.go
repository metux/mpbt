// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"fmt"
	"log"
	"time"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

func SysPackage(pkg *model.Package, cf api.Entry) error {
	name := pkg.GetName()

	if pkg.GetBool("@pkg-config-found", false) {
		return nil
	}

	// FIXME: pull this from config or env ?
	pkgconfig_cmd := []string{"pkg-config", "--modversion"}

	if (true) {
		start := time.Now()
		out := util.ExecOut([]string{"/bin/echo", "wurst"}, "")
		if out == "" {
			return fmt.Errorf("/bin/echo: %s", name)
		}
		elapsed := time.Now().Sub(start)
		log.Printf("[%s] /bin/echo result: %s [%d usec]\n", name, out, elapsed)
	}

	for _, pn := range pkg.GetStrList(model.Package_Key_PkgConfig) {
		start := time.Now()
		out := util.ExecOut(append(pkgconfig_cmd, pn), "")
		if out == "" {
			return fmt.Errorf("[%s] missing pkg-config package: %s", name, pn)
		}
		elapsed := time.Now().Sub(start)
		log.Printf("[%s] pkg-config probe result: %s [%d usec]\n", name, out, elapsed)
	}

	pkg.SetBool("@pkg-config-found", true)

	return nil
}
