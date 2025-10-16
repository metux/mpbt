package core

import (
	"fmt"
	"log"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/util"
)

type Project struct {
	model.Project
}

// resolve what we need to clone
func (prj *Project) ResolvePkg(name string) error {
	log.Printf("NAME=%s\n", name)
	comp := prj.LookupComponent(name)
	if comp == nil {
		return fmt.Errorf("Cant resolve component %s\n", name)
	}

	for _, dep := range comp.GetAllDeps() {
		log.Printf("DEP: %s\n", dep)
		if err := prj.ResolvePkg(dep); err != nil {
			return err
		}
	}

	log.Printf("[%s] cloning component\n", name)
	return prj.CloneComponent(comp)
}

func (prj *Project) CloneComponent(comp *model.Component) error {
	gitspec := comp.Sources.Git
	remotename := "origin"

	if gitspec == nil {
		log.Printf("[%s] no gitspec - nothing to clone here\n", comp.Name)
		return nil
	}

	comp.CloneDir = prj.SourceRoot + "/" + comp.Name

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

// FIXME: yet need to check for recursion and feature flags
func (prj *Project) Resolve() {
	prj.SourceRoot = "sources"
	for _, b := range prj.Solution.Build {
		if err := prj.ResolvePkg(b); err != nil {
			log.Printf("ERR on %s: %s\n", b, err)
		}
	}
}
