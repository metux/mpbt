package workflow

import (
	"log"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

func CloneComponent(comp model.Component, prefix string) error {
	gitspec := comp.Sources.Git
	remotename := "origin"

	if gitspec == nil {
		log.Printf("[%s] no gitspec - nothing to clone here\n", comp.Name)
		return nil
	}

	comp.CloneDir = prefix + "/" + comp.Name

	repo := util.GitRepo{Dir: comp.CloneDir}
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
		repo.SimpleCheckout(gitspec.Ref)
	}
	return nil
}
