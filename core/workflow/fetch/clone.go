package fetch

import (
	"log"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
	"github.com/metux/go-magicdict/api"
)

func ClonePackage(comp model.Package, config api.Entry) error {
	gitspec := comp.Sources.Git
	remotename := "origin"

	if gitspec == nil {
		log.Printf("[%s] no gitspec - nothing to clone here\n", comp.Name)
		return nil
	}

	repo := util.GitRepo{Dir: comp.SourceDir}

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
		return repo.SimpleCheckout(gitspec.Ref)
	}
	return nil
}
