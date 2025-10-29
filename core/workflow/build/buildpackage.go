// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"fmt"
	"log"
	"os"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

func BuildPackage(pkg *model.Package, cf api.Entry) error {
	if !pkg.IsBuildable() {
		return nil
	}

	bs := pkg.GetBuildsystem()
	if bs == "meson" {
		return BuildWithBuilder(pkg, cf, &MesonBuilder{})
	}
	if bs == "autotools" {
		return BuildWithBuilder(pkg, cf, &AutotoolsBuilder{})
	}

	return fmt.Errorf("%s: no known build system defined: %s", pkg.GetName(), bs)
}

func BuildWithBuilder(pkg *model.Package, cf api.Entry, b model.IBuilder) error {
	b.Init(pkg, cf)

	pkgName := pkg.GetName()

	statdir := pkg.GetStatDir()
	os.MkdirAll(statdir, 0755)

	statfile := statdir + "/" + pkg.GetSlug() + ".DONE"
	log.Printf("statfile: %s\n", statfile)

	if _, err := os.Stat(statfile); err == nil {
		log.Printf("[%s] Package already built\n", pkgName)
		return nil
	}

	log.Printf("[%s] building ...\n", pkgName)

	if err := b.RunPrepare(); err != nil {
		log.Printf("[%s] Prepare error: %s\n", pkgName, err)
		return err
	}

	if err := b.RunConfigure(); err != nil {
		log.Printf("[%s] Configure error: %s\n", pkgName, err)
		return err
	}

	if err := b.RunBuild(); err != nil {
		log.Printf("[%s] Build error: %s\n", pkgName, err)
		return err
	}

	if err := b.RunInstall(); err != nil {
		log.Printf("[%s] Install error: %s\n", pkgName, err)
		return err
	}

	if err := util.ExecCmd(pkgName, []string{"touch", statfile}, "."); err != nil {
		log.Printf("[%s] Error: %s\n", pkgName, err)
		return err
	}

	log.Printf("[%s] finished build\n", pkgName)

	return nil
}
