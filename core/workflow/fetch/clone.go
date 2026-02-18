// SPDX-License-Identifier: AGPL-3.0-or-later
package fetch

import (
	"log"

	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

func ClonePackage(pkg model.Package, config api.Entry) error {
	gitspec := pkg.GetGit()
	remotename := "origin"

	if gitspec == nil {
		log.Printf("[%s] no gitspec - nothing to clone here\n", pkg.GetName())
		return nil
	}

	repo := util.GitRepo{Dir: pkg.GetSourceDir()}

	// FIXME: dont fetch if already checked-out
	if repo.IsCheckedOut() {
		return nil
	}

	if err := repo.Init(); err != nil {
		return err
	}
	if err := repo.SetRemoteUrl(remotename, gitspec.Url); err != nil {
		return err
	}
	if err := repo.Fetch(gitspec.Depth, remotename, gitspec.Fetch...); err != nil {
		return err
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
