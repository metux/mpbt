// SPDX-License-Identifier: AGPL-3.0-or-later
package model

import (
	"log"
	"os"
	"path/filepath"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/go-magicdict/magic"
	"github.com/metux/mpbt/core/model/sources"
	"github.com/metux/mpbt/core/util"
)

const (
	Package_Key_BuildDepends  = "build-depends"
	Package_Key_Buildsystem   = "buildsystem"
	Package_Key_BuildDir      = "@builddir"
	Package_Key_Depends       = "depends"
	Package_Key_Filename      = "@filename"
	Package_Key_Basename      = "@basename"
	Package_Key_Name          = "name"
	Package_Key_Slug          = "@slug"
	Package_Key_Project       = "@PROJECT"
	Package_Key_Solution      = "@SOLUTION"
	Package_Key_Provides      = "provides"
	Package_Key_SourceDir     = "source-dir"
	Package_Key_Type          = "type"
	Package_Key_InstallPrefix = "install-prefix"
	Package_Key_StatDir       = "@statdir"
	Package_Key_Parallel      = "parallel"
	Package_Key_Destdir       = "@destdir"
	Package_Key_BinaryTarball = "@binary-tarball"

	Package_Default_BuildDir = "__BUILD"
)

type Package struct {
	util.SpecObj

	// internal only, not in YAML
	cacheGit *sources.Git `yaml:"-"`
}

type PackageMap = map[string]*Package

func (pkg *Package) LoadYaml(fn string) error {
	d, err := magic.YamlLoad(fn, "")
	if err != nil {
		return err
	}
	pkg.MagicDict = d

	// init some presets
	pkg.SetStr(Package_Key_Filename, fn)
	pkg.SetDefaultStr(Package_Key_SourceDir, "${"+Package_Key_Project+"::"+Project_Key_SourceRoot+"}/${name}")
	pkg.SetDefaultStr(Package_Key_InstallPrefix, "${"+Package_Key_Solution+"::"+Solution_Key_InstallPrefix+"}")
	pkg.SetDefaultStr(Package_Key_Parallel, "${"+Package_Key_Solution+"::"+Solution_Key_Parallel+"}")
	pkg.SetDefaultStr(Package_Key_StatDir, "${"+Package_Key_Project+"::"+Project_Key_Workdir+"}/stat")
	pkg.SetDefaultStr(Package_Key_BuildDir, "${"+Package_Key_Project+"::"+Project_Key_Workdir+"}/build/${"+Package_Key_Name+"}")

	// private, should not be used directly in user configs
	pkg.SetStr("@binary-image", "${@PROJECT::@workdir}/install/${name}/image")
	pkg.SetStr(Package_Key_BinaryTarball, "${@PROJECT::@workdir}/tarball/${name}.tar.gz")

	// link some defaults to solution
	pkg.SetDefaultStr("cmake-args", "${@SOLUTION::package-defaults::cmake-args}")
	pkg.SetDefaultStr("cmake-extra-args", "${SOLUTION::package-defaults::cmake-extra-args}")

	return nil
}

func (c Package) GetAllDeps() util.StringList {
	return append(c.GetBuildDepends(), c.GetDepends()...)
}

// tell wether the component should/can be built
// eg. "system" type has nothing to build at all
func (c Package) IsBuildable() bool {
	t := c.GetType()
	return t != "system" && t != "fetchonly"
}

func (c Package) IsFetchable() bool {
	return c.GetGit() != nil
}

func (c Package) GetBuildsystem() string {
	return c.GetStr(Package_Key_Buildsystem)
}

func (c Package) GetType() string {
	return c.GetStr(Package_Key_Type)
}

func (c Package) GetDepends() []string {
	return api.GetStrList(c, Package_Key_Depends)
}

func (c Package) GetBuildDepends() []string {
	return api.GetStrList(c, Package_Key_BuildDepends)
}

func (c Package) GetName() string {
	return c.GetStr(Package_Key_Name)
}

func (c Package) GetProvides() []string {
	return c.GetStrList(Package_Key_Provides)
}

func (pkg Package) GetGit() *sources.Git {
	if pkg.cacheGit != nil {
		return pkg.cacheGit
	}

	ent, err := pkg.Get("sources::git")
	if err != nil {
		log.Printf("[%s] failed getting git entry: %+v\n", pkg.GetName(), err)
	}

	if ent == nil {
		return nil
	}

	git := sources.Git{
		Url:             api.GetStr(ent, "url"),
		Ref:             api.GetStr(ent, "ref"),
		Depth:           api.GetInt(ent, "depth", 0),
		Fetch:           api.GetStrList(ent, "fetch"),
		PostCheckoutCmd: api.GetStrList(ent, "post-checkout-cmd"),
	}

	pkg.cacheGit = &git
	return pkg.cacheGit
}

func (pkg Package) GetSourceDir() string {
	return pkg.GetStr(Package_Key_SourceDir)
}

func (pkg Package) GetBuildDir() string {
	return pkg.GetStr(Package_Key_BuildDir)
}

func (pkg Package) SetSourceDir(src string) error {
	return pkg.SetStr(Package_Key_SourceDir, src)
}

func (pkg Package) GetInstallPrefix() string {
	return pkg.GetStr(Package_Key_InstallPrefix)
}

func (pkg Package) SetProject(prj *Project) {
	pkg.Put(Package_Key_Project, prj)
	pkg.Put(Package_Key_Solution, prj.Solution)
}

func (pkg Package) GetSlug() string {
	return pkg.GetStr(Package_Key_Slug)
}

func (pkg Package) GetStatDir() string {
	return pkg.GetStr(Package_Key_StatDir)
}

func (pkg Package) GetParallel() int {
	return pkg.GetInt(Package_Key_Parallel, 0)
}

func (pkg Package) GetDestdir() string {
	return pkg.GetStr(Package_Key_Destdir)
}

func (pkg Package) EnableBinpkg() bool {
	return pkg.GetBool(Package_Key_Solution+"::enable-binpkg", true)
}

func (pkg Package) GetStatfileBuilt() string {
	return pkg.GetStatDir() + "/" + pkg.GetSlug() + ".DONE"
}

func (pkg Package) CheckStatBuilt() bool {
	statfile := pkg.GetStatfileBuilt()
	if !util.FileExists(statfile) {
		return false
	}

	old_sha_buf, err := os.ReadFile(statfile)
	if err != nil {
		return false
	}

	old_sha := string(old_sha_buf)
	new_sha := pkg.GetGitRepo().GetCurrentRev()

	if old_sha != new_sha {
		return false
	}

	if pkg.EnableBinpkg() {
		tarfile := pkg.GetBinpkgTarball()
		return util.FileExists(tarfile)
	}

	return true
}

func (pkg Package) MarkStatBuilt() error {
	statdir := pkg.GetStatDir()
	os.MkdirAll(statdir, 0755)

	statfile := pkg.GetStatfileBuilt()
	pkgName := pkg.GetName()

	repo := pkg.GetGitRepo()
	rev := repo.GetCurrentRev()

	log.Printf("[%s] current rev is %s\n", pkgName, rev)

	file, err := os.Create(statfile)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(rev); err != nil {
		return err
	}

	if err := file.Sync(); err != nil {
		return err
	}

	return nil
}

func (pkg Package) GetGitRepo() util.GitRepo {
	return util.GitRepo{Dir: pkg.GetSourceDir()}
}

func (pkg Package) GetBinpkgTarball() string {
	return pkg.GetStr(Package_Key_BinaryTarball)
}

func LoadPackageYaml(fn string, name string) (*Package, error) {
	pkg := Package{}
	if err := pkg.LoadYaml(fn); err != nil {
		return nil, err
	}
	pkg.SetDefaultStr(Package_Key_Name, name)
	pkg.SetDefaultStr(Package_Key_Basename, filepath.Base(name))
	pkg.SetDefaultStr(Package_Key_Slug, util.SanitizeFilename(name))
	return &pkg, nil
}
