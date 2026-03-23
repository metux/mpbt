// SPDX-License-Identifier: AGPL-3.0-or-later
package fetch

import (
	"log"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/model/sources"
	"github.com/metux/mpbt/core/util"
)

func addConfig(pkg *model.Package, repo util.GitRepo, config map[api.Key]string) error {
	if config == nil {
		return nil
	}

	for idx, val := range config {
		log.Printf("[%s] adding git config: key=%s val=%s\n", pkg.GetName(), idx, val)
		if err := repo.ConfigSet(string(idx), val); err != nil {
			return nil
		}
	}

	return nil
}

func updatePackage(pkg *model.Package, gitspec *sources.Git, repo util.GitRepo) error {
	log.Printf("[%s] updating package ...\n", pkg.GetName())

	if err := addConfig(pkg, repo, gitspec.Config); err != nil {
		return err
	}

	for _, remote := range gitspec.Remotes {
		if err := repo.Fetch(remote.Depth, remote.Name, true, remote.Fetch...); err != nil {
			return err
		}
	}

	return nil
}

func clonePackage(pkg *model.Package, gitspec *sources.Git, repo util.GitRepo) error {
	log.Printf("[%s] cloning package\n", pkg.GetName())

	if err := repo.Init(); err != nil {
		return err
	}

	for _, remote := range gitspec.Remotes {
		log.Printf("[%s] adding remote %s -- %+v\n", pkg.GetName(), remote.Name, remote)
		if err := repo.SetRemoteUrl(remote.Name, remote.Url); err != nil {
			return err
		}
		if err := repo.Fetch(remote.Depth, remote.Name, false, remote.Fetch...); err != nil {
			return err
		}
		if err := repo.ConfigFetch(remote.Name, remote.Fetch...); err != nil {
			return err
		}
	}

	if err := addConfig(pkg, repo, gitspec.Config); err != nil {
		return err
	}

	if err := repo.SimpleCheckout(gitspec.Ref); err != nil {
		return err
	}

	if len(gitspec.PostCheckoutCmd) > 0 {
		return util.ExecCmd(pkg.GetName(), gitspec.PostCheckoutCmd, pkg.GetSourceDir())
	}

	return nil
}

func FetchPackage(pkg *model.Package, update bool) error {
	gitspec := pkg.GetGit()

	if gitspec == nil {
		log.Printf("[%s] no gitspec - nothing to clone here\n", pkg.GetName())
		return nil
	}

	repo := pkg.GetGitRepo()

	if !repo.IsCheckedOut() {
		return clonePackage(pkg, gitspec, repo)
	}

	if update {
		return updatePackage(pkg, gitspec, repo)
	}

	return nil
}
