// SPDX-License-Identifier: AGPL-3.0-or-later
package sysroot

import (
	"fmt"
	"log"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

func doDeploySysroot(prj *model.Project, parent *model.Package, pkgName string, sysroot string) error {

	pkg := prj.LookupPackage(pkgName)
	if pkg == nil {
		return fmt.Errorf("sysroot: cant resolve component %s\n", pkgName)
	}

	parentName := parent.GetName()

	if !pkg.IsBuildable() {
		log.Printf("[%s] skip sysroot deployment for %s\n", parentName, pkgName)
		return nil
	}

	flagName := api.Key("@@sysroot-deployed-" + pkgName)

	if parent.GetBool(flagName, false) {
		return nil
	}

	for _, dep := range pkg.GetAllDeps() {
		if err := doDeploySysroot(prj, parent, dep, sysroot); err != nil {
			return err
		}
	}

	log.Printf("[%s] deploying %s into sysroot %s\n", parentName, pkgName, sysroot)

	tarFile := pkg.GetBinpkgTarball()
	log.Printf("[%s] binpkg tarball for %s: %s\n", parentName, pkgName, tarFile)

	if err := util.UnpackTarballGz(tarFile, sysroot+"@"+pkg.GetScope()); err != nil {
		return err
	}

	parent.SetBool(flagName, true)
	return nil
}

// FIXME: not honoring build flags yet
func DeploySysroot(prj *model.Project, pkg *model.Package) error {

	sysroot := pkg.GetStr("@sysroot")

	for _, depname := range pkg.GetAllDeps() {
		if err := doDeploySysroot(prj, pkg, depname, sysroot); err != nil {
			return err
		}
	}

	return nil
}
