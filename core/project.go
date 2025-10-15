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
	comp := prj.LookupComponent(name)
	if comp == nil {
		return fmt.Errorf("Cant resolve component %s\n", name)
	}

	for _, dep := range comp.GetAllDeps() {
		if err := prj.ResolvePkg(dep); err != nil {
			return err
		}
	}

	return prj.CloneComponent(comp)
}

func (prj *Project) CloneComponent(comp *model.Component) error {

	gitspec := comp.Sources.Git
	remotename := "origin"

	if gitspec == nil {
		return nil
	}

	comp.CloneDir = prj.SourceRoot + "/" + comp.Name

	log.Printf("should clone: %s\n", comp.Name)
	log.Printf("clonedir %s\n", comp.CloneDir)
	log.Printf("url=%s\n", gitspec.Url)
	log.Printf("ref=%s\n", gitspec.Ref)
	log.Printf("depths=%d\n", gitspec.Depth)
	log.Printf("fetch=%v+\n", gitspec.Fetch)

	util.ExecCmd([]string{"git", "init", comp.CloneDir})
	util.ExecCmd([]string{"git", "-C", comp.CloneDir, "config", "remote." + remotename + ".url", gitspec.Url})

	c1 := []string{"git", "-C", comp.CloneDir, "fetch"}
	if gitspec.Depth > 0 {
		c1 = append(c1, fmt.Sprintf("--depth=%d", gitspec.Depth))
	}

	c1 = append(c1, "origin")
	c1 = append(c1, gitspec.Fetch...)

	util.ExecCmd(c1)
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
