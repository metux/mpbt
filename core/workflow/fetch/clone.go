// SPDX-License-Identifier: AGPL-3.0-or-later
package fetch

import (
	"log"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/model/sources"
	"github.com/metux/mpbt/core/util"
)

func addRemote(repo util.GitRepo, remote sources.GitRemote) error {
	log.Printf("adding remote %s -- %+v\n", remote.Name, remote)

	if err := repo.SetRemoteUrl(remote.Name, remote.Url); err != nil {
		return err
	}
	if err := repo.Fetch(remote.Depth, remote.Name, remote.Fetch...); err != nil {
		return err
	}
	if err := repo.ConfigFetch(remote.Name, remote.Fetch...); err != nil {
		return err
	}
	return nil
}

func addConfig(pkg *model.Package, repo util.GitRepo, config map[api.Key]string) error {
	if config == nil {
		log.Printf("[%s] no config\n", pkg.GetName())
		return nil
	}

	if pkg.GetBool("@git-config-applied", false) {
		return nil
	}

	for idx, val := range config {
		log.Printf("[%s] key=%s val=%s\n", pkg.GetName(), idx, val)
		repo.ConfigSet(string(idx), val)
	}

	pkg.SetBool("@git-config-applied", true)

	return nil
}

func ClonePackage(pkg *model.Package, config api.Entry) error {
	gitspec := pkg.GetGit()

	if gitspec == nil {
		log.Printf("[%s] no gitspec - nothing to clone here\n", pkg.GetName())
		return nil
	}

	repo := pkg.GetGitRepo()

	addConfig(pkg, repo, gitspec.Config)

	// FIXME: dont fetch if already checked-out
	if repo.IsCheckedOut() {
		return nil
	}

	if err := repo.Init(); err != nil {
		return err
	}

	for _, remote := range gitspec.Remotes {
		if err := addRemote(repo, remote); err != nil {
			return err
		}
	}

	if !repo.IsCheckedOut() {
		if err := repo.SimpleCheckout(gitspec.Ref); err != nil {
			return err
		}
	}
	if len(gitspec.PostCheckoutCmd) > 0 {
		return util.ExecCmd(pkg.GetName(), gitspec.PostCheckoutCmd, pkg.GetSourceDir())
	}

	return nil
}
